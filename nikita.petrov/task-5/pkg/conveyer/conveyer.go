package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

var errChanNotFound = errors.New("chan not found")

type Decorator struct {
	DecoratorFunc func(
		ctx context.Context,
		input chan string,
		output chan string,
	) error
	input  string
	output string
}

type Multiplexer struct {
	MultiplexerFunc func(
		ctx context.Context,
		inputs []chan string,
		output chan string,
	) error
	inputs []string
	output string
}

type Separator struct {
	SeparatorFunc func(
		ctx context.Context,
		input chan string,
		outputs []chan string,
	) error
	input   string
	outputs []string
}

type Conveyer struct {
	chansSize    int
	chansMap     sync.Map
	decorators   []*Decorator
	multiplexers []*Multiplexer
	separators   []*Separator
	mutex        sync.RWMutex
}

func (c *Conveyer) RegisterDecorator(
	newDecorator func(
		ctx context.Context,
		input chan string,
		output chan string,
	) error,
	input string,
	output string,
) {
	value, isExist := c.chansMap.Load(input)
	var inputChan chan string
	if !isExist {
		inputChan = make(chan string, c.chansSize)
		c.chansMap.Store(input, inputChan)
	} else {
		inputChan = value.(chan string)
	}

	outputChan, isExist := c.chansMap.Load(output)
	if !isExist {
		outputChan = make(chan string, c.chansSize)
		c.chansMap.Store(output, outputChan)
	}

	c.decorators = append(c.decorators, &Decorator{
		DecoratorFunc: newDecorator,
		input:         input,
		output:        output,
	})
}

func (c *Conveyer) RegisterMultiplexer(
	newMultiplexer func(
		ctx context.Context,
		inputs []chan string,
		output chan string,
	) error,
	inputs []string,
	output string,
) {
	inputChan, isExist := c.chansMap.Load(inputs)
	if !isExist {
		inputChan = make([]chan string, c.chansSize)
		c.chansMap.Store(inputs, inputChan)
	}

	outputChan, isExist := c.chansMap.Load(output)
	if !isExist {
		outputChan = make(chan string, c.chansSize)
		c.chansMap.Store(output, outputChan)
	}

	c.multiplexers = append(c.multiplexers, &Multiplexer{
		MultiplexerFunc: newMultiplexer,
		inputs:          inputs,
		output:          output,
	})
}

func (c *Conveyer) RegisterSeparator(
	newSeparator func(
		ctx context.Context,
		input chan string,
		outputs []chan string,
	) error,
	input string,
	outputs []string,
) {
	inputChan, isExist := c.chansMap.Load(input)
	if !isExist {
		inputChan = make([]chan string, c.chansSize)
		c.chansMap.Store(input, inputChan)
	}

	outputChan, isExist := c.chansMap.Load(outputs)
	if !isExist {
		outputChan = make([]chan string, c.chansSize)
		c.chansMap.Store(outputs, outputChan)
	}

	c.separators = append(c.separators, &Separator{
		SeparatorFunc: newSeparator,
		input:         input,
		outputs:       outputs,
	})
}

func (c *Conveyer) Run(ctx context.Context) error {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	var wg sync.WaitGroup
	var errCh = make(chan error, len(c.decorators)+len(c.multiplexers)+len(c.separators))
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	for _, decorator := range c.decorators {
		wg.Add(1)
		go func(d *Decorator) {
			defer wg.Done()

			inputChanValue, ok := c.chansMap.Load(d.input)
			if !ok {
				errCh <- errChanNotFound
				return
			}

			outputChanValue, ok := c.chansMap.Load(d.output)
			if !ok {
				errCh <- errChanNotFound
				return
			}

			inputChan := inputChanValue.(chan string)
			outputChan := outputChanValue.(chan string)

			if err := d.DecoratorFunc(ctx, inputChan, outputChan); err != nil {
				errCh <- fmt.Errorf("processing error: %w", err)
			}
		}(decorator)
	}

	for _, multiplexer := range c.multiplexers {
		wg.Add(1)
		go func(m *Multiplexer) {
			defer wg.Done()

			inputChans := make([]chan string, 0, len(m.inputs))
			for _, inputKey := range m.inputs {
				inputChanValue, ok := c.chansMap.Load(inputKey)
				if !ok {
					fmt.Errorf("Error")
					return
				}
				inputChans = append(inputChans, inputChanValue.(chan string))
			}

			outputChanValue, ok := c.chansMap.Load(m.output)
			if !ok {
				errCh <- errChanNotFound
				return
			}
			outputChan := outputChanValue.(chan string)

			if err := m.MultiplexerFunc(ctx, inputChans, outputChan); err != nil {
				errCh <- errChanNotFound
			}
		}(multiplexer)
	}

	for _, separator := range c.separators {
		wg.Add(1)
		go func(s *Separator) {
			defer wg.Done()

			inputChanValue, ok := c.chansMap.Load(s.input)
			if !ok {
				errCh <- errChanNotFound
				return
			}
			inputChan := inputChanValue.(chan string)

			outputChans := make([]chan string, 0, len(s.outputs))
			for _, outputKey := range s.outputs {
				outputChanValue, ok := c.chansMap.Load(outputKey)
				if !ok {
					errCh <- errChanNotFound
					return
				}
				outputChans = append(outputChans, outputChanValue.(chan string))
			}

			if err := s.SeparatorFunc(ctx, inputChan, outputChans); err != nil {
				errCh <- fmt.Errorf("processing error: %w", err)
			}
		}(separator)
	}

	done := make(chan struct{})
	var firstError error
	go func() {
		defer close(done)
		for err := range errCh {
			if firstError == nil {
				firstError = err
				cancel()
			}
		}
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	<-done

	return firstError
}

func (c *Conveyer) Send(input string, data string) error {
	value, ok := c.chansMap.Load(input)
	if !ok {
		return fmt.Errorf("Error")
	}

	ch := value.(chan string)
	select {
	case ch <- data:
		return nil
	default:
		return fmt.Errorf("Error")
	}
}

func (c *Conveyer) Recv(output string) (string, error) {
	value, ok := c.chansMap.Load(output)
	if !ok {
		return "", fmt.Errorf("Error")
	}

	ch := value.(chan string)
	select {
	case data := <-ch:
		return data, nil
	default:
		return "", fmt.Errorf("Error")
	}
}

func New(chansSize int) Conveyer {
	return Conveyer{
		chansSize:    chansSize,
		chansMap:     sync.Map{},
		decorators:   make([]*Decorator, 0),
		multiplexers: make([]*Multiplexer, 0),
		separators:   make([]*Separator, 0),
		mutex:        sync.RWMutex{},
	}
}
