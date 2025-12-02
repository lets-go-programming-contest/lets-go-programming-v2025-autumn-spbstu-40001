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
	decorators   []*decoratorInfo
	multiplexers []*multiplexerInfo
	separators   []*separatorInfo
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
		errChan:  make(chan error, 1),
	}
}

// RegisterDecorator регистрирует модификатор данных
func (c *conveyerImpl) RegisterDecorator(fn DecoratorFunc, input string, output string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.decorators = append(c.decorators, &decoratorInfo{
		fn:     fn,
		input:  input,
		output: output,
	})
}

// RegisterMultiplexer регистрирует мультиплексор
func (c *conveyerImpl) RegisterMultiplexer(fn MultiplexerFunc, inputs []string, output string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.multiplexers = append(c.multiplexers, &multiplexerInfo{
		fn:     fn,
		inputs: inputs,
		output: output,
	})
}

// RegisterSeparator регистрирует сепаратор
func (c *conveyerImpl) RegisterSeparator(fn SeparatorFunc, input string, outputs []string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.separators = append(c.separators, &separatorInfo{
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

	// Запускаем декораторы
	for _, d := range c.decorators {
		inputCh := c.getOrCreateChannel(d.input)
		outputCh := c.getOrCreateChannel(d.output)

		c.wg.Add(1)
		go func(d *decoratorInfo, in, out chan string) {
			defer c.wg.Done()
			if err := d.fn(c.ctx, in, out); err != nil {
				select {
				case c.errChan <- err:
				default:
				}
			}
			close(out)
		}(d, inputCh, outputCh)
	}

	// Запускаем сепараторы
	for _, s := range c.separators {
		inputCh := c.getOrCreateChannel(s.input)
		outputs := make([]chan string, len(s.outputs))
		for i, name := range s.outputs {
			outputs[i] = c.getOrCreateChannel(name)
		}

		c.wg.Add(1)
		go func(s *separatorInfo, in chan string, outs []chan string) {
			defer c.wg.Done()
			if err := s.fn(c.ctx, in, outs); err != nil {
				select {
				case c.errChan <- err:
				default:
				}
			}
			for _, ch := range outs {
				close(ch)
			}
		}(s, inputCh, outputs)
	}

	// Запускаем мультиплексоры
	for _, m := range c.multiplexers {
		inputs := make([]chan string, len(m.inputs))
		for i, name := range m.inputs {
			inputs[i] = c.getOrCreateChannel(name)
		}
		outputCh := c.getOrCreateChannel(m.output)

		c.wg.Add(1)
		go func(m *multiplexerInfo, ins []chan string, out chan string) {
			defer c.wg.Done()
			if err := m.fn(c.ctx, ins, out); err != nil {
				select {
				case c.errChan <- err:
				default:
				}
			}
			close(out)
		}(m, inputs, outputCh)
	}

	// Ждем завершения или ошибки
	go func() {
		c.wg.Wait()
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

// stop останавливает конвейер
func (c *conveyerImpl) stop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.running {
		return
	}

	if c.cancel != nil {
		c.cancel()
	}

	// Закрываем все каналы
	for name, ch := range c.channels {
		close(ch)
		delete(c.channels, name)
	}

	c.running = false
}

// Send отправляет данные в канал
func (c *conveyerImpl) Send(input string, data string) error {
	ch, err := c.getChannel(input)
	if err != nil {
		return err
	}

	select {
	case <-c.ctx.Done():
		return c.ctx.Err()
	case ch <- data:
		return nil
	default:
		return errors.New("channel buffer full")
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
