package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

// Errors.
var (
	ErrCannotDecorate      = errors.New("can't be decorated")
	ErrSeparatorCanceled   = errors.New("separator canceled")
	ErrMultiplexerCanceled = errors.New("multiplexer canceled")
)

// Funcs with business logic.
// Decorator.
func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	for {
		select {
		case <-ctx.Done():
			return errors.Join(ctx.Err(), ErrCannotDecorate)
		case val, ok := <-input:
			if !ok {
				close(output)

				return nil
			}

			if strings.Contains(val, "no decorator") {
				close(output)

				return ErrCannotDecorate
			}

			if !strings.HasPrefix(val, "decorated: ") {
				val = "decorated: " + val
			}

			output <- val
		}
	}
}

// Funcs with business logic.
// Separator.
func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	currentIndex := 0

	for {
		select {
		case <-ctx.Done():
			return errors.Join(ctx.Err(), ErrSeparatorCanceled)
		case val, ok := <-input:
			if !ok {
				for _, outChan := range outputs {
					close(outChan)
				}

				return nil
			}

			outputs[currentIndex%len(outputs)] <- val

			currentIndex++
		}
	}
}

// Funcs with business logic.
// Multiplexer.
func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(inputs))

	for _, inputChannel := range inputs {
		go func(inputChan chan string) {
			defer waitGroup.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-inputChan:
					if !ok {
						return
					}

					if strings.Contains(val, "no multiplexer") {
						continue
					}

					output <- val
				}
			}
		}(inputChannel)
	}

	done := make(chan struct{})

	go func() {
		waitGroup.Wait()
		close(done)
	}()

	select {
	case <-ctx.Done():
		return errors.Join(ctx.Err(), ErrMultiplexerCanceled)
	case <-done:
		close(output)

		return nil
	}
}
