package conveyer

import (
	"context"
	"errors"
	"sync"
)

// Ошибки конвейера
var (
	ErrChanNotFound  = errors.New("chan not found")
	ErrAlreadyExists = errors.New("channel already exists")
	ErrNotRunning    = errors.New("conveyer not running")
)

// Типы обработчиков
type DecoratorFunc func(ctx context.Context, input chan string, output chan string) error
type MultiplexerFunc func(ctx context.Context, inputs []chan string, output chan string) error
type SeparatorFunc func(ctx context.Context, input chan string, outputs []chan string) error

// Структура конвейера
type Conveyer struct {
	mu           sync.RWMutex
	running      bool
	size         int
	channels     map[string]chan string
	decorators   []decorator
	multiplexers []multiplexer
	separators   []separator
	ctx          context.Context
	cancel       context.CancelFunc
	wg           sync.WaitGroup
	errors       chan error
}

type decorator struct {
	fn     DecoratorFunc
	input  string
	output string
}

type multiplexer struct {
	fn     MultiplexerFunc
	inputs []string
	output string
}

type separator struct {
	fn      SeparatorFunc
	input   string
	outputs []string
}

// New создает новый конвейер
func New(size int) *Conveyer {
	return &Conveyer{
		size:     size,
		channels: make(map[string]chan string),
		errors:   make(chan error, 10),
	}
}

// RegisterDecorator регистрирует модификатор данных
func (c *Conveyer) RegisterDecorator(
	fn DecoratorFunc,
	input string,
	output string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.decorators = append(c.decorators, decorator{
		fn:     fn,
		input:  input,
		output: output,
	})
}

// RegisterMultiplexer регистрирует мультиплексор
func (c *Conveyer) RegisterMultiplexer(
	fn MultiplexerFunc,
	inputs []string,
	output string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.multiplexers = append(c.multiplexers, multiplexer{
		fn:     fn,
		inputs: inputs,
		output: output,
	})
}

// RegisterSeparator регистрирует сепаратор
func (c *Conveyer) RegisterSeparator(
	fn SeparatorFunc,
	input string,
	outputs []string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.separators = append(c.separators, separator{
		fn:      fn,
		input:   input,
		outputs: outputs,
	})
}

// getOrCreateChannel получает или создает канал
func (c *Conveyer) getOrCreateChannel(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ch, exists := c.channels[name]; exists {
		return ch
	}

	ch := make(chan string, c.size)
	c.channels[name] = ch
	return ch
}

// getChannel получает канал по имени
func (c *Conveyer) getChannel(name string) (chan string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	ch, exists := c.channels[name]
	if !exists {
		return nil, ErrChanNotFound
	}

	return ch, nil
}

// Run запускает конвейер
func (c *Conveyer) Run(ctx context.Context) error {
	c.mu.Lock()
	if c.running {
		c.mu.Unlock()
		return errors.New("conveyer already running")
	}

	c.ctx, c.cancel = context.WithCancel(ctx)
	c.running = true
	c.mu.Unlock()

	// Запускаем все обработчики
	for _, d := range c.decorators {
		inputCh := c.getOrCreateChannel(d.input)
		outputCh := c.getOrCreateChannel(d.output)

		c.wg.Add(1)
		go func(d decorator, input, output chan string) {
			defer c.wg.Done()
			if err := d.fn(c.ctx, input, output); err != nil {
				select {
				case c.errors <- err:
				default:
				}
			}
			close(output)
		}(d, inputCh, outputCh)
	}

	for _, m := range c.multiplexers {
		inputs := make([]chan string, len(m.inputs))
		for i, name := range m.inputs {
			inputs[i] = c.getOrCreateChannel(name)
		}
		outputCh := c.getOrCreateChannel(m.output)

		c.wg.Add(1)
		go func(m multiplexer, inputs []chan string, output chan string) {
			defer c.wg.Done()
			if err := m.fn(c.ctx, inputs, output); err != nil {
				select {
				case c.errors <- err:
				default:
				}
			}
			close(output)
		}(m, inputs, outputCh)
	}

	for _, s := range c.separators {
		inputCh := c.getOrCreateChannel(s.input)
		outputs := make([]chan string, len(s.outputs))
		for i, name := range s.outputs {
			outputs[i] = c.getOrCreateChannel(name)
		}

		c.wg.Add(1)
		go func(s separator, input chan string, outputs []chan string) {
			defer c.wg.Done()
			if err := s.fn(c.ctx, input, outputs); err != nil {
				select {
				case c.errors <- err:
				default:
				}
			}
			for _, ch := range outputs {
				close(ch)
			}
		}(s, inputCh, outputs)
	}

	// Ожидаем завершения или ошибки
	select {
	case <-c.ctx.Done():
		c.stop()
		return c.ctx.Err()
	case err := <-c.errors:
		c.stop()
		return err
	}
}

// stop останавливает конвейер
func (c *Conveyer) stop() {
	c.mu.Lock()
	if !c.running {
		c.mu.Unlock()
		return
	}

	c.cancel()
	c.running = false
	c.mu.Unlock()

	// Закрываем все каналы
	c.mu.Lock()
	for name, ch := range c.channels {
		select {
		case <-ch:
		default:
			close(ch)
		}
		delete(c.channels, name)
	}
	c.mu.Unlock()

	c.wg.Wait()
	close(c.errors)
}

// Send отправляет данные в канал
func (c *Conveyer) Send(input string, data string) error {
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
func (c *Conveyer) Recv(output string) (string, error) {
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

// Wait ожидает завершения всех горутин
func (c *Conveyer) Wait() {
	c.wg.Wait()
}
