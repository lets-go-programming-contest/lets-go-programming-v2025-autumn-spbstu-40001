package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var (
	ErrNoDecorator      = errors.New("can't be decorated")
	ErrEmptyChannelList = errors.New("channels slice is empty")
)

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		select {
		case str, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(str, "no decorator") {
				return ErrNoDecorator
			}

			if !strings.HasPrefix(str, "decorated: ") {
				str = "decorated: " + str
			}

			select {
			case output <- str:
			case <-ctx.Done():
				return nil
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	if len(outputs) == 0 {
		return ErrEmptyChannelList
	}

	var idx int

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		select {
		case str, ok := <-input:
			if !ok {
				return nil
			}

			select {
			case outputs[idx] <- str:
			case <-ctx.Done():
				return nil
			}

			idx = (idx + 1) % len(outputs)
		case <-ctx.Done():
			return nil
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	if len(inputs) == 0 {
		return ErrEmptyChannelList
	}

	var group sync.WaitGroup

	group.Add(len(inputs))

	for idx := range inputs {
		go func() {
			defer group.Done()

			for {
				select {
				case <-ctx.Done():
					return
				default:
				}

				select {
				case str, ok := <-inputs[idx]:
					if !ok {
						return
					}

					if !strings.Contains(str, "no multiplexer") {
						select {
						case output <- str:
						case <-ctx.Done():
							return
						}
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	group.Wait()

	return nil
}
