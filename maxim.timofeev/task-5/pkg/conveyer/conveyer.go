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

func (c *conveyor) ensureChannel(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()
	ch, ok := c.chans[name]
	if !ok {
		ch = make(chan string, c.size)
		c.chans[name] = ch
	}
	return ch
}

func (c *conveyor) getChannel(name string) (chan string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	ch, ok := c.chans[name]
	return ch, ok
}

func (c *conveyor) RegisterDecorator(
	fn func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {

	c.ensureChannel(input)
	c.ensureChannel(output)
	c.decorators = append(c.decorators, decoratorReg{
		fn:      fn,
		inName:  input,
		outName: output,
	})
}

func (c *conveyor) RegisterMultiplexer(
	fn func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	for _, in := range inputs {
		c.ensureChannel(in)
	}
	c.ensureChannel(output)
	c.multiplexers = append(c.multiplexers, multiplexerReg{
		fn:      fn,
		inNames: inputs,
		outName: output,
	})
}

func (c *conveyor) RegisterSeparator(
	fn func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	c.ensureChannel(input)

	for _, o := range outputs {
		c.ensureChannel(o)
	}

	c.separators = append(c.separators, separatorReg{
		fn:       fn,
		inName:   input,
		outNames: outputs,
	})
}

func (c *conveyor) Run(ctx context.Context) error {
	// don't shadow caller ctx; we may still cancel locally on error
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	errCh := make(chan error, 1)

	runHandler := func(fn func(ctx context.Context) error) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := fn(ctx); err != nil {
				// send first error non-blocking
				select {
				case errCh <- err:
				default:
				}
			}
		}()
	}

	// start handlers
	for _, d := range c.decorators {
		in, _ := c.getChannel(d.inName)
		out, _ := c.getChannel(d.outName)
		fn := d.fn
		runHandler(func(ctx context.Context) error {
			return fn(ctx, in, out)
		})
	}

	for _, m := range c.multiplexers {
		inputs := make([]chan string, 0, len(m.inNames))
		for _, n := range m.inNames {
			ch, _ := c.getChannel(n)
			inputs = append(inputs, ch)
		}
		out, _ := c.getChannel(m.outName)
		fn := m.fn
		runHandler(func(ctx context.Context) error {
			return fn(ctx, inputs, out)
		})
	}

	for _, s := range c.separators {
		in, _ := c.getChannel(s.inName)
		outputs := make([]chan string, 0, len(s.outNames))
		for _, n := range s.outNames {
			ch, _ := c.getChannel(n)
			outputs = append(outputs, ch)
		}
		fn := s.fn
		runHandler(func(ctx context.Context) error {
			return fn(ctx, in, outputs)
		})
	}

	// wait for all handlers to finish in a non-blocking way
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	// wait for either: external ctx cancel, first handler error, or all handlers done
	var retErr error
	select {
	case <-ctx.Done():
		retErr = ctx.Err()
	case e := <-errCh:
		// got an error from handler -> cancel others
		retErr = e
		cancel()
	case <-done:
		retErr = nil
	}

	// ensure all goroutines finished
	<-done

	// IMPORTANT: do NOT close channels here — handlers are responsible for closing their own outputs.
	// Closing here may result in double-close panics (even if recovered).
	return retErr
}

func (c *conveyor) Send(input string, data string) error {
	ch, ok := c.getChannel(input)
	if !ok {
		return ErrChanNotFound
	}
	// no recover: tests expect simple behavior; sending to closed channel is a user error in tests
	ch <- data
	return nil
}

func (c *conveyor) Recv(output string) (string, error) {
	ch, ok := c.getChannel(output)
	if !ok {
		return "", ErrChanNotFound
	}
	v, ok := <-ch
	if !ok {
		return Undefined, nil
	}
	return v, nil
}
