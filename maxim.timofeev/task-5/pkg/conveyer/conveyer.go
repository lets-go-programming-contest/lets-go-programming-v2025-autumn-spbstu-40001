package conveyer

import (
	"context"
	"errors"
	"sync"
)

// Типы обработчиков
type DecoratorFunc func(ctx context.Context, input chan string, output chan string) error
type MultiplexerFunc func(ctx context.Context, inputs []chan string, output chan string) error
type SeparatorFunc func(ctx context.Context, input chan string, outputs []chan string) error

// Интерфейс конвейера
type Conveyer interface {
	RegisterDecorator(fn DecoratorFunc, input string, output string)
	RegisterMultiplexer(fn MultiplexerFunc, inputs []string, output string)
	RegisterSeparator(fn SeparatorFunc, input string, outputs []string)
	Run(ctx context.Context) error
	Send(input string, data string) error
	Recv(output string) (string, error)
}

// Реализация конвейера
type conveyerImpl struct {
	mu       sync.RWMutex
	size     int
	channels map[string]chan string
	tasks    []func(context.Context) error
	started  bool
	wg       sync.WaitGroup
}

// New создает новый конвейер
func New(size int) Conveyer {
	return &conveyerImpl{
		size:     size,
		channels: make(map[string]chan string),
		tasks:    make([]func(context.Context) error, 0),
	}
}

// getOrCreateChannel получает или создает канал с указанным именем
func (c *conveyerImpl) getOrCreateChannel(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ch, ok := c.channels[name]; ok {
		return ch
	}

	ch := make(chan string, c.size)
	c.channels[name] = ch
	return ch
}

// getChannel получает канал по имени
func (c *conveyerImpl) getChannel(name string) (chan string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	ch, ok := c.channels[name]
	if !ok {
		return nil, errors.New("chan not found")
	}
	return ch, nil
}

// RegisterDecorator регистрирует модификатор данных
func (c *conveyerImpl) RegisterDecorator(fn DecoratorFunc, input string, output string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	inputCh := c.getOrCreateChannel(input)
	outputCh := c.getOrCreateChannel(output)

	task := func(ctx context.Context) error {
		defer close(outputCh)
		return fn(ctx, inputCh, outputCh)
	}

	c.tasks = append(c.tasks, task)
}

// RegisterMultiplexer регистрирует мультиплексор
func (c *conveyerImpl) RegisterMultiplexer(fn MultiplexerFunc, inputs []string, output string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	inputChannels := make([]chan string, len(inputs))
	for i, name := range inputs {
		inputChannels[i] = c.getOrCreateChannel(name)
	}
	outputCh := c.getOrCreateChannel(output)

	task := func(ctx context.Context) error {
		defer close(outputCh)
		return fn(ctx, inputChannels, outputCh)
	}

	c.tasks = append(c.tasks, task)
}

// RegisterSeparator регистрирует сепаратор
func (c *conveyerImpl) RegisterSeparator(fn SeparatorFunc, input string, outputs []string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	inputCh := c.getOrCreateChannel(input)
	outputChannels := make([]chan string, len(outputs))
	for i, name := range outputs {
		outputChannels[i] = c.getOrCreateChannel(name)
	}

	task := func(ctx context.Context) error {
		defer func() {
			for _, ch := range outputChannels {
				close(ch)
			}
		}()
		return fn(ctx, inputCh, outputChannels)
	}

	c.tasks = append(c.tasks, task)
}

// Run запускает конвейер
func (c *conveyerImpl) Run(ctx context.Context) error {
	c.mu.Lock()
	if c.started {
		c.mu.Unlock()
		return errors.New("conveyer already started")
	}
	c.started = true
	tasks := make([]func(context.Context) error, len(c.tasks))
	copy(tasks, c.tasks)
	c.mu.Unlock()

	// Создаем канал для ошибок
	errChan := make(chan error, len(tasks))

	// Запускаем все задачи
	for _, task := range tasks {
		c.wg.Add(1)
		go func(t func(context.Context) error) {
			defer c.wg.Done()

			if err := t(ctx); err != nil {
				select {
				case errChan <- err:
				default:
				}
			}
		}(task)
	}

	// Ждем завершения всех задач или ошибки
	go func() {
		c.wg.Wait()
		close(errChan)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}

// Send отправляет данные в канал
func (c *conveyerImpl) Send(input string, data string) error {
	ch, err := c.getChannel(input)
	if err != nil {
		return err
	}

	// Потокобезопасная отправка с recover
	select {
	case ch <- data:
		return nil
	default:
		// Буфер заполнен
		return errors.New("channel buffer full")
	}
}

// Recv получает данные из канала
func (c *conveyerImpl) Recv(output string) (string, error) {
	ch, err := c.getChannel(output)
	if err != nil {
		return "", err
	}

	// Потокобезопасное чтение
	select {
	case value, ok := <-ch:
		if !ok {
			return "undefined", nil
		}
		return value, nil
	default:
		// Нет данных в канале
		return "", errors.New("no data available")
	}
}
