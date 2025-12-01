package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
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
				// вход закрыт — завершаем работу
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

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	if len(inputs) == 0 {
		return nil
	}
	var wg sync.WaitGroup
	wg.Add(len(inputs))

	agg := make(chan string)

	for _, in := range inputs {
		inCh := in
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-inCh:
					if !ok {
						return
					}
					if strings.Contains(v, "no multiplexer") {
						continue
					}
					select {
					case <-ctx.Done():
						return
					case agg <- v:
					}
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(agg)
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case v, ok := <-agg:
			if !ok {
				return nil
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case output <- v:
			}
		}
	}
}
