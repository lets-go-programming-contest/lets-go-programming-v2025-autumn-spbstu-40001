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
	var count int

	numOutputs := len(outputs)
	if numOutputs == 0 {
		return ErrNoOutputs
	}

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

	var wgr sync.WaitGroup

	wgr.Add(len(inputs))

	multiplex := func(chn chan string) {
		defer wgr.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case data, ok := <-chn:
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

	for _, chn := range inputs {
		ch := chn

		go multiplex(ch)
	}

	wgr.Wait()

	return nil
}
