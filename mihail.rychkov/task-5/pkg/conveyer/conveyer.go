package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Conveyer[T any] struct {
	channelCapacity int
	pipes           map[string]chan T
	nodes           []func(c context.Context) error
	mutex           sync.RWMutex
}

var (
	ErrChannelNotFound   = errors.New("chan not found")
	ErrClosedChanelEmpty = errors.New("requested channel was closed and is empty")
)

func NewConveyer[T any](channelCapacity int) Conveyer[T] {
	return Conveyer[T]{channelCapacity, make(map[string]chan T), []func(ctx context.Context) error{}, sync.RWMutex{}}
}

func (obj *Conveyer[T]) reserveChannel(name string) chan T {
	channel, exists := obj.pipes[name]
	if exists {
		return channel
	}

	channel = make(chan T, obj.channelCapacity)
	obj.pipes[name] = channel

	return channel
}

func (obj *Conveyer[T]) Run(ctx context.Context) error {
	defer func() {
		obj.mutex.RLock()
		defer obj.mutex.RUnlock()

		for _, channel := range obj.pipes {
			close(channel)
		}
	}()

	obj.mutex.RLock()

	group, ctx := errgroup.WithContext(ctx)
	for _, functor := range obj.nodes {
		group.Go(func() error { return functor(ctx) })
	}

	obj.mutex.RUnlock()

	err := group.Wait()
	if err != nil {
		return fmt.Errorf("Conveyer finished with error: %w", err)
	}

	return nil
}

func (obj *Conveyer[T]) Send(inChName string, data T) error {
	obj.mutex.RLock()
	channel, exists := obj.pipes[inChName]
	obj.mutex.RUnlock()

	if !exists {
		return ErrChannelNotFound
	}

	channel <- data

	return nil
}

func (obj *Conveyer[T]) Recv(outChName string) (T, error) {
	obj.mutex.RLock()
	channel, exists := obj.pipes[outChName]
	obj.mutex.RUnlock()

	if !exists {
		var res T

		return res, ErrChannelNotFound
	}

	res, ok := <-channel
	if !ok {
		return res, ErrClosedChanelEmpty
	}

	return res, nil
}

func (obj *Conveyer[T]) RegisterDecorator(
	functor func(c context.Context, input chan T, output chan T) error,
	input string, output string,
) {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()

	in := obj.reserveChannel(input)
	out := obj.reserveChannel(output)
	obj.nodes = append(obj.nodes, func(c context.Context) error {
		return functor(c, in, out)
	})
}

func (obj *Conveyer[T]) RegisterMultiplexer(
	functor func(c context.Context, input []chan T, output chan T) error,
	input []string, output string,
) {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()

	inputs := make([]chan T, len(input))
	for idx, name := range input {
		inputs[idx] = obj.reserveChannel(name)
	}

	out := obj.reserveChannel(output)
	obj.nodes = append(obj.nodes, func(c context.Context) error {
		return functor(c, inputs, out)
	})
}

func (obj *Conveyer[T]) RegisterSeparator(
	functor func(c context.Context, input chan T, output []chan T) error,
	input string, output []string,
) {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()

	inCh := obj.reserveChannel(input)

	outputs := make([]chan T, len(output))
	for idx, name := range output {
		outputs[idx] = obj.reserveChannel(name)
	}

	obj.nodes = append(obj.nodes, func(c context.Context) error {
		return functor(c, inCh, outputs)
	})
}
