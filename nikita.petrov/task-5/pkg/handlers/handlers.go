package handlers

import (
	"context"
	"fmt"
	"strings"
)

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	data := <-input
	if strings.Contains(data, "no decorator") {
		return fmt.Errorf("can't be decorated")
	}

	result := "decorated: " + data
	output <- result

	return nil
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {

}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) {

}
