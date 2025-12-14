package handlers

import (
	"context"
)

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	defer func() {
		for _, out := range outputs {
			close(out)
		}
	}()

	if len(outputs) == 0 {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case _, ok := <-input:
				if !ok {
					return nil
				}
			}
		}
	}

	idx := 0
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case v, ok := <-input:
			if !ok {
				return nil
			}

			for i := 0; i < len(outputs); i++ {
				currentIdx := (idx + i) % len(outputs)
				select {
				case <-ctx.Done():
					return ctx.Err()
				case outputs[currentIdx] <- v:
					break
				default:
					continue
				}
				break
			}

			idx = (idx + 1) % len(outputs)
		}
	}
}
