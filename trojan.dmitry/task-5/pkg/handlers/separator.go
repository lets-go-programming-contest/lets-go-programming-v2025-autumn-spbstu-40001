package handlers

import (
	"context"
)

func SeparatorFunc(
	ctx context.Context,
	input chan string,
	outputs []chan string,
) error {
	if len(outputs) == 0 {
		return ErrEmptyChannel
	}

	var index int

	for {
		select {
		case <-ctx.Done():
			return nil
		case message, ok := <-input:
			if !ok {
				return nil
			}

			targetChannel := outputs[index]
			index = (index + 1) % len(outputs)

			select {
			case targetChannel <- message:
			case <-ctx.Done():
				return nil
			}
		}
	}
}
