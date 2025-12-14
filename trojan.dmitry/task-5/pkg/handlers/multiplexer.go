package handlers

import (
	"context"
	"strings"
	"sync"
)

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	defer close(output)
	if len(inputs) == 0 {
		return nil
	}

	inMerged := make(chan string, len(inputs)*10)
	var wg sync.WaitGroup

	for _, ch := range inputs {
		wg.Add(1)
		go func(cch chan string) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-cch:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case inMerged <- v:
					}
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(inMerged)
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case v, ok := <-inMerged:
			if !ok {
				return nil
			}
			if strings.Contains(v, "no multiplexer") {
				continue
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case output <- v:
			}
		}
	}
}
