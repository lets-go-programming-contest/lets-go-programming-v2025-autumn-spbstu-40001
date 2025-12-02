package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

const Undefined = "undefined"

var ErrChannelMissing = errors.New("chan not found")

type Pipeline struct {
	lock            sync.Mutex
	chans           map[string]chan string
	jobs            []func(context.Context) error
	channelCapacity int
}

func New(capacity int) *Pipeline {
	return &Pipeline{
		chans:           make(map[string]chan string),
		jobs:            []func(context.Context) error{},
		channelCapacity: capacity,
	}
}

func (p *Pipeline) getOrCreateChan(name string) chan string {
	p.lock.Lock()
	defer p.lock.Unlock()

	if ch, exists := p.chans[name]; exists {
		return ch
	}

	ch := make(chan string, p.channelCapacity)
	p.chans[name] = ch
	return ch
}

func (p *Pipeline) getChan(name string) (chan string, bool) {
	p.lock.Lock()
	defer p.lock.Unlock()

	ch, exists := p.chans[name]
	return ch, exists
}

func (p *Pipeline) RegisterDecorator(
	fn func(ctx context.Context, input chan string, output chan string) error,
	inputName, outputName string,
) {
	inCh := p.getOrCreateChan(inputName)
	outCh := p.getOrCreateChan(outputName)

	job := func(ctx context.Context) error {
		defer close(outCh)
		return fn(ctx, inCh, outCh)
	}

	p.lock.Lock()
	p.jobs = append(p.jobs, job)
	p.lock.Unlock()
}

func (p *Pipeline) RegisterMultiplexer(
	fn func(ctx context.Context, inputs []chan string, output chan string) error,
	inputNames []string,
	outputName string,
) {
	inChans := make([]chan string, 0, len(inputNames))
	for _, n := range inputNames {
		inChans = append(inChans, p.getOrCreateChan(n))
	}

	outCh := p.getOrCreateChan(outputName)

	job := func(ctx context.Context) error {
		defer close(outCh)
		return fn(ctx, inChans, outCh)
	}

	p.lock.Lock()
	p.jobs = append(p.jobs, job)
	p.lock.Unlock()
}

func (p *Pipeline) RegisterSeparator(
	fn func(ctx context.Context, input chan string, outputs []chan string) error,
	inputName string,
	outputNames []string,
) {
	inCh := p.getOrCreateChan(inputName)
	outChans := make([]chan string, 0, len(outputNames))
	for _, n := range outputNames {
		outChans = append(outChans, p.getOrCreateChan(n))
	}

	job := func(ctx context.Context) error {
		defer func() {
			for _, ch := range outChans {
				close(ch)
			}
		}()
		return fn(ctx, inCh, outChans)
	}

	p.lock.Lock()
	p.jobs = append(p.jobs, job)
	p.lock.Unlock()
}

func (p *Pipeline) Run(ctx context.Context) error {
	p.lock.Lock()
	copiedJobs := make([]func(context.Context) error, len(p.jobs))
	copy(copiedJobs, p.jobs)
	p.lock.Unlock()

	group, gCtx := errgroup.WithContext(ctx)

	for _, j := range copiedJobs {
		job := j
		group.Go(func() error {
			return job(gCtx)
		})
	}

	if err := group.Wait(); err != nil {
		return fmt.Errorf("conveyer finished with error: %w", err)
	}
	return nil
}

func (p *Pipeline) Send(name, value string) error {
	ch, ok := p.getChan(name)
	if !ok {
		return ErrChannelMissing
	}

	defer func() { _ = recover() }()

	ch <- value
	return nil
}

func (p *Pipeline) Recv(name string) (string, error) {
	ch, ok := p.getChan(name)
	if !ok {
		return "", ErrChannelMissing
	}

	select {
	case val, ok := <-ch:
		if !ok {
			return Undefined, nil
		}
		return val, nil
	default:
		return Undefined, nil
	}
}
