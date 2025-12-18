package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var (
	ErrCannotBeDecorated = errors.New("can't be decorated")
	ErrCannotMultiplex   = errors.New("can't multiplex")
	ErrNoInputChannels   = errors.New("no input channels")
	ErrNoOutputChannels  = errors.New("no output channels")
)

const (
	NoDecoratorMarker   = "no decorator"
	NoMultiplexerMarker = "no multiplexer"
	DecoratedPrefix     = "decorated: "
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

			if strings.Contains(val, NoDecoratorMarker) {
				return ErrCannotBeDecorated
			}

			if !strings.HasPrefix(val, DecoratedPrefix) {
				val = DecoratedPrefix + val
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
	if len(inputs) == 0 {
		return ErrNoInputChannels
	}

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

				if strings.Contains(val, NoMultiplexerMarker) {
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

		localInputChan := inputChan
		localProcessInput := processInput

		go func() {
			localProcessInput(localInputChan)
		}()
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
	if len(outputs) == 0 {
		return ErrNoOutputChannels
	}

	counter := -1

	outputsCount := len(outputs)

	for {
		select {
		case <-ctx.Done():
			return nil
		case val, ok := <-input:
			if !ok {
				return nil
			}

			counter = (counter + 1) % outputsCount
			out := outputs[counter]

			select {
			case out <- val:
			case <-ctx.Done():
				return nil
			}
		}
	}
}
