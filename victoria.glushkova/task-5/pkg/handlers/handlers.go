package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var (
	ErrNoDecorator = errors.New("can't be decorated")
)

const (
	noDecoratorData   = "no decorator"
	decoratorPrefix   = "decorated: "
	noMultiplexerData = "no multiplexer"
)

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	defer close(output)

	for {
		select {
		case <-ctx.Done():
			return nil

		case data, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(data, noDecoratorData) {
				return ErrNoDecorator
			}

			if !strings.HasPrefix(data, decoratorPrefix) {
				data = decoratorPrefix + data
			}

			select {
			case <-ctx.Done():
				return nil
			case output <- data:
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	defer func() {
		for _, out := range outputs {
			close(out)
		}
	}()

	outputCount := len(outputs)
	currentIndex := 0

	for {
		select {
		case <-ctx.Done():
			return nil

		case data, ok := <-input:
			if !ok {
				return nil
			}

			if outputCount == 0 {
				continue
			}

			select {
			case <-ctx.Done():
				return nil
			case outputs[currentIndex%outputCount] <- data:
				currentIndex++
			}
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	defer close(output)

	if len(inputs) == 0 {
		return nil
	}

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(inputs))

	for _, inputChannel := range inputs {
		localInputChannel := inputChannel

		go func(currentChannel chan string) {
			defer waitGroup.Done()

			for {
				select {
				case <-ctx.Done():
					return

				case data, ok := <-currentChannel:
					if !ok {
						return
					}

					if strings.Contains(data, noMultiplexerData) {
						continue
					}

					select {
					case <-ctx.Done():
						return
					case output <- data:
					}
				}
			}
		}(localInputChannel)
	}

	waitGroup.Wait()

	return nil
}
