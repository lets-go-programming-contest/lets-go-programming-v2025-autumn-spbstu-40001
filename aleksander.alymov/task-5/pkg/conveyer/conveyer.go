package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

var (
	ErrChanNotFound = errors.New("chan not found")
	ErrNoData       = errors.New("no data")
)

const Undefined = "undefined"

type conveyer struct {
	size    int
	mu      sync.RWMutex
	chans   map[string]chan string
	tasks   []func(context.Context) error
	running bool
}

func New(size int) *conveyer {
	return &conveyer{
		size:    size,
		chans:   make(map[string]chan string),
		tasks:   []func(context.Context) error{},
		running: false,
		mu:      sync.RWMutex{},
	}
}

func (c *conveyer) getOrCreateChan(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ch, ok := c.chans[name]; ok {
		return ch
	}

	ch := make(chan string, c.size)
	c.chans[name] = ch

	return ch
}

func (c *conveyer) getChan(name string) (chan string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	ch, ok := c.chans[name]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrChanNotFound, name)
	}

	return ch, nil
}

func (c *conveyer) RegisterDecorator(
	decoratorFunc func(context.Context, chan string, chan string) error,
	input, output string,
) {
	inCh := c.getOrCreateChan(input)
	outCh := c.getOrCreateChan(output)

	task := func(ctx context.Context) error {
		defer close(outCh)
		return decoratorFunc(ctx, inCh, outCh)
	}

	c.mu.Lock()
	c.tasks = append(c.tasks, task)
	c.mu.Unlock()
}

func (c *conveyer) RegisterMultiplexer(
	multiplexerFunc func(context.Context, []chan string, chan string) error,
	inputs []string,
	output string,
) {
	inputChans := make([]chan string, len(inputs))
	for i, name := range inputs {
		inputChans[i] = c.getOrCreateChan(name)
	}

	outCh := c.getOrCreateChan(output)

	task := func(ctx context.Context) error {
		defer close(outCh)
		return multiplexerFunc(ctx, inputChans, outCh)
	}

	c.mu.Lock()
	c.tasks = append(c.tasks, task)
	c.mu.Unlock()
}

func (c *conveyer) RegisterSeparator(
	separatorFunc func(context.Context, chan string, []chan string) error,
	input string,
	outputs []string,
) {
	inCh := c.getOrCreateChan(input)

	outputChans := make([]chan string, len(outputs))
	for i, name := range outputs {
		outputChans[i] = c.getOrCreateChan(name)
	}

	task := func(ctx context.Context) error {
		defer func() {
			for _, out := range outputChans {
				close(out)
			}
		}()
		return separatorFunc(ctx, inCh, outputChans)
	}

	c.mu.Lock()
	c.tasks = append(c.tasks, task)
	c.mu.Unlock()
}

func (c *conveyer) Run(ctx context.Context) error {
	c.mu.Lock()
	if c.running {
		c.mu.Unlock()
		return errors.New("conveyer already running")
	}

	c.running = true
	c.mu.Unlock()

	defer func() {
		c.mu.Lock()
		c.running = false
		c.mu.Unlock()
	}()

	errGroup, ctx := errgroup.WithContext(ctx)

	for _, task := range c.tasks {
		errGroup.Go(func() error {
			return task(ctx)
		})
	}

	if err := errGroup.Wait(); err != nil {
		return fmt.Errorf("conveyer run failed: %w", err)
	}

	return nil
}

func (c *conveyer) Send(input string, data string) error {
	channel, err := c.getChan(input)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrChanNotFound, input)
	}

	defer func() {
		_ = recover()
	}()

	channel <- data

	return nil
}

func (c *conveyer) Recv(output string) (string, error) {
	channel, err := c.getChan(output)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrChanNotFound, output)
	}

	val, ok := <-channel
	if !ok {
		return Undefined, nil
	}

	return val, nil
}
