package handlers

import (
	"context"
	"errors"
	"strings"
	"sync/atomic"
)

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	defer close(output)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(data, "no decorator") {
				return errors.New("can't be decorated")
			}

			const prefix = "decorated: "
			if !strings.HasPrefix(data, prefix) {
				data = prefix + data
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
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

	var counter uint64

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-input:
			if !ok {
				return nil
			}

			idx := atomic.AddUint64(&counter, 1) % uint64(len(outputs))

			select {
			case <-ctx.Done():
				return ctx.Err()
			case outputs[idx] <- data:
			}
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	defer close(output)

	if len(inputs) == 0 {
		return nil
	}

	type result struct {
		data string
		ok   bool
	}

	results := make(chan result, len(inputs))

	for _, input := range inputs {
		go func(in chan string) {
			defer func() {
				select {
				case results <- result{"", false}:
				case <-ctx.Done():
				}
			}()

			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-in:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case results <- result{data, true}:
					}
				}
			}
		}(input)
	}

	closedInputs := 0
	totalInputs := len(inputs)

	for closedInputs < totalInputs {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case res := <-results:
			if !res.ok {
				closedInputs++
				continue
			}

			if strings.Contains(res.data, "no multiplexer") {
				continue
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			case output <- res.data:
			}
		}
	}

	return nil
}
