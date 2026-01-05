package handlers

import (
	"context"
	"errors"
	"strings"
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
