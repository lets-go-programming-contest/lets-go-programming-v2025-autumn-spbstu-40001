package handlers

import (
	"context"
	"errors"
	"reflect"
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

		case data, ok := <-input:
			if !ok {
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

	cases := make([]reflect.SelectCase, len(inputs)+1)
	cases[0] = reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(ctx.Done()),
	}

	for i, input := range inputs {
		cases[i+1] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(input),
		}
	}

	remaining := len(inputs)

	for remaining > 0 {
		chosen, value, ok := reflect.Select(cases)

		if chosen == 0 {
			return ctx.Err()
		}

		if !ok {
			cases[chosen].Chan = reflect.ValueOf(nil)
			remaining--
			continue
		}

		data := value.String()

		if strings.Contains(data, "no multiplexer") {
			continue
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case output <- data:
		}
	}

	return nil
}

func SeparatorFunc(
	ctx context.Context,
	input chan string,
	outputs []chan string,
) error {
	onceList := make([]*sync.Once, len(outputs))
	for i := range onceList {
		onceList[i] = &sync.Once{}
	}

	defer func() {
		for i, output := range outputs {
			onceList[i].Do(func() {
				close(output)
			})
		}
	}()

	for i := 0; ; {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case data, ok := <-input:
			if !ok {
				return nil
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			case outputs[i%len(outputs)] <- data:
				i++
			}
		}
	}
}
