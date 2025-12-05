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

type Conveyer struct {
	mu           sync.RWMutex
	channels     map[string]chan string
	workers      []func(context.Context) error
	chanCapacity int
}

func New(size int) *Conveyer {
	return &Conveyer{
		mu:           sync.RWMutex{},
		channels:     make(map[string]chan string),
		workers:      []func(context.Context) error{},
		chanCapacity: size,
	}
}

func (c *Conveyer) getOrCreateChan(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()

	tempChan, exists := c.channels[name]
	if exists {
		return tempChan
	}

	tempChan = make(chan string, c.chanCapacity)
	c.channels[name] = tempChan

	return tempChan
}

func (c *Conveyer) getChannel(name string) (chan string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	tempChan, exists := c.channels[name]

	return tempChan, exists
}

func (c *Conveyer) RegisterDecorator(
	funkt func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	inCh := c.getOrCreateChan(input)
	outCh := c.getOrCreateChan(output)

	task := func(ctx context.Context) error {
		defer close(outCh)

		return funkt(ctx, inCh, outCh)
	}
	c.workers = append(c.workers, task)
}

func (c *Conveyer) RegisterMultiplexer(
	funkt func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	inputChannels := make([]chan string, 0, len(inputs))
	for _, name := range inputs {
		inputChannels = append(inputChannels, c.getOrCreateChan(name))
	}

	outCh := c.getOrCreateChan(output)

	task := func(ctx context.Context) error {
		defer close(outCh)

		return funkt(ctx, inputChannels, outCh)
	}
	c.workers = append(c.workers, task)
}

func (c *Conveyer) RegisterSeparator(
	funkt func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	inCh := c.getOrCreateChan(input)

	outputChannels := make([]chan string, 0, len(outputs))
	for _, name := range outputs {
		outputChannels = append(outputChannels, c.getOrCreateChan(name))
	}

	task := func(ctx context.Context) error {
		defer func() {
			for _, ch := range outputChannels {
				close(ch)
			}
		}()

		return funkt(ctx, inCh, outputChannels)
	}
	c.workers = append(c.workers, task)
}

func (c *Conveyer) Run(ctx context.Context) error {
	c.mu.RLock()
	workers := c.workers
	c.mu.RUnlock()

	group, gctx := errgroup.WithContext(ctx)

	for i := range workers {
		job := workers[i]

		group.Go(func() error {
			return job(gctx)
		})
	}

	if err := group.Wait(); err != nil {
		return fmt.Errorf("conveyer run failed: %w", err)
	}

	return nil
}

func (c *Conveyer) Send(input, data string) error {
	channel, exists := c.getChannel(input)
	if !exists {
		return ErrChannelNotFound
	}

	defer func() { _ = recover() }()
	channel <- data

	return nil
}

func (c *Conveyer) Recv(output string) (string, error) {
	channel, exists := c.getChannel(output)
	if !exists {
		return "", ErrChannelNotFound
	}

	val, ok := <-channel
	if !ok {
		return Undefined, nil
	}

	return val, nil
}
