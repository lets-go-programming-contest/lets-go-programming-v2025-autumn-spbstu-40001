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

			sent := false
			for attempts := 0; attempts < len(outputs); attempts++ {
				currentIdx := (idx + attempts) % len(outputs)
				select {
				case <-ctx.Done():
					return ctx.Err()
				case outputs[currentIdx] <- v:
					sent = true
					idx = (currentIdx + 1) % len(outputs)
					break
				default:
					continue
				}
				if sent {
					break
				}
			}

			if !sent {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case outputs[idx%len(outputs)] <- v:
				}
				idx = (idx + 1) % len(outputs)
			}
		}
	}
}
