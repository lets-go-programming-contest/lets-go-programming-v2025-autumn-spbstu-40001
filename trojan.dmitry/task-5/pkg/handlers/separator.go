package handlers

import (
	"context"
)

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
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

			for attempts := 0; attempts < len(outputs); attempts++ {
				currentIdx := (idx + attempts) % len(outputs)
				select {
				case <-ctx.Done():
					return ctx.Err()
				case outputs[currentIdx] <- v:
					idx = (currentIdx + 1) % len(outputs)
					goto next
				default:
					continue
				}
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			case outputs[idx%len(outputs)] <- v:
			}

		next:
		}
	}
}
