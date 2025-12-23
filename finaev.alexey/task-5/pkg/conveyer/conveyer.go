package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Handler func(context.Context) error

type Conveyer struct {
	mu       sync.RWMutex
	channels map[string]chan string
	size     int
	handlers []Handler
}

func New(size int) *Conveyer {
	return &Conveyer{
		channels: make(map[string]chan string),
		size:     size,
		handlers: make([]Handler, 0),
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
	defer c.mu.Unlock()

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		defer func() {
			c.mu.Lock()
			defer c.mu.Unlock()
			if ch, exists := c.channels[output]; exists {
				close(ch)
				delete(c.channels, output)
			}
		}()

		return funct(ctx, inputChan, outputChan)
	})
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
	defer c.mu.Unlock()

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		defer func() {
			c.mu.Lock()
			defer c.mu.Unlock()
			if ch, exists := c.channels[output]; exists {
				close(ch)
				delete(c.channels, output)
			}
		}()

		return funct(ctx, inChans, outCh)
	})
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
	defer c.mu.Unlock()

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		defer func() {
			c.mu.Lock()
			defer c.mu.Unlock()
			for _, name := range outputs {
				if ch, exists := c.channels[name]; exists {
					close(ch)
					delete(c.channels, name)
				}
			}
		}()

		return funct(ctx, inCh, outChans)
	})
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
	handlers := c.handlers
	c.mu.RUnlock()

	group, gctx := errgroup.WithContext(ctx)

	for i := range handlers {
		job := handlers[i]

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

	select {
	case channel <- data:
		return nil
	default:
		return errors.New("channel is full or closed")
	}
}

func (c *Conveyer) Recv(output string) (string, error) {
	channel, err := c.getChannel(output)
	if err != nil {
		return "", err
	}

	val, ok := <-channel
	if !ok {
		return "", errors.New("channel closed")
	}

	return val, nil
}
