package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var ErrEmptyChannel = errors.New("channel can't be empty")

const NoMultiplexer = "no multiplexer"

func MultiplexerFunc(
	ctx context.Context,
	inputs []chan string,
	output chan string,
) error {
	if len(inputs) == 0 {
		return ErrEmptyChannel
	}

	var waitGrp sync.WaitGroup

	waitGrp.Add(len(inputs))

	for _, channel := range inputs {
		go func(inputChannel chan string) {
			defer waitGrp.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case message, ok := <-inputChannel:
					if !ok {
						return
					}

					if strings.Contains(message, NoMultiplexer) {
						continue
					}

					select {
					case output <- message:
					case <-ctx.Done():
						return
					}
				}
			}
		}(channel)
	}

	waitGrp.Wait()

	return nil
}
