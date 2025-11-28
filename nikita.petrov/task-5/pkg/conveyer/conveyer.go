package conveyer

import (
	"context"
	"sync"
)

type Conveyer struct {
	chansSize int
	chansMap  sync.Map
	Decorator *func(
		ctx context.Context,
		input chan string,
		output chan string,
	) error
	Multiplexer *func(
		ctx context.Context,
		inputs []chan string,
		output chan string,
	) error
	Separator *func(
		ctx context.Context,
		input chan string,
		outputs []chan string,
	) error
}

func (c *Conveyer) RegisterDecorator(
	newDecorator func(
		ctx context.Context,
		input chan string,
		output chan string,
	) error,
	input string,
	output string,
) {
	inputChan, isExist := c.chansMap.Load(input)
	if !isExist {
		inputChan = make(chan string, c.chansSize)
		c.chansMap.Store(input, inputChan)
	}

	outputChan, isExist := c.chansMap.Load(output)
	if !isExist {
		outputChan = make(chan string, c.chansSize)
		c.chansMap.Store(input, outputChan)
	}

	*c.Decorator = newDecorator
}

func (c *Conveyer) RegisterMultiplexer(
	newMultiplexer func(
		ctx context.Context,
		inputs []chan string,
		output chan string,
	) error,
	inputs []string,
	output string,
) {
	inputChan, isExist := c.chansMap.Load(inputs)
	if !isExist {
		inputChan = make([]chan string, c.chansSize)
		c.chansMap.Store(inputs, inputChan)
	}

	outputChan, isExist := c.chansMap.Load(output)
	if !isExist {
		outputChan = make(chan string, c.chansSize)
		c.chansMap.Store(output, outputChan)
	}

	*c.Multiplexer = newMultiplexer
}

func (c *Conveyer) RegisterSeparator(
	newSeparator func(
		ctx context.Context,
		input chan string,
		outputs []chan string,
	) error,
	input string,
	outputs []string,
) {
	inputChan, isExist := c.chansMap.Load(input)
	if !isExist {
		inputChan = make([]chan string, c.chansSize)
		c.chansMap.Store(input, inputChan)
	}

	outputChan, isExist := c.chansMap.Load(outputs)
	if !isExist {
		outputChan = make([]chan string, c.chansSize)
		c.chansMap.Store(outputs, outputChan)
	}

	*c.Separator = newSeparator
}

func (*Conveyer) Run(ctx context.Context) error {

	return nil
}

func (*Conveyer) Send(input string, data string) error {

	return nil
}

func (*Conveyer) Recv(output string) (string, error) {

}

func New(size int) Conveyer {
	return Conveyer{size, sync.Map{}, nil, nil, nil}
}
