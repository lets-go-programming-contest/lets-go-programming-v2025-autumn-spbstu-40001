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

type Worker func(context.Context) error

type Conveyer struct {
	mu           sync.RWMutex
	channels     map[string]chan string
	workers      []Worker
	chanCapacity int
}

func New(size int) *Conveyer {
	return &Conveyer{
		mu:           sync.RWMutex{},
		channels:     make(map[string]chan string),
		workers:      make([]Worker, 0),
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

func (c *Conveyer) getChannel(name string) (chan string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	tempChan, exists := c.channels[name]
	if exists {
		return tempChan, nil
	}

	return nil, ErrChannelNotFound
}

func (c *Conveyer) RegisterDecorator(
	funct func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	inCh := c.getOrCreateChan(input)
	outCh := c.getOrCreateChan(output)

	c.mu.Lock()
	c.workers = append(c.workers, func(ctx context.Context) error {
		return funct(ctx, inCh, outCh)
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
		inChans = append(inChans, c.getOrCreateChan(name))
	}

	outCh := c.getOrCreateChan(output)

	c.mu.Lock()
	c.workers = append(c.workers, func(ctx context.Context) error {
		return funct(ctx, inChans, outCh)
	})
	c.mu.Unlock()
}

func (c *Conveyer) RegisterSeparator(
	funct func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	inCh := c.getOrCreateChan(input)

	outChans := make([]chan string, 0, len(outputs))
	for _, name := range outputs {
		outChans = append(outChans, c.getOrCreateChan(name))
	}

	c.mu.Lock()
	c.workers = append(c.workers, func(ctx context.Context) error {
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
	channel, err := c.getChannel(input)
	if err != nil {
		return err
	}

	defer func() {
		_ = recover()
	}()

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
		return Undefined, nil
	}

	return val, nil
}
