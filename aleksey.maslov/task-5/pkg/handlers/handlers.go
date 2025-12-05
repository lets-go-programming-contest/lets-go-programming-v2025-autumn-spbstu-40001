package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var (
	ErrDecorator    = errors.New("can't be decorated")
	ErrEmptyChannel = errors.New("channel can't be empty")
)

const (
	Decorator     = "decorated: "
	NoDecorator   = "no decorator"
	NoMultiplexer = "no multiplexer"
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

func MultiplexerFunc(
	ctx context.Context,
	inputs []chan string,
	output chan string,
) error {
	if len(inputs) == 0 {
		return ErrEmptyChannel
	}

	var waitGrp sync.WaitGroup

	waitGrp.Add(len(inputs))

	for _, channel := range inputs {
		go func(inputChannel chan string) {
			defer waitGrp.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case message, ok := <-inputChannel:
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

	waitGrp.Wait()

	return nil
}
