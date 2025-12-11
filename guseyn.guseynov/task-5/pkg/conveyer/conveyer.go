package conveyer

import (
	"context"
	"errors"
	"sync"

	"golang.org/x/sync/errgroup"
)

const (
	ErrSendChanNotFound = "conveyer.Send: chan not found"
	ErrRecvChanNotFound = "conveyer.Recv: chan not found"
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
		size: size,
	}
}

func (cr *ChannelRegistry) GetOrCreate(name string) chan string {
	if ch, ok := cr.channels.Load(name); ok {
		return ch.(chan string)
	}

	ch := make(chan string, cr.size)
	cr.channels.Store(name, ch)
	return ch
}

func (cr *ChannelRegistry) Get(name string) (chan string, bool) {
	ch, ok := cr.channels.Load(name)
	if !ok {
		return nil, false
	}
	return ch.(chan string), true
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
		for i, inputName := range inputs {
			inputChannels[i] = conveyer.channels.GetOrCreate(inputName)
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
		for i, outputName := range outputs {
			outputChannels[i] = conveyer.channels.GetOrCreate(outputName)
		}
		return separator(ctx, inputChan, outputChannels)
	})
}

func (conveyer *Conveyer) Run(context context.Context) error {
	group, contextWithErrs := errgroup.WithContext(context)
	workers := conveyer.pool.GetAll()

	for _, worker := range workers {
		worker := worker
		group.Go(func() error {
			return worker(contextWithErrs)
		})
	}

	return group.Wait()
}

func (conveyer *Conveyer) Send(input string, data string) error {
	channel, ok := conveyer.channels.Get(input)
	if !ok {
		return errors.New(ErrSendChanNotFound)
	}

	channel <- data
	return nil
}

func (conveyer *Conveyer) Recv(output string) (string, error) {
	channel, ok := conveyer.channels.Get(output)
	if !ok {
		return "", errors.New(ErrRecvChanNotFound)
	}

	data, ok := <-channel
	if !ok {
		return "undefined", nil
	}

	return data, nil
}
