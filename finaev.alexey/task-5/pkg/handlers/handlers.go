package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var (
	decoratorErr = errors.New("can't be decorated")
	noOutputErr  = errors.New("no output channels")
	noInputErr   = errors.New("no input channels")
)

const (
	prefixMark      = "decorated: "
	decoratorMarker = "no decorator"
	skipMuxMarker   = "no multiplexer"
)

func PrefixDecoratorFunc(ctx context.Context, source, destination chan string) error {
	for item := range source {
		if strings.Contains(item, decoratorMarker) {
			return decoratorErr
		}

		if !strings.HasPrefix(item, prefixMark) {
			item = prefixMark + item
		}

		select {
		case destination <- item:
			continue
		case <-ctx.Done():
			return nil
		}
	}
	return nil
}

func SeparatorFunc(ctx context.Context, source chan string, destinations []chan string) error {
	if len(destinations) == 0 {
		return noOutputErr
	}

	for counter := 0; ; counter++ {
		select {
		case <-ctx.Done():
			return nil
		case item, active := <-source:
			if !active {
				return nil
			}

			targetIdx := counter % len(destinations)

			select {
			case destinations[targetIdx] <- item:
			case <-ctx.Done():
				return nil
			}
		}
	}
}

func MultiplexerFunc(ctx context.Context, sources []chan string, destination chan string) error {
	if len(sources) == 0 {
		return noInputErr
	}

	var workers sync.WaitGroup
	errorChan := make(chan error, 1)

	for _, src := range sources {
		workers.Add(1)

		go func(input chan string) {
			defer workers.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case item, active := <-input:
					if !active {
						return
					}

					if strings.Contains(item, skipMuxMarker) {
						continue
					}

					select {
					case destination <- item:
					case <-ctx.Done():
						return
					}
				}
			}
		}(src)
	}

	workers.Wait()
	close(errorChan)

	select {
	case err := <-errorChan:
		return err
	default:
		return nil
	}
}
