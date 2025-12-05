package conveyer

import (
	"context"
	"errors"
	"sync"
)

const Undefined = "undefined"

var ErrChanNotFound = errors.New("chan not found")

type Task func(ctx context.Context) error

type ConveyerType struct {
	size     int
	channels map[string]chan string
	tasks    []Task
	mutex    sync.RWMutex
}

func New(size int) *ConveyerType {
	return &ConveyerType{
		size:     size,
		channels: make(map[string]chan string),
		tasks:    make([]Task, 0),
	}
}

func (c *ConveyerType) getChannel(name string) (chan string, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if channel, ok := c.channels[name]; ok {
		return channel, nil
	}

	return nil, ErrChanNotFound
}

func (c *ConveyerType) getOrCreateChannel(name string) chan string {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if channel, exists := c.channels[name]; exists {
		return channel
	}

	channel := make(chan string, c.size)
	c.channels[name] = channel

	return channel
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
