package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

var errChanNotFound = errors.New("chan not found")

const errUndefinedStr = "undefined"

type Conveyer struct {
	chansSize int
	chansMap  map[string]chan string
	mutex     sync.RWMutex
	modifiers []func(context.Context) error
}

func New(size int) *Conveyer {
	return &Conveyer{
		chansSize: size,
		chansMap:  make(map[string]chan string),
		mutex:     sync.RWMutex{},
		modifiers: []func(context.Context) error{},
	}
}

func (c *Conveyer) getChannel(name string) (chan string, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	ch, ok := c.chansMap[name]
	if !ok {
		return nil, fmt.Errorf("%w: %s", errChanNotFound, name)
	}

	return ch, nil
}

func (c *Conveyer) getOrCreateChannel(name string) chan string {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if ch, exists := c.chansMap[name]; exists {
		return ch
	}

	ch := make(chan string, c.chansSize)
	c.chansMap[name] = ch

	return ch
}

func (c *Conveyer) RegisterDecorator(
	newDecorator func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	inputChan := c.getOrCreateChannel(input)
	outputChan := c.getOrCreateChannel(output)

	newModifier := func(ctx context.Context) error {
		defer close(outputChan)

		return newDecorator(ctx, inputChan, outputChan)
	}

	c.modifiers = append(c.modifiers, newModifier)
}

func (c *Conveyer) RegisterMultiplexer(
	newMultiplexer func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	inputChans := make([]chan string, 0, len(inputs))
	for _, input := range inputs {
		inputChans = append(inputChans, c.getOrCreateChannel(input))
	}

	outputChan := c.getOrCreateChannel(output)

	newModifier := func(ctx context.Context) error {
		defer close(outputChan)

		return newMultiplexer(ctx, inputChans, outputChan)
	}

	c.modifiers = append(c.modifiers, newModifier)
}

func (c *Conveyer) RegisterSeparator(
	newSeparator func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	inputChan := c.getOrCreateChannel(input)

	outputChans := make([]chan string, 0, len(outputs))
	for _, output := range outputs {
		outputChans = append(outputChans, c.getOrCreateChannel(output))
	}

	newModifier := func(ctx context.Context) error {
		defer func() {
			for _, channel := range outputChans {
				close(channel)
			}
		}()

		return newSeparator(ctx, inputChan, outputChans)
	}

	c.modifiers = append(c.modifiers, newModifier)
}

func (c *Conveyer) Run(ctx context.Context) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	errGroup, ctx := errgroup.WithContext(ctx)

	for _, modifier := range c.modifiers {
		errGroup.Go(func() error {
			return modifier(ctx)
		})
	}

	if err := errGroup.Wait(); err != nil {
		return fmt.Errorf("conveyer finished with error: %w", err)
	}

	return nil
}

func (c *Conveyer) Send(input string, data string) error {
	channel, err := c.getChannel(input)
	if err != nil {
		return fmt.Errorf("%w: %s", errChanNotFound, input)
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
		return "", fmt.Errorf("%w: %s", errChanNotFound, output)
	}

	val, ok := <-channel
	if !ok {
		return errUndefinedStr, nil
	}

	return val, nil
}
