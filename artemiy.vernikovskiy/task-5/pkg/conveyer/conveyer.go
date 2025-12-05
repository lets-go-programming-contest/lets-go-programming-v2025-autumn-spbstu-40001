package conveyer

import (
	"context"
	"errors"
	"sync"
)

// Errors. Simple enough.
var (
	ErrChanNotFound     = errors.New("chan not found")
	ErrPipelineCanceled = errors.New("pipeline run canceled")
)

// The whole basis structure. Used for gorutines, you know.
type pipeline struct {
	mu       sync.RWMutex
	channels map[string]chan string
	handlers []func(ctx context.Context) error
	size     int
}

// Construct (god famn this golang without classes).
func New(size int) *pipeline {
	return &pipeline{
		mu:       sync.RWMutex{},
		channels: make(map[string]chan string),
		handlers: make([]func(ctx context.Context) error, 0),
		size:     size,
	}
}

// Registrate Decorator (because task needs it).
func (p *pipeline) RegisterDecorator(
	workingFunc func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	inCh := p.ensureChannel(input)
	outCh := p.ensureChannel(output)
	p.mu.Lock()
	p.handlers = append(p.handlers, func(ctx context.Context) error {
		return workingFunc(ctx, inCh, outCh)
	})
	p.mu.Unlock()
}

// Registrate Multiplexer (because task needs it).
func (p *pipeline) RegisterMultiplexer(
	workingFunc func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	inChans := make([]chan string, 0, len(inputs))
	for _, name := range inputs {
		inChans = append(inChans, p.ensureChannel(name))
	}

	outCh := p.ensureChannel(output)
	p.mu.Lock()
	p.handlers = append(p.handlers, func(ctx context.Context) error {
		return workingFunc(ctx, inChans, outCh)
	})
	p.mu.Unlock()
}

// Registrate Separator (because task needs it).
func (p *pipeline) RegisterSeparator(
	workingFunc func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	inCh := p.ensureChannel(input)

	outChans := make([]chan string, 0, len(outputs))
	for _, name := range outputs {
		outChans = append(outChans, p.ensureChannel(name))
	}

	p.mu.Lock()
	p.handlers = append(p.handlers, func(ctx context.Context) error {
		return workingFunc(ctx, inCh, outChans)
	})
	p.mu.Unlock()
}

// Here we run pipeline with gorutines. Main logic.
func (p *pipeline) Run(ctx context.Context) error {
	var pipelineWaitGroup sync.WaitGroup

	errCh := make(chan error, len(p.handlers))

	for _, handler := range p.handlers {
		pipelineWaitGroup.Add(1)

		go func(h func(ctx context.Context) error) {
			defer pipelineWaitGroup.Done()

			err := h(ctx)
			if err != nil {
				select {
				case errCh <- err:
				case <-ctx.Done():
				}
			}
		}(handler)
	}

	done := make(chan struct{})

	go func() {
		pipelineWaitGroup.Wait()
		close(done)
	}()

	select {
	case err := <-errCh:
		return err
	case <-done:
		return nil
		//	case <-ctx.Done():
		//		return errors.Join(ErrPipelineCanceled, ctx.Err())
	}
}

// Work with pipeline indirectly.
func (p *pipeline) Send(name string, data string) error {
	p.mu.RLock()
	channel, ok := p.channels[name]
	p.mu.RUnlock()

	if !ok {
		return ErrChanNotFound
	}

	channel <- data

	return nil
}

// Get pipeline info indirectly.
func (p *pipeline) Recv(name string) (string, error) {
	p.mu.RLock()
	channel, isThisOKLongEnogh := p.channels[name]
	p.mu.RUnlock()

	if !isThisOKLongEnogh {
		return "", ErrChanNotFound
	}

	// okay, lets do without select
	val, ok := <-channel
	if !ok {
		return "undefined", nil
	}

	return val, nil
}

// Get channel properly.
func (p *pipeline) ensureChannel(name string) chan string {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.channels[name]; !ok {
		p.channels[name] = make(chan string, p.size)
	}

	return p.channels[name]
}
