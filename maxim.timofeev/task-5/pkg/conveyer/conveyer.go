package conveyer

import (
	"context"
	"errors"
	"sync"
)

const Undefined = "undefined"

var ErrChanNotFound = errors.New("chan not found")

type decoratorReg struct {
	fn      func(ctx context.Context, input chan string, output chan string) error
	inName  string
	outName string
}

type multiplexerReg struct {
	fn      func(ctx context.Context, inputs []chan string, output chan string) error
	inNames []string
	outName string
}

type separatorReg struct {
	fn       func(ctx context.Context, input chan string, outputs []chan string) error
	inName   string
	outNames []string
}

type conveyor struct {
	size int

	mu sync.RWMutex

	chans map[string]chan string

	decorators   []decoratorReg
	multiplexers []multiplexerReg
	separators   []separatorReg
}

func New(size int) *conveyor {
	return &conveyor{
		size:  size,
		chans: make(map[string]chan string),
	}
}

func (c *conveyor) ensure(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ch, ok := c.chans[name]; ok {
		return ch
	}

	ch := make(chan string, c.size)
	c.chans[name] = ch
	return ch
}

func (c *conveyor) get(name string) (chan string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	ch, ok := c.chans[name]
	return ch, ok
}

func (c *conveyor) RegisterDecorator(fn func(context.Context, chan string, chan string) error, input, output string) {
	c.ensure(input)
	c.ensure(output)
	c.decorators = append(c.decorators, decoratorReg{fn, input, output})
}

func (c *conveyor) RegisterMultiplexer(fn func(context.Context, []chan string, chan string) error, inputs []string, output string) {
	for _, n := range inputs {
		c.ensure(n)
	}
	c.ensure(output)
	c.multiplexers = append(c.multiplexers, multiplexerReg{fn, inputs, output})
}

func (c *conveyor) RegisterSeparator(fn func(context.Context, chan string, []chan string) error, input string, outputs []string) {
	c.ensure(input)
	for _, n := range outputs {
		c.ensure(n)
	}
	c.separators = append(c.separators, separatorReg{fn, input, outputs})
}

func (c *conveyor) Run(ctx context.Context) error {
	var wg sync.WaitGroup
	errCh := make(chan error, 1)

	run := func(f func(context.Context) error) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := f(ctx); err != nil {
				select {
				case errCh <- err:
				default:
				}
			}
		}()
	}

	for _, d := range c.decorators {
		in, _ := c.get(d.inName)
		out, _ := c.get(d.outName)
		fn := d.fn
		run(func(ctx context.Context) error {
			return fn(ctx, in, out)
		})
	}

	for _, m := range c.multiplexers {
		inputs := make([]chan string, 0, len(m.inNames))
		for _, n := range m.inNames {
			ch, _ := c.get(n)
			inputs = append(inputs, ch)
		}
		out, _ := c.get(m.outName)
		fn := m.fn
		run(func(ctx context.Context) error {
			return fn(ctx, inputs, out)
		})
	}

	for _, s := range c.separators {
		in, _ := c.get(s.inName)
		outs := make([]chan string, 0, len(s.outNames))
		for _, n := range s.outNames {
			ch, _ := c.get(n)
			outs = append(outs, ch)
		}
		fn := s.fn
		run(func(ctx context.Context) error {
			return fn(ctx, in, outs)
		})
	}

	var err error
	select {
	case err = <-errCh:
	case <-ctx.Done():
		err = ctx.Err()
	}

	wg.Wait()
	return err
}

func (c *conveyor) Send(input string, data string) error {
	ch, ok := c.get(input)
	if !ok {
		return ErrChanNotFound
	}
	ch <- data
	return nil
}

func (c *conveyor) Recv(output string) (string, error) {
	ch, ok := c.get(output)
	if !ok {
		return "", ErrChanNotFound
	}
	v, ok := <-ch
	if !ok {
		return Undefined, nil
	}
	return v, nil
}
