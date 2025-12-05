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
			return errors.Join(ctx.Err())

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
					return errors.Join(ctx.Err())
				case output <- data:
				}

				continue
			}

			select {
			case <-ctx.Done():
				return errors.Join(ctx.Err())
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
		Send: reflect.Value{},
	}

	for idx, input := range inputs {
		cases[idx+1] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(input),
			Send: reflect.Value{},
		}
	}

	remaining := len(inputs)

	for remaining > 0 {
		chosen, value, isOpen := reflect.Select(cases)

		if chosen == 0 {
			return errors.Join(ctx.Err())
		}

		if !isOpen {
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
			return errors.Join(ctx.Err())
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
			return errors.Join(ctx.Err())

		case data, isOpen := <-input:
			if !isOpen {
				return nil
			}

			select {
			case <-ctx.Done():
				return errors.Join(ctx.Err())
			case outputs[counter%len(outputs)] <- data:
				counter++
			}
		}
	}
}
