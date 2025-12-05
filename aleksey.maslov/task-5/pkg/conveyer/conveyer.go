package conveyer

import (
	"context"
	"errors"
)

var (
	ErrChanNotFound = errors.New("chan not found")
)

type Task func(ctx context.Context) error

type ConveyerType struct {
	size     int
	channels map[string]chan string
	tasks    []Task
}

func New(size int) *ConveyerType {
	return &ConveyerType{
		size:     size,
		channels: make(map[string]chan string),
		tasks:    make([]Task, 0),
	}
}

type Conveyer interface {
	RegisterDecorator(
		fn func(ctx context.Context, input chan string, output chan string) error,
		input string,
		output string,
	)
	RegisterMultiplexer(
		fn func(ctx context.Context, inputs []chan string, output chan string) error,
		inputs []string,
		output string,
	)
	RegisterSeparator(
		fn func(ctx context.Context, input chan string, outputs []chan string) error,
		input string,
		outputs []string,
	)

	Run(ctx context.Context) error
	Send(input string, data string) error
	Recv(output string) (string, error)
}
