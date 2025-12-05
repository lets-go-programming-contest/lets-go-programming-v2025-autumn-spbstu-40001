package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
	"sync/atomic"
)

var (
	ErrCannotBeDecorated = errors.New("can't be decorated")
	ErrCannotMultiplex   = errors.New("can't multiplex")
)

func PrefixDecoratorFunc(ctx context.Context, input, output chan string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case val, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(val, "no decorator") {
				return ErrCannotBeDecorated
			}

			if !strings.HasPrefix(val, "decorated: ") {
				val = "decorated: " + val
			}

			select {
			case output <- val:
			case <-ctx.Done():
				return nil
			}
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	var waitGroup sync.WaitGroup

	processInput := func(inputChan chan string) {
		defer waitGroup.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-inputChan:
				if !ok {
					return
				}

				if strings.Contains(val, "no multiplexer") {
					continue
				}

				select {
				case output <- val:
				case <-ctx.Done():
					return
				}
			}
		}
	}

	for _, inputChan := range inputs {
		waitGroup.Add(1)
		processInput := processInput // capture variable
		go func(ch chan string) {
			processInput(ch)
		}(inputChan)
	}

	done := make(chan struct{})
	go func() {
		waitGroup.Wait()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return nil
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	var counter int64 = -1
	outputsCount := len(outputs)

	for {
		select {
		case <-ctx.Done():
			return nil
		case val, ok := <-input:
			if !ok {
				return nil
			}

			idx := atomic.AddInt64(&counter, 1) % int64(outputsCount)
			out := outputs[idx]

			select {
			case out <- val:
			case <-ctx.Done():
				return nil
			}
		}
	}
}
