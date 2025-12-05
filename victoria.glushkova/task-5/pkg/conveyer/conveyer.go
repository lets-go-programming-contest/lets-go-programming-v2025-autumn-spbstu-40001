package conveyer

import (
	"context"
	"errors"
	"sync"
)

type Pipeline struct {
	size      int
	channels  map[string]chan string
	workers   []worker
	mu        sync.RWMutex
}

type worker struct {
	process func(ctx context.Context) error
}

type PipelineInterface interface {
	AddDecorator(
		fn func(ctx context.Context, in <-chan string, out chan<- string) error,
		inputChan, outputChan string,
	)
	AddMultiplexer(
		fn func(ctx context.Context, ins []<-chan string, out chan<- string) error,
		inputs []string, output string,
	)
	AddSeparator(
		fn func(ctx context.Context, in <-chan string, outs []chan<- string) error,
		input string, outputs []string,
	)
	Execute(ctx context.Context) error
	SendTo(channelID, data string) error
	ReceiveFrom(channelID string) (string, error)
}

func New(size int) *Pipeline {
	return &Pipeline{
		size:     size,
		channels: make(map[string]chan string),
		workers:  make([]worker, 0),
	}
}

func (p *Pipeline) AddDecorator(
	fn func(ctx context.Context, in <-chan string, out chan<- string) error,
	inputID, outputID string,
) {
	p.mu.Lock()
	defer p.mu.Unlock()

	inChan := p.obtainChannel(inputID)
	outChan := p.obtainChannel(outputID)

	p.workers = append(p.workers, worker{
		process: func(ctx context.Context) error {
			return fn(ctx, inChan, outChan)
		},
	})
}

func (p *Pipeline) AddMultiplexer(
	fn func(ctx context.Context, ins []<-chan string, out chan<- string) error,
	inputIDs []string, outputID string,
) {
	p.mu.Lock()
	defer p.mu.Unlock()

	inChans := make([]<-chan string, len(inputIDs))
	for i, id := range inputIDs {
		inChans[i] = p.obtainChannel(id)
	}
	outChan := p.obtainChannel(outputID)

	p.workers = append(p.workers, worker{
		process: func(ctx context.Context) error {
			return fn(ctx, inChans, outChan)
		},
	})
}

func (p *Pipeline) AddSeparator(
	fn func(ctx context.Context, in <-chan string, outs []chan<- string) error,
	inputID string, outputIDs []string,
) {
	p.mu.Lock()
	defer p.mu.Unlock()

	inChan := p.obtainChannel(inputID)
	outChans := make([]chan<- string, len(outputIDs))
	for i, id := range outputIDs {
		ch := p.obtainChannel(id)
		outChans[i] = ch
	}

	p.workers = append(p.workers, worker{
		process: func(ctx context.Context) error {
			return fn(ctx, inChan, outChans)
		},
	})
}

func (p *Pipeline) Execute(ctx context.Context) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(p.workers))

	for _, w := range p.workers {
		wg.Add(1)
		go func(w worker) {
			defer wg.Done()
			if err := w.process(ctx); err != nil {
				select {
				case errCh <- err:
				default:
				}
			}
		}(w)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	select {
	case <-ctx.Done():
		p.closeAllChannels()
		return ctx.Err()
	case err := <-errCh:
		p.closeAllChannels()
		return err
	}
}

func (p *Pipeline) SendTo(channelID, data string) error {
	p.mu.RLock()
	ch, exists := p.channels[channelID]
	p.mu.RUnlock()

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

func (p *Pipeline) ReceiveFrom(channelID string) (string, error) {
	p.mu.RLock()
	ch, exists := p.channels[channelID]
	p.mu.RUnlock()

	if !exists {
		return "", errors.New("chan not found")
	}

	value, ok := <-ch
	if !ok {
		return "undefined", nil
	}

	return value, nil
}

func (p *Pipeline) obtainChannel(name string) chan string {
	if ch, found := p.channels[name]; found {
		return ch
	}
	ch := make(chan string, p.size)
	p.channels[name] = ch
	return ch
}

func (p *Pipeline) closeAllChannels() {
	p.mu.Lock()
	defer p.mu.Unlock()

	for name, ch := range p.channels {
		close(ch)
		delete(p.channels, name)
	}
}
