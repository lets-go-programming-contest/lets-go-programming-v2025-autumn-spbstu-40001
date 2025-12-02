package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var ErrCannotDecorate = errors.New("can't be decorated")

func PrefixDecoratorFunc(ctx context.Context, src chan string, dst chan string) error {
	const prefix = "decorated: "

	for {
		select {
		case <-ctx.Done():
			return nil
		case val, ok := <-src:
			if !ok {
				return nil
			}

			if strings.Contains(val, "no decorator") {
				return ErrCannotDecorate
			}

			if !strings.HasPrefix(val, prefix) {
				val = prefix + val
			}

			select {
			case dst <- val:
			case <-ctx.Done():
				return nil
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	if len(outputs) == 0 {
		return nil
	}

	counter := 0

	for {
		select {
		case <-ctx.Done():
			return nil
		case val, ok := <-input:
			if !ok {
				return nil
			}

			target := counter % len(outputs)
			counter++

			select {
			case outputs[target] <- val:
			case <-ctx.Done():
				return nil
			}
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, out chan string) error {
	var wg sync.WaitGroup

	worker := func(ch chan string) {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-ch:
				if !ok {
					return
				}

				if strings.Contains(val, "no multiplexer") {
					continue
				}

				select {
				case out <- val:
				case <-ctx.Done():
					return
				}
			}
		}
	}

	for _, ch := range inputs {
		wg.Add(1)
		go worker(ch)
	}

	wg.Wait()
	return nil
}
