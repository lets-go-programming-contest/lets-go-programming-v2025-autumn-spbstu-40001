package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

const (
	Decorator     = "decorated: "
	NoDecorator   = "no decorator"
	NoMultiplexer = "no multiplexer"
)

var (
	ErrDecoration   = errors.New("can't be decorated")
	ErrEmptyChannel = errors.New("channel can't be empty")
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
				return ErrDecoration
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

func SeparatorFunc(
	ctx context.Context,
	input chan string,
	outputs []chan string,
) error {
	if len(outputs) == 0 {
		return ErrEmptyChannel
	}

	index := 0

	for {
		select {
		case <-ctx.Done():
			return nil
		case message, ok := <-input:
			if !ok {
				return nil
			}

			target := outputs[index]
			index = (index + 1) % len(outputs)

			select {
			case target <- message:
			case <-ctx.Done():
				return nil
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
		return ErrEmptyChannel
	}

	var waitGroup sync.WaitGroup

	for _, channel := range inputs {
		go func(inp chan string) {
			defer waitGroup.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case message, ok := <-inp:
					if !ok {
						return
					}

					if strings.Contains(message, NoMultiplexer) {
						continue
					}

					select {
					case output <- message:
					case <-ctx.Done():
						return
					}
				}
			}
		}(channel)
	}

	waitGroup.Wait()

	return nil
}
