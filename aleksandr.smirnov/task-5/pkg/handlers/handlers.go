package handlers

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/sync/errgroup"
)

var (
	ErrCannotDecorate = errors.New("can't be decorated")
	ErrNoOutputs      = errors.New("outputs must not be empty")
)

const (
	noDecoratorText   = "no decorator"
	noMultiplexerText = "no multiplexer"
	decoratorPrefix   = "decorated: "
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

			if strings.Contains(data, noDecoratorText) {
				return ErrCannotDecorate
			}

			if !strings.HasPrefix(data, decoratorPrefix) {
				data = decoratorPrefix + data
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

	counter := 0

	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-input:
			if !ok {
				return nil
			}

			select {
			case outputs[counter] <- data:
			case <-ctx.Done():
				return nil
			}

			counter = (counter + 1) % len(outputs)
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	if len(inputs) == 0 {
		return nil
	}

	group, groupCtx := errgroup.WithContext(ctx)

	for _, inputChan := range inputs {
		group.Go(func() error {
			for {
				select {
				case <-groupCtx.Done():
					return nil
				case data, ok := <-inputChan:
					if !ok {
						return nil
					}

					if strings.Contains(data, noMultiplexerText) {
						continue
					}

					select {
					case output <- data:
					case <-groupCtx.Done():
						return nil
					}
				}
			}
		})
	}

	if err := group.Wait(); err != nil {
		return fmt.Errorf("multiplexer error: %w", err)
	}

	return nil
}
