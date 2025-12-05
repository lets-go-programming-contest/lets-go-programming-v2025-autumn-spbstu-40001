package conveyer

import (
	"context"
	"errors"
	"fmt"

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
}

func New(size int) *Conveyer {
	return &Conveyer{
		chansSize:    size,
		chansMap:     make(map[string]chan string),
		decorators:   make([]Decorator, 0),
		multiplexers: make([]Multiplexer, 0),
		separators:   make([]Separator, 0),
	}
}

func (c *Conveyer) getChannel(name string) (chan string, error) {
	channel, ok := c.chansMap[name]
	if !ok {
		return nil, fmt.Errorf("%w: %s", errChanNotFound, name)
	}

	return channel, nil
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
	inputCh, inputOk := c.chansMap[d.input]
	outputCh, outputOk := c.chansMap[d.output]

	if !inputOk || !outputOk {
		return errChanNotFound
	}

	return d.DecoratorFunc(ctx, inputCh, outputCh)
}

func (c *Conveyer) runMultiplexer(ctx context.Context, m Multiplexer) error {
	inputs := make([]chan string, len(m.inputs))
	for i, inputName := range m.inputs {
		inputCh, ok := c.chansMap[inputName]
		if !ok {
			return errChanNotFound
		}
		inputs[i] = inputCh
	}

	outputCh, outputOk := c.chansMap[m.output]

	if !outputOk {
		return errChanNotFound
	}

	return m.MultiplexerFunc(ctx, inputs, outputCh)
}

func (c *Conveyer) runSeparator(ctx context.Context, s Separator) error {
	inputCh, inputOk := c.chansMap[s.input]
	if !inputOk {
		return errChanNotFound
	}

	outputs := make([]chan string, len(s.outputs))
	for i, outputName := range s.outputs {
		outputCh, ok := c.chansMap[outputName]
		if !ok {
			return errChanNotFound
		}
		outputs[i] = outputCh
	}

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
	defer c.closeChansMap()

	errGroup, ctx := errgroup.WithContext(ctx)

	for _, decorator := range c.decorators {
		d := decorator
		errGroup.Go(func() error {
			return c.runDecorator(ctx, d)
		})
	}

	for _, multiplexer := range c.multiplexers {
		m := multiplexer
		errGroup.Go(func() error {
			return c.runMultiplexer(ctx, m)
		})
	}

	for _, separator := range c.separators {
		s := separator
		errGroup.Go(func() error {
			return c.runSeparator(ctx, s)
		})
	}

	if err := errGroup.Wait(); err != nil {
		return fmt.Errorf("conveyer finished with error: %w", err)
	}

	return nil
}

func (c *Conveyer) Send(input string, data string) error {
	ch, ok := c.chansMap[input]

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
	channel, err := c.getChannel(output)
	if err != nil {
		return "", errChanNotFound
	}

	data, ok := <-channel
	if !ok {
		return errUndefinedStr, nil
	}

	return data, nil
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
