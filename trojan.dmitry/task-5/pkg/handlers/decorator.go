package handlers

import (
	"context"
	"errors"
	"strings"
)

var ErrDecorator = errors.New("can't be decorated")

const (
	Decorator   = "decorated: "
	NoDecorator = "no decorator"
)

func PrefixDecoratorFunc(
	ctx context.Context,
	input chan string,
	output chan string,
) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case message, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(message, NoDecorator) {
				return ErrDecorator
			}

			if !strings.HasPrefix(message, Decorator) {
				message = Decorator + message
			}

			select {
			case output <- message:
			case <-ctx.Done():
				return nil
			}
		}
	}
}
