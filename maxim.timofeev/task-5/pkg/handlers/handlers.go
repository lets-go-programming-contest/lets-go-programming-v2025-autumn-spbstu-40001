package handlers

import (
	"context"
	"errors"
	"strings"
)

// PrefixDecoratorFunc - модификатор данных с префиксом
func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	const prefix = "decorated: "

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

			if !strings.HasPrefix(data, prefix) {
				data = prefix + data
			}

			select {
			case output <- data:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}

// SeparatorFunc - сепаратор по порядковому номеру
func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	if len(outputs) == 0 {
		return nil
	}

	counter := 0

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-input:
			if !ok {
				return nil
			}

			idx := counter % len(outputs)
			counter++

			select {
			case outputs[idx] <- data:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}

// MultiplexerFunc - мультиплексор с фильтрацией
func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	// Простая реализация - обрабатываем входы по очереди
	for {
		dataReceived := false

		for _, input := range inputs {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case data, ok := <-input:
				if !ok {
					continue
				}
				dataReceived = true

				if strings.Contains(data, "no multiplexer") {
					continue
				}

				select {
				case output <- data:
				case <-ctx.Done():
					return ctx.Err()
				}
			default:
				// Пропускаем если нет данных
			}
		}

		// Если ни один канал не дал данных, делаем небольшую паузу
		if !dataReceived {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				// Продолжаем цикл
			}
		}
	}
}
