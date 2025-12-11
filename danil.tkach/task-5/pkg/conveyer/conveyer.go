package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

const Undefined = "undefined"

var ErrChannelNotFound = errors.New("chan not found")

type Task func(context.Context) error

type ConveyerImpl struct {
	mu       sync.RWMutex
	channels map[string]chan string
	tasks    []Task
	chanSize int
}

func New(size int) *ConveyerImpl {
	return &ConveyerImpl{
		mu:       sync.RWMutex{},
		channels: make(map[string]chan string),
		chanSize: size,
		tasks:    make([]Task, 0),
	}
}

func (c *ConveyerImpl) getOrCreateChannel(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ch, exists := c.channels[name]; exists {
		return ch
	}

	ch := make(chan string, c.chanSize)
	c.channels[name] = ch

	return ch
}

func (c *ConveyerImpl) getChannel(name string) (chan string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	ch, exists := c.channels[name]
	if !exists {
		return nil, ErrChannelNotFound
	}

	return ch, nil
}

func (c *ConveyerImpl) RegisterDecorator(
	function func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	inCh := c.getOrCreateChannel(input)
	outCh := c.getOrCreateChannel(output)

	task := func(ctx context.Context) error {
		return function(ctx, inCh, outCh)
	}

	c.mu.Lock()
	c.tasks = append(c.tasks, task)
	c.mu.Unlock()
}

func (c *ConveyerImpl) RegisterMultiplexer(
	function func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	inputChannels := make([]chan string, 0, len(inputs))
	for _, name := range inputs {
		inputChannels = append(inputChannels, c.getOrCreateChannel(name))
	}

	outCh := c.getOrCreateChannel(output)

	task := func(ctx context.Context) error {
		return function(ctx, inputChannels, outCh)
	}

	c.mu.Lock()
	c.tasks = append(c.tasks, task)
	c.mu.Unlock()
}

func (c *ConveyerImpl) RegisterSeparator(
	function func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	inCh := c.getOrCreateChannel(input)

	outputChannels := make([]chan string, 0, len(outputs))
	for _, name := range outputs {
		outputChannels = append(outputChannels, c.getOrCreateChannel(name))
	}

	task := func(ctx context.Context) error {
		return function(ctx, inCh, outputChannels)
	}

	c.mu.Lock()
	c.tasks = append(c.tasks, task)
	c.mu.Unlock()
}

func (c *ConveyerImpl) Run(ctx context.Context) error {
	defer func() {
		c.mu.Lock()
		defer c.mu.Unlock()

		for _, ch := range c.channels {
			close(ch)
		}
	}()

	c.mu.RLock()

	group, groupCtx := errgroup.WithContext(ctx)

	for _, task := range c.tasks {
		currentTask := task

		group.Go(func() error {
			return currentTask(groupCtx)
		})
	}

	c.mu.RUnlock()

	if err := group.Wait(); err != nil {
		return fmt.Errorf("conveyer finished with error: %w", err)
	}

	return nil
}

func (c *ConveyerImpl) Send(input string, data string) error {
	targetChannel, err := c.getChannel(input)
	if err != nil {
		return err
	}

	defer func() {
		_ = recover()
	}()

	targetChannel <- data

	return nil
}

func (c *ConveyerImpl) Recv(output string) (string, error) {
	ch, err := c.getChannel(output)
	if err != nil {
		return "", err
	}

	val, ok := <-ch
	if !ok {
		return Undefined, nil
	}

	return val, nil
}
