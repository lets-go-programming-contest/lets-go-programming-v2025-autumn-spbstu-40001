package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Conveyer struct {
	size int

	chans   map[string]chan string
	chansMu sync.RWMutex

	decorators   []decoratorHandler
	multiplexers []multiplexerHandler
	separators   []separatorHandler

	wg sync.WaitGroup
}

type decoratorHandler struct {
	fn     func(context.Context, chan string, chan string) error
	input  string
	output string
}

type multiplexerHandler struct {
	fn     func(context.Context, []chan string, chan string) error
	inputs []string
	output string
}

type separatorHandler struct {
	fn      func(context.Context, chan string, []chan string) error
	input   string
	outputs []string
}

func New(size int) *Conveyer {
	return &Conveyer{
		size:         size,
		chans:        make(map[string]chan string),
		wg:           sync.WaitGroup{},
		decorators:   make([]decoratorHandler, 0),
		multiplexers: make([]multiplexerHandler, 0),
		separators:   make([]separatorHandler, 0),
	}
}

func (c *Conveyer) ensureChan(name string) chan string {
	c.chansMu.Lock()
	defer c.chansMu.Unlock()
	ch, ok := c.chans[name]
	if ok && ch != nil {
		return ch
	}
	ch = make(chan string, c.size)
	c.chans[name] = ch
	return ch
}

func (c *Conveyer) getChan(name string) (chan string, bool) {
	c.chansMu.RLock()
	defer c.chansMu.RUnlock()
	ch, ok := c.chans[name]
	return ch, ok
}

func (c *Conveyer) RegisterDecorator(
	fn func(context.Context, chan string, chan string) error,
	input string,
	output string,
) {
	c.ensureChan(input)
	c.ensureChan(output)
	c.decorators = append(c.decorators, decoratorHandler{
		fn:     fn,
		input:  input,
		output: output,
	})
}

func (c *Conveyer) RegisterMultiplexer(
	fn func(context.Context, []chan string, chan string) error,
	inputs []string,
	output string,
) {
	for _, in := range inputs {
		c.ensureChan(in)
	}
	c.ensureChan(output)
	c.multiplexers = append(c.multiplexers, multiplexerHandler{
		fn:     fn,
		inputs: inputs,
		output: output,
	})
}

func (c *Conveyer) RegisterSeparator(
	fn func(context.Context, chan string, []chan string) error,
	input string,
	outputs []string,
) {
	c.ensureChan(input)
	for _, o := range outputs {
		c.ensureChan(o)
	}
	c.separators = append(c.separators, separatorHandler{
		fn:      fn,
		input:   input,
		outputs: outputs,
	})
}

func (c *Conveyer) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	for _, d := range c.decorators {
		in := c.ensureChan(d.input)
		out := c.ensureChan(d.output)

		d := d
		g.Go(func() error {
			return d.fn(ctx, in, out)
		})
	}

	for _, m := range c.multiplexers {
		out := c.ensureChan(m.output)

		inputs := make([]chan string, len(m.inputs))
		for i, name := range m.inputs {
			inputs[i] = c.ensureChan(name)
		}

		m := m
		g.Go(func() error {
			return m.fn(ctx, inputs, out)
		})
	}

	for _, s := range c.separators {
		in := c.ensureChan(s.input)

		outs := make([]chan string, len(s.outputs))
		for i, name := range s.outputs {
			outs[i] = c.ensureChan(name)
		}

		s := s
		g.Go(func() error {
			return s.fn(ctx, in, outs)
		})
	}

	err := g.Wait()

	c.chansMu.Lock()
	for _, ch := range c.chans {
		close(ch)
	}
	c.chansMu.Unlock()

	return err
}

func (c *Conveyer) Send(input string, data string) error {
	ch, ok := c.getChan(input)
	if !ok || ch == nil {
		return errors.New("chan not found")
	}

	defer func() (err error) {
		if r := recover(); r != nil {
			err = fmt.Errorf("send panic: %v", r)
			return err
		}
		return nil
	}()

	ch <- data
	return nil
}

func (c *Conveyer) Recv(output string) (string, error) {
	ch, ok := c.getChan(output)
	if !ok || ch == nil {
		return "", errors.New("chan not found")
	}

	v, ok := <-ch
	if !ok {
		return "undefined", nil
	}
	return v, nil
}
