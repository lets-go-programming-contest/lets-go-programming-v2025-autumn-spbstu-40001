package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

const UndefinedMsg = "undefined"

var ErrChanNotFound = errors.New("chan not found")

type Task func(ctx context.Context) error

type ConveyerT struct {
	size     int
	channels map[string]chan string
	tasks    []Task
	mutex    sync.RWMutex
}

func New(size int) *ConveyerT {
	return &ConveyerT{
		size:     size,
		channels: make(map[string]chan string),
		tasks:    make([]Task, 0),
		mutex:    sync.RWMutex{},
	}
}

func (c *ConveyerT) getOrCreateChannel(name string) chan string {
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

func (c *ConveyerT) RegisterDecorator(
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

func (c *ConveyerT) RegisterMultiplexer(
	handler func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	out := c.getOrCreateChannel(output)
	inps := make([]chan string, len(inputs))

	for i, name := range inputs {
		inps[i] = c.getOrCreateChannel(name)
	}

	c.mutex.Lock()
	c.tasks = append(c.tasks, func(ctx context.Context) error {
		return handler(ctx, inps, out)
	})
	c.mutex.Unlock()
}

func (c *ConveyerT) RegisterSeparator(
	handler func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	inp := c.getOrCreateChannel(input)
	outs := make([]chan string, len(outputs))

	for i, name := range outputs {
		outs[i] = c.getOrCreateChannel(name)
	}

	c.mutex.Lock()
	c.tasks = append(c.tasks, func(ctx context.Context) error {
		return handler(ctx, inp, outs)
	})
	c.mutex.Unlock()
}

func (c *ConveyerT) getChannel(name string) (chan string, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	channel, ok := c.channels[name]
	if ok {
		return channel, nil
	}

	return nil, ErrChanNotFound
}

func (c *ConveyerT) Send(input string, data string) error {
	ch, err := c.getChannel(input)
	if err != nil {
		return err
	}
	ch <- data

	return nil
}

func (c *ConveyerT) Recv(output string) (string, error) {
	ch, err := c.getChannel(output)
	if err != nil {
		return "", err
	}

	data, ok := <-ch
	if !ok {
		return UndefinedMsg, nil
	}

	return data, nil
}

func (c *ConveyerT) Run(ctx context.Context) error {
	defer c.closeChannels()

	errgr, ctx := errgroup.WithContext(ctx)

	c.mutex.RLock()

	for _, t := range c.tasks {
		errgr.Go(func() error {
			return t(ctx)
		})
	}

	c.mutex.RUnlock()

	err := errgr.Wait()
	if err != nil {
		return fmt.Errorf("failed to run conveyer: %w", err)
	}

	return nil
}

func (c *ConveyerT) closeChannels() {
	c.mutex.Lock()
	for _, ch := range c.channels {
		close(ch)
	}
	c.mutex.Unlock()
}
