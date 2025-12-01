package handlers

import (
	"context"
	"errors"
	"strings"
)

type conveyer interface {
	RegisterDecorator(
		fn func(
			ctx context.Context,
			input chan string,
			output chan string,
		) error,
		input string,
		output string,
	)

	RegisterMultiplexer(
		fn func(
			ctx context.Context,
			inputs []chan string,
			output chan string,
		) error,
		inputs []string,
		output string,
	)

	RegisterSeparator(
		fn func(
			ctx context.Context,
			input chan string,
			outputs []chan string,
		) error,
		input string,
		outputs []string,
	)

	Run(ctx context.Context) error

	Send(data string, input string) error
	Recv(output string) (string, error)
}

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	defer func() {
		close(output)
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case v, ok := <-input:
			if !ok {
				return nil
			}
			if strings.Contains(v, "no decorator") {
				return errors.New("can't be decorated: contains 'no decorator'")
			}
			if strings.HasPrefix(v, "decorated: ") {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case output <- v:
				}
			} else {
				toSend := "decorated: " + v
				select {
				case <-ctx.Done():
					return ctx.Err()
				case output <- toSend:
				}
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	defer func() {
		for _, ch := range outputs {
			close(ch)
		}
	}()

	if len(outputs) == 0 {
		return nil
	}
	idx := 0
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case v, ok := <-input:
			if !ok {
				return nil
			}
			outCh := outputs[idx%len(outputs)]
			idx++
			select {
			case <-ctx.Done():
				return ctx.Err()
			case outCh <- v:
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

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			allClosed := true

			for _, ch := range inputs {
				select {
				case v, ok := <-ch:
					if ok {
						if strings.Contains(v, "no multiplexer") {
							continue
						}
						output <- v
						allClosed = false
					}
				default:
				}
			}

			if allClosed {
				return nil
			}
		}
	}
}
