package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var (
	ErrNoDecorator = errors.New("can't be decorated")
	ErrNoOutputs   = errors.New("no output channels")
)

const (
	Decorated     = "decorated: "
	NoDecorator   = "no decorator"
	NoMultiplexer = "no multiplexer"
)

func PrefixDecoratorFunc(ctx context.Context, inChan, outChan chan string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case val, ok := <-inChan:
			if !ok {
				return nil
			}

			if strings.Contains(val, NoDecorator) {
				return ErrNoDecorator
			}

			if !strings.HasPrefix(val, Decorated) {
				val = Decorated + val
			}

			select {
			case outChan <- val:
			case <-ctx.Done():
				return nil
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, inChan chan string, outChans []chan string) error {
	if len(outChans) == 0 {
		return ErrNoOutputs
	}

	index := 0

	for {
		select {
		case <-ctx.Done():
			return nil
		case val, ok := <-inChan:
			if !ok {
				return nil
			}

			target := outChans[index%len(outChans)]
			index++

			select {
			case target <- val:
			case <-ctx.Done():
				return nil
			}
		}
	}
}

func MultiplexerFunc(ctx context.Context, inChans []chan string, outChan chan string) error {
	if len(inChans) == 0 {
		return ErrNoOutputs
	}

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(inChans))

	for _, channel := range inChans {
		go func(inputChannel chan string) {
			defer waitGroup.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-inputChannel:
					if !ok {
						return
					}

					if strings.Contains(val, NoMultiplexer) {
						continue
					}

					select {
					case outChan <- val:
					case <-ctx.Done():
						return
					}
				}
			}
		}(channel)
	}

	waitGroup.Wait()

	return nil
}
