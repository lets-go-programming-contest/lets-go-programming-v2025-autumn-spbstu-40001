package conveyer

import (
	"context"
	"sync"
)

type Conveyer struct {
	chansSize int
	chansMap  sync.Map
}

func (c *Conveyer) RegisterDecorator(
	fn func(
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

}

func (*Conveyer) RegisterMultiplexer(
	fn func(
		ctx context.Context,
		inputs []chan string,
		output chan string,
	) error,
	inputs []string,
	output string,
) {

}

func (*Conveyer) RegisterSeparator(
	fn func(
		ctx context.Context,
		input chan string,
		outputs []chan string,
	) error,
	input string,
	outputs []string,
) {

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
	return Conveyer{size, sync.Map{}}
}
