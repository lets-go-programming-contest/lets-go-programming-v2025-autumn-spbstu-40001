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
	ErrChannelFull  = errors.New("channel is full")
	ErrNoData       = errors.New("no data available in channel")
)

const (
	UndefinedValue = "undefined"
)

type Handler func(context.Context) error

type conveyerImpl struct {
	size     int
	channels map[string]chan string
	handlers []Handler
	mu       sync.RWMutex
}

func New(size int) *conveyerImpl {
	return &conveyerImpl{
		size:     size,
		channels: make(map[string]chan string),
		handlers: make([]Handler, 0),
		mu:       sync.RWMutex{},
	}
}

func (c *conveyerImpl) getChannel(name string) (chan string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	ch, exists := c.channels[name]

	return ch, exists
}

func (c *conveyerImpl) getOrCreateChannel(name string) chan string {
	if ch, exists := c.channels[name]; exists {
		return ch
	}

	ch := make(chan string, c.size)
	c.channels[name] = ch

	return ch
}

func (c *conveyerImpl) RegisterDecorator(
	decoratorFunc func(ctx context.Context, input chan string, output chan string) error,
	inputName, outputName string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	inputChan := c.getOrCreateChannel(inputName)
	outputChan := c.getOrCreateChannel(outputName)

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return decoratorFunc(ctx, inputChan, outputChan)
	})
}

func (c *conveyerImpl) RegisterMultiplexer(
	multiplexerFunc func(ctx context.Context, inputs []chan string, output chan string) error,
	inputNames []string, outputName string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	inputs := make([]chan string, len(inputNames))
	for i, name := range inputNames {
		inputs[i] = c.getOrCreateChannel(name)
	}

	outputChan := c.getOrCreateChannel(outputName)

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return multiplexerFunc(ctx, inputs, outputChan)
	})
}

func (c *conveyerImpl) RegisterSeparator(
	separatorFunc func(ctx context.Context, input chan string, outputs []chan string) error,
	inputName string, outputNames []string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	inputChan := c.getOrCreateChannel(inputName)
	outputs := make([]chan string, len(outputNames))

	for i, name := range outputNames {
		outputs[i] = c.getOrCreateChannel(name)
	}

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return separatorFunc(ctx, inputChan, outputs)
	})
}

func (c *conveyerImpl) Run(ctx context.Context) error {
	errorGroup, groupContext := errgroup.WithContext(ctx)

	for _, handler := range c.handlers {
		currentHandler := handler

		errorGroup.Go(func() error {
			return currentHandler(groupContext)
		})
	}

	if err := errorGroup.Wait(); err != nil {
		return fmt.Errorf("conveyer run: %w", err)
	}

	return nil
}

func (c *conveyerImpl) Send(inputName string, data string) error {
	channel, exists := c.getChannel(inputName)
	if !exists {
		return fmt.Errorf("%s: %w", inputName, ErrChanNotFound)
	}

	select {
	case channel <- data:
		return nil
	default:
		return fmt.Errorf("%s: %w", inputName, ErrChannelFull)
	}
}

func (c *conveyerImpl) Recv(outputName string) (string, error) {
	channel, exists := c.getChannel(outputName)
	if !exists {
		return "", fmt.Errorf("%s: %w", outputName, ErrChanNotFound)
	}

	select {
	case val, ok := <-channel:
		if !ok {
			return UndefinedValue, nil
		}

		return val, nil
	default:
		return "", fmt.Errorf("%s: %w", outputName, ErrNoData)
	}
}
