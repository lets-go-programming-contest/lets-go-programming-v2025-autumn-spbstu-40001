package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

const (
	errSendChanNotFoundMsg = "conveyer.Send: chan not found"
	errRecvChanNotFoundMsg = "conveyer.Recv: chan not found"
)

var (
	ErrSendChanNotFound = errors.New(errSendChanNotFoundMsg)
	ErrRecvChanNotFound = errors.New(errRecvChanNotFoundMsg)
)

type Decorator func(
	context.Context,
	chan string,
	chan string,
) error

type Multiplexer func(
	context.Context,
	[]chan string,
	chan string,
) error

type Separator func(
	context.Context,
	chan string,
	[]chan string,
) error

type WorkerPool struct {
	workers []func(context.Context) error
	mu      sync.RWMutex
}

func NewWorkerPool() *WorkerPool {
	return &WorkerPool{
		workers: make([]func(context.Context) error, 0),
		mu:      sync.RWMutex{},
	}
}

func (wp *WorkerPool) Add(worker func(context.Context) error) {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	wp.workers = append(wp.workers, worker)
}

func (wp *WorkerPool) GetAll() []func(context.Context) error {
	wp.mu.RLock()
	defer wp.mu.RUnlock()

	return wp.workers
}

type ChannelRegistry struct {
	channels sync.Map
	size     int
}

func NewChannelRegistry(size int) *ChannelRegistry {
	return &ChannelRegistry{
		channels: sync.Map{},
		size:     size,
	}
}

func (cr *ChannelRegistry) GetOrCreate(name string) chan string {
	if channel, ok := cr.channels.Load(name); ok {
		if ch, typeOk := channel.(chan string); typeOk {
			return ch
		}
	}

	channel := make(chan string, cr.size)
	cr.channels.Store(name, channel)

	return channel
}

func (cr *ChannelRegistry) Get(name string) (chan string, bool) {
	channel, channelFound := cr.channels.Load(name)
	if !channelFound {
		return nil, false
	}

	ch, typeOk := channel.(chan string)

	return ch, typeOk
}

type Conveyer struct {
	channelSize int
	channels    *ChannelRegistry
	pool        *WorkerPool
}

func New(channelSize int) *Conveyer {
	return &Conveyer{
		channelSize: channelSize,
		channels:    NewChannelRegistry(channelSize),
		pool:        NewWorkerPool(),
	}
}

func (conveyer *Conveyer) RegisterDecorator(
	decorator Decorator,
	input string,
	output string,
) {
	conveyer.pool.Add(func(ctx context.Context) error {
		inputChan := conveyer.channels.GetOrCreate(input)
		outputChan := conveyer.channels.GetOrCreate(output)

		return decorator(ctx, inputChan, outputChan)
	})
}

func (conveyer *Conveyer) RegisterMultiplexer(
	multiplexer Multiplexer,
	inputs []string,
	output string,
) {
	conveyer.pool.Add(func(ctx context.Context) error {
		inputChannels := make([]chan string, len(inputs))
		for index, inputName := range inputs {
			inputChannels[index] = conveyer.channels.GetOrCreate(inputName)
		}

		outputChan := conveyer.channels.GetOrCreate(output)

		return multiplexer(ctx, inputChannels, outputChan)
	})
}

func (conveyer *Conveyer) RegisterSeparator(
	separator Separator,
	input string,
	outputs []string,
) {
	conveyer.pool.Add(func(ctx context.Context) error {
		inputChan := conveyer.channels.GetOrCreate(input)

		outputChannels := make([]chan string, len(outputs))
		for index, outputName := range outputs {
			outputChannels[index] = conveyer.channels.GetOrCreate(outputName)
		}

		return separator(ctx, inputChan, outputChannels)
	})
}

func (conveyer *Conveyer) Run(ctx context.Context) error {
	group, ctxWithErrs := errgroup.WithContext(ctx)
	workers := conveyer.pool.GetAll()

	for _, worker := range workers {
		group.Go(func() error {
			return worker(ctxWithErrs)
		})
	}

	if err := group.Wait(); err != nil {
		return fmt.Errorf("failed to run workers: %w", err)
	}

	return nil
}

func (conveyer *Conveyer) Send(input string, data string) error {
	channel := conveyer.channels.GetOrCreate(input)

	channel <- data

	return nil
}

func (conveyer *Conveyer) Recv(output string) (string, error) {
	channel, channelFound := conveyer.channels.Get(output)
	if !channelFound {
		return "", ErrRecvChanNotFound
	}

	data, ok := <-channel
	if !ok {
		return "undefined", nil
	}

	return data, nil
}
