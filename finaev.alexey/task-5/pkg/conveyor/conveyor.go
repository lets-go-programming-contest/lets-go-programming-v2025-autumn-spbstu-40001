package conveyor

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Conveyer struct {
	mu       sync.RWMutex
	channels map[string]chan string
	size     int
	handlers []func(context.Context) error
}

func New(size int) *Conveyer {
	return &Conveyer{
		mu:       sync.RWMutex{},
		channels: make(map[string]chan string),
		size:     size,
		handlers: []func(context.Context) error{},
	}
}

func (c *Conveyer) getOrCreateChannel(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ch, exists := c.channels[name]; exists {
		return ch
	}

	ch := make(chan string, c.size)
	c.channels[name] = ch

	return ch
}

func (c *Conveyer) getChannel(name string) (chan string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	ch, exists := c.channels[name]
	if !exists {
		return nil, errors.New("chan not found")
	}

	return ch, nil
}

func (c *Conveyer) RegisterDecorator(
	funct func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	inputChan := c.getOrCreateChannel(input)
	outputChan := c.getOrCreateChannel(output)

	c.mu.Lock()
	c.handlers = append(c.handlers, func(ctx context.Context) error {
		defer close(outputChan)

		return funct(ctx, inputChan, outputChan)
	})
	c.mu.Unlock()
}

func (c *Conveyer) RegisterMultiplexer(
	funct func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	inChans := make([]chan string, 0, len(inputs))
	for _, name := range inputs {
		inChans = append(inChans, c.getOrCreateChannel(name))
	}

	outCh := c.getOrCreateChannel(output)

	c.mu.Lock()
	c.handlers = append(c.handlers, func(ctx context.Context) error {
		defer close(outCh)

		return funct(ctx, inChans, outCh)
	})
	c.mu.Unlock()
}

func (c *Conveyer) RegisterSeparator(
	funct func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	inCh := c.getOrCreateChannel(input)

	outChans := make([]chan string, 0, len(outputs))
	for _, name := range outputs {
		outChans = append(outChans, c.getOrCreateChannel(name))
	}

	c.mu.Lock()
	c.handlers = append(c.handlers, func(ctx context.Context) error {
		defer func() {
			for _, ch := range outChans {
				close(ch)
			}
		}()

		return funct(ctx, inCh, outChans)
	})
	c.mu.Unlock()
}

func (c *Conveyer) Run(ctx context.Context) error {
	defer func() {
		c.mu.Lock()
		for _, ch := range c.channels {
			close(ch)
		}
		c.mu.Unlock()
	}()

	c.mu.RLock()
	workers := c.handlers
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
	channel, err := c.getChannel(input)
	if err != nil {
		return err
	}
	channel <- data

	return nil
}

func (c *Conveyer) Recv(output string) (string, error) {
	channel, err := c.getChannel(output)
	if err != nil {
		return "", err
	}

	val, ok := <-channel
	if !ok {
		return "undefined", nil
	}

	return val, nil
}
