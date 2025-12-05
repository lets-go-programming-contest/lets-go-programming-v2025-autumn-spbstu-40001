package conveyer

import (
	"context"
	"errors"
	"sync"
)

type conveyerImpl struct {
	size     int
	channels map[string]chan string
	handlers []handler
	mu       sync.RWMutex
}

type handler struct {
	run func(ctx context.Context) error
}

func New(size int) *conveyerImpl {
	return &conveyerImpl{
		size:     size,
		channels: make(map[string]chan string),
		handlers: make([]handler, 0),
	}
}

func (c *conveyerImpl) RegisterDecorator(
	fn func(ctx context.Context, input chan string, output chan string) error,
	inputName, outputName string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	inputChan := c.getOrCreateChannel(inputName)
	outputChan := c.getOrCreateChannel(outputName)

	c.handlers = append(c.handlers, handler{
		run: func(ctx context.Context) error {
			return fn(ctx, inputChan, outputChan)
		},
	})
}

func (c *conveyerImpl) RegisterMultiplexer(
	fn func(ctx context.Context, inputs []chan string, output chan string) error,
	inputNames []string, outputName string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	inputs := make([]chan string, len(inputNames))
	for i, name := range inputNames {
		inputs[i] = c.getOrCreateChannel(name)
	}
	outputChan := c.getOrCreateChannel(outputName)

	c.handlers = append(c.handlers, handler{
		run: func(ctx context.Context) error {
			return fn(ctx, inputs, outputChan)
		},
	})
}

func (c *conveyerImpl) RegisterSeparator(
	fn func(ctx context.Context, input chan string, outputs []chan string) error,
	inputName string, outputNames []string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	inputChan := c.getOrCreateChannel(inputName)
	outputs := make([]chan string, len(outputNames))
	for i, name := range outputNames {
		outputs[i] = c.getOrCreateChannel(name)
	}

	c.handlers = append(c.handlers, handler{
		run: func(ctx context.Context) error {
			return fn(ctx, inputChan, outputs)
		},
	})
}

func (c *conveyerImpl) Run(ctx context.Context) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(c.handlers))

	for _, h := range c.handlers {
		wg.Add(1)
		go func(h handler) {
			defer wg.Done()
			if err := h.run(ctx); err != nil {
				select {
				case errCh <- err:
				default:
				}
			}
		}(h)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	select {
	case <-ctx.Done():
		c.closeAllChannels()
		return ctx.Err()
	case err := <-errCh:
		c.closeAllChannels()
		return err
	}
}

func (c *conveyerImpl) Send(inputName string, data string) error {
	c.mu.RLock()
	ch, exists := c.channels[inputName]
	c.mu.RUnlock()

	if !exists {
		return errors.New("chan not found")
	}

	select {
	case ch <- data:
		return nil
	default:
		return errors.New("channel is full")
	}
}

func (c *conveyerImpl) Recv(outputName string) (string, error) {
	c.mu.RLock()
	ch, exists := c.channels[outputName]
	c.mu.RUnlock()

	if !exists {
		return "", errors.New("chan not found")
	}

	val, ok := <-ch
	if !ok {
		return "undefined", nil
	}

	return val, nil
}

func (c *conveyerImpl) getOrCreateChannel(name string) chan string {
	if ch, exists := c.channels[name]; exists {
		return ch
	}
	ch := make(chan string, c.size)
	c.channels[name] = ch
	return ch
}

func (c *conveyerImpl) closeAllChannels() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for name, ch := range c.channels {
		close(ch)
		delete(c.channels, name)
	}
}
