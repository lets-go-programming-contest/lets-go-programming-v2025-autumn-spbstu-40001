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
	mu           sync.RWMutex
	size         int
	channels     map[string]chan string
	decorators   []decoratorInfo
	multiplexers []multiplexerInfo
	separators   []separatorInfo
	ctx          context.Context
	cancel       context.CancelFunc
	wg           sync.WaitGroup
	errChan      chan error
	running      bool
}

type decoratorInfo struct {
	fn     DecoratorFunc
	input  string
	output string
}

type multiplexerInfo struct {
	fn     MultiplexerFunc
	inputs []string
	output string
}

type separatorInfo struct {
	fn      SeparatorFunc
	input   string
	outputs []string
}

// New создает новый конвейер
func New(size int) Conveyer {
	return &conveyerImpl{
		size:     size,
		channels: make(map[string]chan string),
		errChan:  make(chan error, 10),
	}
}

// RegisterDecorator регистрирует модификатор данных
func (c *conveyerImpl) RegisterDecorator(fn DecoratorFunc, input string, output string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.decorators = append(c.decorators, decoratorInfo{
		fn:     fn,
		input:  input,
		output: output,
	})
}

// RegisterMultiplexer регистрирует мультиплексор
func (c *conveyerImpl) RegisterMultiplexer(fn MultiplexerFunc, inputs []string, output string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.multiplexers = append(c.multiplexers, multiplexerInfo{
		fn:     fn,
		inputs: inputs,
		output: output,
	})
}

// RegisterSeparator регистрирует сепаратор
func (c *conveyerImpl) RegisterSeparator(fn SeparatorFunc, input string, outputs []string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.separators = append(c.separators, separatorInfo{
		fn:      fn,
		input:   input,
		outputs: outputs,
	})
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

// createAllChannels создает все каналы, которые еще не созданы
func (c *conveyerImpl) createAllChannels() {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Создаем каналы для декораторов
	for _, d := range c.decorators {
		if _, ok := c.channels[d.input]; !ok {
			c.channels[d.input] = make(chan string, c.size)
		}
		if _, ok := c.channels[d.output]; !ok {
			c.channels[d.output] = make(chan string, c.size)
		}
	}

	// Создаем каналы для сепараторов
	for _, s := range c.separators {
		if _, ok := c.channels[s.input]; !ok {
			c.channels[s.input] = make(chan string, c.size)
		}
		for _, output := range s.outputs {
			if _, ok := c.channels[output]; !ok {
				c.channels[output] = make(chan string, c.size)
			}
		}
	}

	// Создаем каналы для мультиплексоров
	for _, m := range c.multiplexers {
		for _, input := range m.inputs {
			if _, ok := c.channels[input]; !ok {
				c.channels[input] = make(chan string, c.size)
			}
		}
		if _, ok := c.channels[m.output]; !ok {
			c.channels[m.output] = make(chan string, c.size)
		}
	}
}

// Run запускает конвейер
func (c *conveyerImpl) Run(ctx context.Context) error {
	c.mu.Lock()
	if c.running {
		c.mu.Unlock()
		return errors.New("conveyer already running")
	}
	c.running = true
	c.ctx, c.cancel = context.WithCancel(ctx)
	c.mu.Unlock()

	// Создаем все необходимые каналы
	c.createAllChannels()

	// Запускаем декораторы
	for _, d := range c.decorators {
		c.wg.Add(1)
		go func(d decoratorInfo) {
			defer c.wg.Done()

			inputCh, _ := c.getChannel(d.input)
			outputCh, _ := c.getChannel(d.output)

			if err := d.fn(c.ctx, inputCh, outputCh); err != nil {
				select {
				case c.errChan <- err:
				default:
				}
			}
		}(d)
	}

	// Запускаем сепараторы
	for _, s := range c.separators {
		c.wg.Add(1)
		go func(s separatorInfo) {
			defer c.wg.Done()

			inputCh, _ := c.getChannel(s.input)
			outputs := make([]chan string, len(s.outputs))
			for i, name := range s.outputs {
				outputs[i], _ = c.getChannel(name)
			}

			if err := s.fn(c.ctx, inputCh, outputs); err != nil {
				select {
				case c.errChan <- err:
				default:
				}
			}
		}(s)
	}

	// Запускаем мультиплексоры
	for _, m := range c.multiplexers {
		c.wg.Add(1)
		go func(m multiplexerInfo) {
			defer c.wg.Done()

			inputs := make([]chan string, len(m.inputs))
			for i, name := range m.inputs {
				inputs[i], _ = c.getChannel(name)
			}
			outputCh, _ := c.getChannel(m.output)

			if err := m.fn(c.ctx, inputs, outputCh); err != nil {
				select {
				case c.errChan <- err:
				default:
				}
			}
		}(m)
	}

	// Ждем завершения или ошибки
	go func() {
		c.wg.Wait()
		c.closeAllChannels()
		close(c.errChan)
	}()

	select {
	case <-c.ctx.Done():
		c.stop()
		return c.ctx.Err()
	case err := <-c.errChan:
		c.stop()
		return err
	}
}

// closeAllChannels закрывает все каналы
func (c *conveyerImpl) closeAllChannels() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for name, ch := range c.channels {
		close(ch)
		delete(c.channels, name)
	}
}

// stop останавливает конвейер
func (c *conveyerImpl) stop() {
	if c.cancel != nil {
		c.cancel()
	}
}

// Send отправляет данные в канал
func (c *conveyerImpl) Send(input string, data string) error {
	// Если конвейер не запущен, создаем канал на лету
	if !c.isRunning() {
		ch := c.getOrCreateChannel(input)
		ch <- data
		return nil
	}

	ch, err := c.getChannel(input)
	if err != nil {
		return err
	}

	select {
	case <-c.ctx.Done():
		return c.ctx.Err()
	case ch <- data:
		return nil
	}
}

// Recv получает данные из канала
func (c *conveyerImpl) Recv(output string) (string, error) {
	ch, err := c.getChannel(output)
	if err != nil {
		return "", err
	}

	select {
	case <-c.ctx.Done():
		return "", c.ctx.Err()
	case data, ok := <-ch:
		if !ok {
			return "undefined", nil
		}
		return data, nil
	}
}

// isRunning проверяет, запущен ли конвейер
func (c *conveyerImpl) isRunning() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.running
}
