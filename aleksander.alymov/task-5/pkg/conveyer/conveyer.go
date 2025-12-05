package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

var (
	ErrChanNotFound = errors.New("chan not found")
	ErrNoData       = errors.New("no data")
)

const Undefined = "undefined"

type conveyer struct {
	size     int
	mu       sync.RWMutex
	chans    map[string]chan string
	handlers []handler
	running  bool
}

type handler struct {
	fn          interface{}
	inputNames  []string
	outputNames []string
	htype       handlerType
}

type handlerType int

const (
	decorator handlerType = iota
	multiplexer
	separator
)

func New(size int) *conveyer {
	return &conveyer{
		size:  size,
		chans: make(map[string]chan string),
	}
}

func (c *conveyer) getOrCreateChan(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ch, ok := c.chans[name]; ok {
		return ch
	}

	ch := make(chan string, c.size)
	c.chans[name] = ch
	return ch
}

func (c *conveyer) getChan(name string) (chan string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	ch, ok := c.chans[name]
	if !ok {
		return nil, ErrChanNotFound
	}
	return ch, nil
}

func (c *conveyer) RegisterDecorator(
	fn func(context.Context, chan string, chan string) error,
	input, output string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.handlers = append(c.handlers, handler{
		fn:          fn,
		inputNames:  []string{input},
		outputNames: []string{output},
		htype:       decorator,
	})
}

func (c *conveyer) RegisterMultiplexer(
	fn func(context.Context, []chan string, chan string) error,
	inputs []string,
	output string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.handlers = append(c.handlers, handler{
		fn:          fn,
		inputNames:  inputs,
		outputNames: []string{output},
		htype:       multiplexer,
	})
}

func (c *conveyer) RegisterSeparator(
	fn func(context.Context, chan string, []chan string) error,
	input string,
	outputs []string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.handlers = append(c.handlers, handler{
		fn:          fn,
		inputNames:  []string{input},
		outputNames: outputs,
		htype:       separator,
	})
}

func (c *conveyer) Run(ctx context.Context) error {
	c.mu.Lock()
	if c.running {
		c.mu.Unlock()
		return errors.New("conveyer already running")
	}
	c.running = true
	c.mu.Unlock()

	defer func() {
		c.mu.Lock()
		c.running = false
		c.mu.Unlock()
	}()

	g, ctx := errgroup.WithContext(ctx)

	for _, h := range c.handlers {
		h := h
		g.Go(func() error {
			return c.runHandler(ctx, h)
		})
	}

	return g.Wait()
}

func (c *conveyer) runHandler(ctx context.Context, h handler) error {
	switch h.htype {
	case decorator:
		fn := h.fn.(func(context.Context, chan string, chan string) error)
		input := c.getOrCreateChan(h.inputNames[0])
		output := c.getOrCreateChan(h.outputNames[0])
		defer close(output)
		return fn(ctx, input, output)

	case multiplexer:
		fn := h.fn.(func(context.Context, []chan string, chan string) error)
		inputs := make([]chan string, len(h.inputNames))
		for i, name := range h.inputNames {
			inputs[i] = c.getOrCreateChan(name)
		}
		output := c.getOrCreateChan(h.outputNames[0])
		defer close(output)
		return fn(ctx, inputs, output)

	case separator:
		fn := h.fn.(func(context.Context, chan string, []chan string) error)
		input := c.getOrCreateChan(h.inputNames[0])
		outputs := make([]chan string, len(h.outputNames))
		for i, name := range h.outputNames {
			outputs[i] = c.getOrCreateChan(name)
		}
		defer func() {
			for _, out := range outputs {
				close(out)
			}
		}()
		return fn(ctx, input, outputs)

	default:
		return errors.New("unknown handler type")
	}
}

func (c *conveyer) Send(input string, data string) error {
	ch := c.getOrCreateChan(input)

	select {
	case ch <- data:
		return nil
	default:
		return errors.New("channel is full")
	}
}

func (c *conveyer) Recv(output string) (string, error) {
	ch, err := c.getChan(output)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrChanNotFound, output)
	}

	select {
	case val, ok := <-ch:
		if !ok {
			return Undefined, nil
		}
		return val, nil
	default:
		return "", ErrNoData
	}
}
