package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

const Undefined = "undefined"

var ErrChannelNotFound = errors.New("chan not found")

type Conveyer struct {
	lock     sync.RWMutex
	pipes    map[string]chan string
	workers  []func(context.Context) error
	capacity int
}

func New(size int) *Conveyer {
	return &Conveyer{
		lock:     sync.RWMutex{},
		pipes:    make(map[string]chan string),
		workers:  []func(context.Context) error{},
		capacity: size,
	}
}

func (c *Conveyer) ensure(name string) chan string {
	c.lock.RLock()
	exist := c.pipes[name]
	c.lock.RUnlock()

	if exist != nil {
		return exist
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	exist = c.pipes[name]
	if exist != nil {
		return exist
	}

	created := make(chan string, c.capacity)
	c.pipes[name] = created

	return created
}

func (c *Conveyer) lookup(name string) (chan string, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	ch, ok := c.pipes[name]

	return ch, ok
}

func (c *Conveyer) RegisterDecorator(
	functionn func(ctx context.Context, input chan string, output chan string) error,
	input, output string,
) {
	inp := c.ensure(input)
	out := c.ensure(output)

	c.lock.Lock()
	c.workers = append(c.workers, func(ctx context.Context) error {
		defer close(out)

		return functionn(ctx, inp, out)
	})
	c.lock.Unlock()
}

func (c *Conveyer) RegisterMultiplexer(
	functionn func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	inList := make([]chan string, 0, len(inputs))
	for _, n := range inputs {
		inList = append(inList, c.ensure(n))
	}

	out := c.ensure(output)

	job := func(ctx context.Context) error {
		defer close(out)

		return functionn(ctx, inList, out)
	}

	c.workers = append(c.workers, job)
}

func (c *Conveyer) RegisterSeparator(
	functionn func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	inp := c.ensure(input)

	outs := make([]chan string, 0, len(outputs))
	for _, n := range outputs {
		outs = append(outs, c.ensure(n))
	}

	job := func(ctx context.Context) error {
		defer func() {
			for _, ch := range outs {
				close(ch)
			}
		}()

		return functionn(ctx, inp, outs)
	}

	c.lock.Lock()
	c.workers = append(c.workers, job)
	c.lock.Unlock()
}

func (c *Conveyer) Run(ctx context.Context) error {
	c.lock.RLock()
	workers := c.workers
	c.lock.RUnlock()

	group, gctx := errgroup.WithContext(ctx)

	for i := range workers {
		job := workers[i]

		group.Go(func() error {
			return job(gctx)
		})
	}

	if err := group.Wait(); err != nil {
		return fmt.Errorf("conveyer run failed: %w", err)
	}

	return nil
}

func (c *Conveyer) Send(input, data string) error {
	channel, exists := c.lookup(input)
	if !exists {
		return ErrChannelNotFound
	}

	defer func() { _ = recover() }()
	channel <- data

	return nil
}

func (c *Conveyer) Recv(output string) (string, error) {
	channel, exists := c.lookup(output)
	if !exists {
		return "", ErrChannelNotFound
	}

	val, ok := <-channel
	if !ok {
		return Undefined, nil
	}

	return val, nil
}
