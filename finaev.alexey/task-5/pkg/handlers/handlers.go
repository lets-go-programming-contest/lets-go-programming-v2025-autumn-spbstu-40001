package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
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

			if strings.Contains(val, "no decorator") {
				return errors.New("error decorated!")
			}

			if !strings.HasPrefix(val, "decorated: ") {
				val = "decorated: " + val
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
	numOut := len(outChans)
	if numOut == 0 {
		return errors.New("no output channel")
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

			target := index % numOut
			index++

			select {
			case outChans[target] <- val:
			case <-ctx.Done():
				return nil
			}
		}
	}
}

func MultiplexerFunc(ctx context.Context, inChans []chan string, outChan chan string) error {
	numInput := len(inChans)
	if numInput == 0 {
		return errors.New("no input channel")
	}

	var waitGroup sync.WaitGroup

	waitGroup.Add(numInput)

	for _, channel := range inChans {
		go func(inChan chan string) {
			defer waitGroup.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-inChan:
					if !ok {
						return
					}

					if !strings.Contains(val, "no multiplexer") {
						select {
						case outChan <- val:
						case <-ctx.Done():
							return
						}
					}
				}
			}
		}(channel)
	}

	waitGroup.Wait()

	return nil
}
