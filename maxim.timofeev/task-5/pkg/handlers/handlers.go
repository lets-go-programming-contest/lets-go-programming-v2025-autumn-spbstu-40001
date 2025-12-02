package handlers

import (
	"context"
	"errors"
	"strings"
	"sync/atomic"
)

// PrefixDecoratorFunc - модификатор данных с префиксом
func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	prefix := "decorated: "

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-input:
			if !ok {
				return nil
			}

			// Проверяем, содержит ли данные подстроку "no decorator"
			if strings.Contains(data, "no decorator") {
				return errors.New("can't be decorated")
			}

			// Проверяем, не добавлен ли уже префикс
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

// SeparatorFunc - сепаратор по порядковому номеру
func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	if len(outputs) == 0 {
		return nil
	}

	var counter int64 = -1

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-input:
			if !ok {
				return nil
			}

			// Увеличиваем счетчик атомарно
			idx := int(atomic.AddInt64(&counter, 1)) % len(outputs)

			select {
			case <-ctx.Done():
				return ctx.Err()
			case outputs[idx] <- data:
			}
		}
	}
}

// MultiplexerFunc - мультиплексор с фильтрацией
func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	// Создаем канал для слияния всех входов
	merged := make(chan string, len(inputs)*10)

	// Запускаем горутины для каждого входного канала
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
					select {
					case <-ctx.Done():
						return
					case merged <- data:
					}
				}
			}
		}(input)
	}

	// Ждем завершения всех горутин
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

	// Обрабатываем объединенные данные
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-merged:
			if !ok {
				return nil
			}

			// Фильтруем данные с подстрокой "no multiplexer"
			if strings.Contains(data, "no multiplexer") {
				continue
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			case output <- data:
			}
		}
	}
}
