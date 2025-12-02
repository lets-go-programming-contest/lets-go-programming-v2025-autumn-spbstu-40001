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
	defer close(output)

	for {
		select {
		case <-ctx.Done():
			return nil

		case v, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(v, "no decorator") {
				return errors.New("can't be decorated: contains 'no decorator'")
			}

			if !strings.HasPrefix(v, "decorated: ") {
				v = "decorated: " + v
			}

			select {
			case <-ctx.Done():
				return nil
			case output <- v:
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

	i := 0

	for {
		select {
		case <-ctx.Done():
			return nil

		case v, ok := <-input:
			if !ok {
				return nil
			}

			out := outputs[i%len(outputs)]
			i++

			select {
			case <-ctx.Done():
				return nil
			case out <- v:
			}
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	defer close(output)

	merged := make(chan string)
	var wg sync.WaitGroup

	for _, in := range inputs {
		wg.Add(1)
		go func(ch chan string) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-ch:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case merged <- v:
					}
				}
			}
		}(in)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		case v, ok := <-merged:
			if !ok {
				return nil
			}

			if strings.Contains(v, "no multiplexer") {
				continue
			}

			select {
			case <-ctx.Done():
				return nil
			case output <- v:
			}
		}
	}
}
