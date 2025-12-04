package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

var errChanNotFound = errors.New("chan not found")

const errUndefinedStr = "undefined"

type Decorator struct {
	DecoratorFunc func(ctx context.Context, input chan string, output chan string) error
	input         string
	output        string
}

type Multiplexer struct {
	MultiplexerFunc func(ctx context.Context, inputs []chan string, output chan string) error
	inputs          []string
	output          string
}

type Separator struct {
	SeparatorFunc func(ctx context.Context, input chan string, outputs []chan string) error
	input         string
	outputs       []string
}

type Conveyer struct {
	chansSize    int
	chansMap     map[string]chan string
	decorators   []Decorator
	multiplexers []Multiplexer
	separators   []Separator
	mutex        sync.RWMutex
}

func New(size int) *Conveyer {
	return &Conveyer{
		chansSize:    size,
		chansMap:     make(map[string]chan string),
		decorators:   make([]Decorator, 0),
		multiplexers: make([]Multiplexer, 0),
		separators:   make([]Separator, 0),
		mutex:        sync.RWMutex{},
	}
}

func (c *Conveyer) getOrCreateChannel(name string) chan string {
	if ch, exists := c.chansMap[name]; exists {
		return ch
	}

	ch := make(chan string, c.chansSize)
	c.chansMap[name] = ch
	return ch
}

func (c *Conveyer) runDecorator(ctx context.Context, d Decorator) error {
	c.mutex.RLock()
	inputCh, inputOk := c.chansMap[d.input]
	outputCh, outputOk := c.chansMap[d.output]
	c.mutex.RUnlock()

	if !inputOk || !outputOk {
		return errChanNotFound
	}

	return d.DecoratorFunc(ctx, inputCh, outputCh)
}

func (c *Conveyer) runMultiplexer(ctx context.Context, m Multiplexer) error {
	c.mutex.RLock()
	inputs := make([]chan string, len(m.inputs))
	for i, inputName := range m.inputs {
		inputCh, ok := c.chansMap[inputName]
		if !ok {
			c.mutex.RUnlock()
			return errChanNotFound
		}
		inputs[i] = inputCh
	}

	outputCh, outputOk := c.chansMap[m.output]
	c.mutex.RUnlock()

	if !outputOk {
		return errChanNotFound
	}

	return m.MultiplexerFunc(ctx, inputs, outputCh)
}

func (c *Conveyer) runSeparator(ctx context.Context, s Separator) error {
	c.mutex.RLock()
	inputCh, inputOk := c.chansMap[s.input]
	if !inputOk {
		c.mutex.RUnlock()
		return errChanNotFound
	}

	outputs := make([]chan string, len(s.outputs))
	for i, outputName := range s.outputs {
		outputCh, ok := c.chansMap[outputName]
		if !ok {
			c.mutex.RUnlock()
			return errChanNotFound
		}
		outputs[i] = outputCh
	}
	c.mutex.RUnlock()

	return s.SeparatorFunc(ctx, inputCh, outputs)
}

func (c *Conveyer) RegisterDecorator(
	fn func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	c.getOrCreateChannel(input)
	c.getOrCreateChannel(output)

	c.decorators = append(c.decorators, Decorator{
		DecoratorFunc: fn,
		input:         input,
		output:        output,
	})
}

func (c *Conveyer) RegisterMultiplexer(
	fn func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	for _, input := range inputs {
		c.getOrCreateChannel(input)
	}
	c.getOrCreateChannel(output)

	c.multiplexers = append(c.multiplexers, Multiplexer{
		MultiplexerFunc: fn,
		inputs:          inputs,
		output:          output,
	})
}

func (c *Conveyer) RegisterSeparator(
	fn func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	c.getOrCreateChannel(input)
	for _, output := range outputs {
		c.getOrCreateChannel(output)
	}

	c.separators = append(c.separators, Separator{
		SeparatorFunc: fn,
		input:         input,
		outputs:       outputs,
	})
}

func (c *Conveyer) Run(ctx context.Context) error {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	errgroup, ctx := errgroup.WithContext(ctx)

	for _, decorator := range c.decorators {
		d := decorator
		errgroup.Go(func() error {
			return c.runDecorator(ctx, d)
		})
	}

	for _, multiplexer := range c.multiplexers {
		m := multiplexer
		errgroup.Go(func() error {
			return c.runMultiplexer(ctx, m)
		})
	}

	for _, separator := range c.separators {
		s := separator
		errgroup.Go(func() error {
			return c.runSeparator(ctx, s)
		})
	}

	if err := errgroup.Wait(); err != nil {
		return fmt.Errorf("conveyer finished with error: %w", err)
	}

	return nil
}

func (c *Conveyer) Send(input string, data string) error {
	c.mutex.RLock()
	ch, ok := c.chansMap[input]
	c.mutex.RUnlock()

	if !ok {
		return errChanNotFound
	}

	select {
	case ch <- data:
		return nil
	default:
		return fmt.Errorf("channel is full")
	}
}

func (c *Conveyer) Recv(output string) (string, error) {
	c.mutex.RLock()
	ch, ok := c.chansMap[output]
	c.mutex.RUnlock()

	if !ok {
		return "", errChanNotFound
	}

	select {
	case data, ok := <-ch:
		if !ok {
			return errUndefinedStr, nil
		}
		return data, nil
	default:
		return "", fmt.Errorf("channel is empty")
	}
}

func (c *Conveyer) closeChansMap() {
	for name, ch := range c.chansMap {
		select {
		case _, ok := <-ch:
			if ok {
				close(ch)
			}
		default:
			close(ch)
		}
		delete(c.chansMap, name)
	}
}
