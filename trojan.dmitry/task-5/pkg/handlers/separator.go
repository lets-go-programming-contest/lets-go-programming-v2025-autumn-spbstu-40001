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
				out := outputs[idx%len(outputs)]
				select {
				case <-ctx.Done():
					return ctx.Err()
				case out <- v:
					goto next
				default:
					idx = (idx + 1) % len(outputs)
				}
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			case outputs[idx%len(outputs)] <- v:
			}

		next:
			idx = (idx + 1) % len(outputs)
		}
	}
}
