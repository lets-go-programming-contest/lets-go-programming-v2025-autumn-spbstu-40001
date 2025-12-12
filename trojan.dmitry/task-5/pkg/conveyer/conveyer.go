package conveyer

import (
	"context"
	"errors"
	"sync"

	"golang.org/x/sync/errgroup"
)

const UndefinedMsg = "undefined"

var ErrChanNotFound = errors.New("chan not found")

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

type Conveyer struct {
	size int

	chans   map[string]chan string
	chansMu sync.RWMutex

	decorators   []decoratorHandler
	multiplexers []multiplexerHandler
	separators   []separatorHandler

	cancelFunc context.CancelFunc
	running    bool
	runningMu  sync.RWMutex
}

func New(size int) *Conveyer {
	return &Conveyer{
		size:         size,
		chans:        make(map[string]chan string),
		decorators:   make([]decoratorHandler, 0),
		multiplexers: make([]multiplexerHandler, 0),
		separators:   make([]separatorHandler, 0),
		running:      false,
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

func (c *Conveyer) getChan(name string) (chan string, error) {
	c.chansMu.RLock()
	defer c.chansMu.RUnlock()

	ch, ok := c.chans[name]
	if !ok || ch == nil {
		return nil, ErrChanNotFound
	}
	return ch, nil
}

func (c *Conveyer) isRunning() bool {
	c.runningMu.RLock()
	defer c.runningMu.RUnlock()
	return c.running
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
	for _, out := range outputs {
		c.ensureChan(out)
	}

	c.separators = append(c.separators, separatorHandler{
		fn:      fn,
		input:   input,
		outputs: outputs,
	})
}

func (c *Conveyer) Send(input string, data string) error {
	if !c.isRunning() {
		return errors.New("conveyer is not running")
	}

	ch, err := c.getChan(input)
	if err != nil {
		return err
	}

	select {
	case ch <- data:
		return nil
	default:
		return errors.New("channel is full or closed")
	}
}

func (c *Conveyer) Recv(output string) (string, error) {
	if !c.isRunning() {
		return "", errors.New("conveyer is not running")
	}

	ch, err := c.getChan(output)
	if err != nil {
		return "", err
	}

	select {
	case v, ok := <-ch:
		if !ok {
			return UndefinedMsg, nil
		}
		return v, nil
	default:
		return "", errors.New("no data available")
	}
}

func (c *Conveyer) Run(ctx context.Context) error {
	c.runningMu.Lock()
	if c.running {
		c.runningMu.Unlock()
		return errors.New("conveyer is already running")
	}

	ctx, cancel := context.WithCancel(ctx)
	c.cancelFunc = cancel
	c.running = true
	c.runningMu.Unlock()

	defer func() {
		c.runningMu.Lock()
		c.running = false
		c.cancelFunc = nil
		c.runningMu.Unlock()

		c.chansMu.Lock()
		for name, ch := range c.chans {
			if ch != nil {
				select {
				case <-ch:
				default:
				}
				close(ch)
				c.chans[name] = nil
			}
		}
		c.chansMu.Unlock()
	}()

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

		ins := make([]chan string, len(m.inputs))
		for i, name := range m.inputs {
			ins[i] = c.ensureChan(name)
		}

		m := m
		g.Go(func() error {
			return m.fn(ctx, ins, out)
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

	return g.Wait()
}

func (c *Conveyer) Stop() {
	c.runningMu.RLock()
	if c.cancelFunc != nil {
		c.cancelFunc()
	}
	c.runningMu.RUnlock()
}
