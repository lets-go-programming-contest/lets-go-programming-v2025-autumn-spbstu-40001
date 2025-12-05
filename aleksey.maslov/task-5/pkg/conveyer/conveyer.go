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

	channel, ok := c.channels[name]
	if ok {
		return channel, nil
	}

	return nil, ErrChanNotFound
}

func (c *ConveyerType) getOrCreateChannel(name string) chan string {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	channel, ok := c.channels[name]
	if ok {
		return channel
	}

	channel = make(chan string, c.size)
	c.channels[name] = channel

	return channel
}
func (c *ConveyerType) RegisterDecorator(
	handler func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	in := c.getOrCreateChannel(input)
	out := c.getOrCreateChannel(output)

	c.mutex.Lock()
	c.tasks = append(c.tasks, func(ctx context.Context) error {
		return handler(ctx, in, out)
	})
	c.mutex.Unlock()
}

func (c *ConveyerType) RegisterMultiplexer(
	handler func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	out := c.getOrCreateChannel(output)
	ins := make([]chan string, len(inputs))

	for i, name := range inputs {
		ins[i] = c.getOrCreateChannel(name)
	}

	c.mutex.Lock()
	c.tasks = append(c.tasks, func(ctx context.Context) error {
		return handler(ctx, ins, out)
	})
	c.mutex.Unlock()
}

func (c *ConveyerType) RegisterSeparator(
	handler func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	in := c.getOrCreateChannel(input)
	outs := make([]chan string, len(outputs))

	for i, name := range outputs {
		outs[i] = c.getOrCreateChannel(name)
	}

	c.mutex.Lock()
	c.tasks = append(c.tasks, func(ctx context.Context) error {
		return handler(ctx, in, outs)
	})
	c.mutex.Unlock()
}

Run(ctx context.Context) error
Send(input string, data string) error
Recv(output string) (string, error)
