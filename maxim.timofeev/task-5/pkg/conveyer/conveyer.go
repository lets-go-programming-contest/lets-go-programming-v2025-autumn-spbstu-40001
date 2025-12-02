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

type Conveyer struct {
	mu              sync.Mutex
	chans           map[string]chan string
	jobs            []func(context.Context) error
	channelCapacity int
}

func New(capacity int) *Conveyer {
	return &Conveyer{
		mu:              sync.Mutex{},
		chans:           make(map[string]chan string),
		jobs:            []func(context.Context) error{},
		channelCapacity: capacity,
	}
}

func (p *Conveyer) getOrCreateChan(name string) chan string {
	p.mu.Lock()
	defer p.mu.Unlock()

	if ch, exists := p.chans[name]; exists {
		return ch
	}

	ch := make(chan string, p.channelCapacity)
	p.chans[name] = ch

	return ch
}

func (p *Conveyer) getChan(name string) (chan string, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	ch, exists := p.chans[name]

	return ch, exists
}

func (p *Conveyer) RegisterDecorator(
	funct func(ctx context.Context, input chan string, output chan string) error,
	inputName, outputName string,
) {
	inCh := p.getOrCreateChan(inputName)
	outCh := p.getOrCreateChan(outputName)

	job := func(ctx context.Context) error {
		defer close(outCh)

		return funct(ctx, inCh, outCh)
	}

	p.mu.Lock()
	p.jobs = append(p.jobs, job)
	p.mu.Unlock()
}

func (p *Conveyer) RegisterMultiplexer(
	funct func(ctx context.Context, inputs []chan string, output chan string) error,
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

		return funct(ctx, inChans, outCh)
	}

	p.mu.Lock()
	p.jobs = append(p.jobs, job)
	p.mu.Unlock()
}

func (p *Conveyer) RegisterSeparator(
	funct func(ctx context.Context, input chan string, outputs []chan string) error,
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

		return funct(ctx, inCh, outChans)
	}

	p.mu.Lock()
	p.jobs = append(p.jobs, job)
	p.mu.Unlock()
}

func (p *Conveyer) Run(ctx context.Context) error {
	p.mu.Lock()
	copiedJobs := make([]func(context.Context) error, len(p.jobs))
	copy(copiedJobs, p.jobs)
	p.mu.Unlock()

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

func (p *Conveyer) Send(name, value string) error {
	channel, ok := p.getChan(name)
	if !ok {
		return ErrChannelMissing
	}

	defer func() { _ = recover() }()

	channel <- value

	return nil
}

func (p *Conveyer) Recv(name string) (string, error) {
	channel, ok := p.getChan(name)
	if !ok {
		return "", ErrChannelMissing
	}

	select {
	case val, ok := <-channel:
		if !ok {
			return Undefined, nil
		}
		return val, nil
	default:
		return Undefined, nil
	}
}
