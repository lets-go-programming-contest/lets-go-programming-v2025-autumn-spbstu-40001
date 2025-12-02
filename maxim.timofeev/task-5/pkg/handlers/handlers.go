package handlers

import (
	"context"
	"errors"
	"strings"
)

// PrefixDecoratorFunc - модификатор данных с префиксом
func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(data, "no decorator") {
				return errors.New("can't be decorated")
			}

			if !strings.HasPrefix(data, "decorated: ") {
				data = "decorated: " + data
			}

			select {
			case <-ctx.Done():
				return nil
			case output <- data:
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
			return nil
		case data, ok := <-input:
			if !ok {
				return nil
			}

			idx := counter % len(outputs)
			counter++

			select {
			case <-ctx.Done():
				return nil
			case outputs[idx] <- data:
			}
		}
	}
}

// MultiplexerFunc - мультиплексор с фильтрацией
func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	// Используем простой подход - читаем из всех каналов по очереди
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			// Проверяем каждый входной канал
			for _, input := range inputs {
				select {
				case <-ctx.Done():
					return nil
				case data, ok := <-input:
					if !ok {
						continue
					}

					// Фильтрация
					if strings.Contains(data, "no multiplexer") {
						continue
					}

					select {
					case <-ctx.Done():
						return nil
					case output <- data:
					}
				default:
					// Нет данных в этом канале, продолжаем
				}
			}
		}
	}
}
