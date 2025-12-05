package handlers

import (
	"context"
	"errors"
	"strings"
	"sync/atomic"
)

func AddPrefixDecorator(ctx context.Context, in <-chan string, out chan<- string) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-in:
			if !ok {
				return nil
			}

			if strings.Contains(data, "no decorator") {
				return errors.New("can't be decorated")
			}

			const marker = "decorated: "
			if !strings.HasPrefix(data, marker) {
				data = marker + data
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			case out <- data:
			}
		}
	}
}

func SequentialSeparator(ctx context.Context, in <-chan string, outs []chan<- string) error {
	var counter uint64

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-in:
			if !ok {
				return nil
			}

			idx := atomic.AddUint64(&counter, 1) % uint64(len(outs))

			select {
			case <-ctx.Done():
				return ctx.Err()
			case outs[idx] <- data:
			}
		}
	}
}

func SelectiveMultiplexer(ctx context.Context, ins []<-chan string, out chan<- string) error {
	type received struct {
		value string
		valid bool
	}

	merged := make(chan received, len(ins))

	for _, input := range ins {
		go func(in <-chan string) {
			for {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-in:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case merged <- received{val, ok}:
					}
				}
			}
		}(input)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case item := <-merged:
			if !item.valid {
				continue
			}

			if strings.Contains(item.value, "no multiplexer") {
				continue
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			case out <- item.value:
			}
		}
	}
}
