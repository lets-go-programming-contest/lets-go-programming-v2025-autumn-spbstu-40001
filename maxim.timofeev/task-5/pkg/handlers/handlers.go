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
			case output <- data:
			case <-ctx.Done():
				return nil
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
			case outputs[idx] <- data:
			case <-ctx.Done():
				return nil
			}
		}
	}
}

// MultiplexerFunc - мультиплексор с фильтрацией
func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	// Создаем канал для слияния данных
	merged := make(chan string, len(inputs)*10)

	// Запускаем горутину для каждого входа
	done := make(chan struct{})
	defer close(done)

	for _, input := range inputs {
		go func(in chan string) {
			defer func() {
				select {
				case done <- struct{}{}:
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

					// Фильтрация
					if strings.Contains(data, "no multiplexer") {
						continue
					}

					select {
					case merged <- data:
					case <-ctx.Done():
						return
					}
				}
			}
		}(input)
	}

	// Закрываем merged после завершения всех горутин
	go func() {
		for i := 0; i < len(inputs); i++ {
			select {
			case <-done:
			case <-ctx.Done():
				return
			}
		}
		close(merged)
	}()

	// Отправляем данные в выходной канал
	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-merged:
			if !ok {
				return nil
			}

			select {
			case output <- data:
			case <-ctx.Done():
				return nil
			}
		}
	}
}
