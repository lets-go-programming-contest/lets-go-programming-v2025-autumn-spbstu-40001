package handlers

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/sync/errgroup"
)

const (
	ErrNoDecorator   = "no decorator"
	DecoratedPrefix  = "decorated: "
	ErrNoMultiplexer = "no multiplexer"
)

var (
	ErrPrefixDecoratorCantBeDecorated = errors.New("handlers.PrefixDecoratorFunc: can't be decorated")
	ErrNoOutputChannels               = errors.New("handlers.SeparatorFunc: no output channels")
	ErrNoInputChannels                = errors.New("handlers.MultiplexerFunc: no input channels")
)

func PrefixDecoratorFunc(
	ctx context.Context,
	input chan string,
	output chan string,
) error {
	defer close(output)

	for {
		select {
		case <-ctx.Done():
			return nil

		case data, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(data, ErrNoDecorator) {
				return ErrPrefixDecoratorCantBeDecorated
			}

			if strings.HasPrefix(data, DecoratedPrefix) {
				select {
				case <-ctx.Done():
					return nil
				case output <- data:
				}

				continue
			}

			select {
			case <-ctx.Done():
				return nil
			case output <- DecoratedPrefix + data:
			}
		}
	}
}

func MultiplexerFunc(
	ctx context.Context,
	inputs []chan string,
	output chan string,
) error {
	if len(inputs) == 0 {
		return ErrNoInputChannels
	}

	defer close(output)

	errGroup, ctx := errgroup.WithContext(ctx)

	for _, inputChan := range inputs {
		errGroup.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return nil
				case data, ok := <-inputChan:
					if !ok {
						return nil
					}

					if strings.Contains(data, ErrNoMultiplexer) {
						continue
					}

					select {
					case <-ctx.Done():
						return nil
					case output <- data:
					}
				}
			}
		})
	}

	if err := errGroup.Wait(); err != nil {
		return fmt.Errorf("multiplexer wait error: %w", err)
	}

	return nil
}

func SeparatorFunc(
	ctx context.Context,
	input chan string,
	outputs []chan string,
) error {
	if len(outputs) == 0 {
		return ErrNoOutputChannels
	}

	defer func() {
		for _, out := range outputs {
			close(out)
		}
	}()

	index := 0

	for {
		select {
		case <-ctx.Done():
			return nil

		case data, ok := <-input:
			if !ok {
				return nil
			}

			select {
			case <-ctx.Done():
				return nil
			case outputs[index] <- data:
				index = (index + 1) % len(outputs)
			}
		}
	}
}
