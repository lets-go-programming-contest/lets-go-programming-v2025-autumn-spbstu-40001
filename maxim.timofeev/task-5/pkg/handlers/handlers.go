package handlers

import (
	"context"
	"errors"
	"strings"
	"sync/atomic"
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
	// Используем отдельную горутину для каждого входного канала
	done := make(chan struct{}, len(inputs))

	// Запускаем горутины для чтения из каждого входного канала
	for i, input := range inputs {
		go func(idx int, in chan string) {
			defer func() {
				done <- struct{}{}
			}()

			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-in:
					if !ok {
						return
					}

					// Фильтруем данные с подстрокой "no multiplexer"
					if strings.Contains(data, "no multiplexer") {
						continue
					}

					select {
					case <-ctx.Done():
						return
					case output <- data:
					}
				}
			}
		}(i, input)
	}

	// Ждем завершения всех горутин
	for i := 0; i < len(inputs); i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-done:
		}
	}

	return nil
}
