package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var (
	ErrNoDecorator = errors.New("can't be decorated")
	ErrNoOutputs   = errors.New("no output channels provided for separator")
	ErrNoInputs    = errors.New("no input channels provided for multiplexer")
)

const (
	DecoratorPrefix     = "decorated: "
	ErrNoDecoratorMsg   = "no decorator"
	ErrNoMultiplexerMsg = "no multiplexer"
)

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(data, ErrNoDecoratorMsg) {
				return ErrNoDecorator
			}

			if !strings.HasPrefix(data, DecoratorPrefix) {
				data = DecoratorPrefix + data
			}

			select {
			case output <- data:
			case <-ctx.Done():
				return nil
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	if len(outputs) == 0 {
		return ErrNoOutputs
	}

	var count int
	var numOutputs = len(outputs)

	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-input:
			if !ok {
				return nil
			}

			targetIndex := count % numOutputs
			count++

			select {
			case outputs[targetIndex] <- data:
			case <-ctx.Done():
				return nil
			}
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	if len(inputs) == 0 {
		return ErrNoInputs
	}

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(inputs))

	multiplex := func(channel chan string) {
		defer waitGroup.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case data, ok := <-channel:
				if !ok {
					return
				}

				if strings.Contains(data, ErrNoMultiplexerMsg) {
					continue
				}

				select {
				case output <- data:
				case <-ctx.Done():
					return
				}
			}
		}
	}

	for _, channel := range inputs {
		go multiplex(channel)
	}

	waitGroup.Wait()

	return nil
}
