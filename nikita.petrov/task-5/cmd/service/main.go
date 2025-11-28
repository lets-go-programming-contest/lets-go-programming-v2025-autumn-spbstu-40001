package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Nekich06/task-5/pkg/conveyer"
	"github.com/Nekich06/task-5/pkg/handlers"
	"golang.org/x/sync/errgroup"
)

func ReadAll(name string, recv func(name string) (string, error)) ([]string, error) {
	results := []string{}

	for {
		data, err := recv(name)
		if err != nil {
			return nil, err
		}

		results = append(results, data)
		if data == undefinedData {
			break
		}
	}

	return results, nil
}

const (
	chansSize          = 256
	undefinedData      = "undefined"
	chanNotExitMessage = "chan not found"
)

func main() {
	stringsProcessor := conveyer.New(chansSize)
	stringsProcessor.RegisterDecorator(handlers.PrefixDecoratorFunc, "input", "d1-output")
	stringsProcessor.RegisterDecorator(handlers.PrefixDecoratorFunc, "d1-output", "output")

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()

	errgr, ctx := errgroup.WithContext(ctx)
	errgr.Go(func() error {
		return stringsProcessor.Run(ctx)
	})

	stringsProcessor.Send("input", "val-1")
	stringsProcessor.Send("input", "decorated: val-2")

	errgr.Wait()

	output, err := ReadAll("output", stringsProcessor.Recv)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
