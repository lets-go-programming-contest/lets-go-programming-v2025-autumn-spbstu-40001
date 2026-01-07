package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

const UndefinedMsg = "undefined"

var ErrChanelNotFound = errors.New("chan not found")

type Handler func(ctx context.Context) error

type Conveyer struct {
	channels map[string]chan string
	size     int
	mutex    sync.RWMutex
	handlers []Handler
}

func New(size int) Conveyer {
	return Conveyer{
		size:     size,
		channels: make(map[string]chan string),
		mutex:    sync.RWMutex{},
		handlers: make([]Handler, 0),
	}
}

func (obj *Conveyer) RegisterDecorator(
	handler func(
		ctx context.Context,
		input chan string,
		output chan string,
	) error,
	input string,
	output string,
) {
	inp := obj.createChanel(input)
	out := obj.createChanel(output)

	obj.handlers = append(obj.handlers, func(ctx context.Context) error {
		return handler(ctx, inp, out)
	})
}

func (obj *Conveyer) RegisterMultiplexer(
	handler func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()

	ins := make([]chan string, len(inputs))
	for idx, name := range inputs {
		ins[idx] = obj.createChanel(name)
	}

	out := obj.createChanel(output)

	obj.handlers = append(obj.handlers, func(ctx context.Context) error {
		return handler(ctx, ins, out)
	})
}

func (obj *Conveyer) RegisterSeparator(
	handler func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()

	inp := obj.createChanel(input)

	outs := make([]chan string, len(outputs))
	for idx, name := range outputs {
		outs[idx] = obj.createChanel(name)
	}

	obj.handlers = append(obj.handlers, func(ctx context.Context) error {
		return handler(ctx, inp, outs)
	})
}

func (obj *Conveyer) Run(ctx context.Context) error {
	defer obj.closeAllChannels()

	obj.mutex.RLock()
	defer obj.mutex.RUnlock()

	group, ctx := errgroup.WithContext(ctx)
	for _, handler := range obj.handlers {
		group.Go(func() error {
			return handler(ctx)
		})
	}

	err := group.Wait()
	if err != nil {
		return fmt.Errorf("failed to run conveyer: %w", err)
	}

	return nil
}

func (obj *Conveyer) Send(input string, data string) error {
	ch, err := obj.getChanel(input)
	if err != nil {
		return err
	}

	ch <- data

	return nil
}

func (obj *Conveyer) Recv(output string) (string, error) {
	ch, err := obj.getChanel(output)
	if err != nil {
		return "", err
	}

	data, ok := <-ch
	if !ok {
		return UndefinedMsg, nil
	}

	return data, nil
}

func (obj *Conveyer) createChanel(name string) chan string {
	chanel, exists := obj.channels[name]
	if exists {
		return chanel
	}

	chanel = make(chan string, obj.size)
	obj.channels[name] = chanel

	return chanel
}

func (obj *Conveyer) getChanel(name string) (chan string, error) {
	chanel, exists := obj.channels[name]
	if exists {
		return chanel, nil
	}

	return nil, ErrChanelNotFound
}

func (obj *Conveyer) closeAllChannels() {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()

	for _, ch := range obj.channels {
		close(ch)
	}
}
