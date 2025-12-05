package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var errPrefixDecoratorFuncCantBeDecorated = errors.New("handlers.PrefixDecoratorFunc: can't be decorated")

func PrefixDecoratorFunc(
	ctx context.Context,
	input chan string,
	output chan string,
) error {
	defer close(output)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case data, isOpen := <-input:
			if !isOpen {
				return nil
			}

			if strings.Contains(data, "no decorator") {
				return errPrefixDecoratorFuncCantBeDecorated
			}

			if strings.HasPrefix(data, "decorated: ") {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case output <- data:
				}

				continue
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			case output <- "decorated: " + data:
			}
		}
	}
}

func MultiplexerFunc(
	ctx context.Context,
	inputs []chan string,
	output chan string,
) error {
	defer close(output)

	var wg sync.WaitGroup
	errChan := make(chan error, 1)

	for _, inputChan := range inputs {
		wg.Add(1)
		go func(in chan string) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case data, isOpen := <-in:
					if !isOpen {
						return
					}

					if strings.Contains(data, "no multiplexer") {
						continue
					}

					select {
					case <-ctx.Done():
						return
					case output <- data:
					}
				}
			}
		}(inputChan)
	}

	go func() {
		wg.Wait()
		errChan <- nil
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}

func SeparatorFunc(
	ctx context.Context,
	input chan string,
	outputs []chan string,
) error {
	onceList := make([]*sync.Once, len(outputs))
	for idx := range onceList {
		onceList[idx] = &sync.Once{}
	}

	defer func() {
		for idx, output := range outputs {
			onceList[idx].Do(func() {
				close(output)
			})
		}
	}()

	counter := 0

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case data, isOpen := <-input:
			if !isOpen {
				return nil
			}

			sent := false
			for attempt := 0; attempt < len(outputs) && !sent; attempt++ {
				idx := (counter + attempt) % len(outputs)
				select {
				case <-ctx.Done():
					return ctx.Err()
				case outputs[idx] <- data:
					counter = idx + 1
					sent = true
				default:
				}
			}

			if !sent {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case outputs[counter%len(outputs)] <- data:
					counter++
				}
			}
		}
	}
}
