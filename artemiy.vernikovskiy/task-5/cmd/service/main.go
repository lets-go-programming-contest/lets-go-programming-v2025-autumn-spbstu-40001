package main

import (
	"context"
	"log"
	"time"

	"github.com/Aapng-cmd/task-5/internal/conveyer"
	"github.com/Aapng-cmd/task-5/internal/handlers"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conveyerPipeline := conveyer.New(10)

	conveyerPipeline.RegisterDecorator(handlers.PrefixDecoratorFunc, "in", "out1")
	conveyerPipeline.RegisterSeparator(handlers.SeparatorFunc, "out1", []string{"sep1", "sep2"})
	conveyerPipeline.RegisterMultiplexer(handlers.MultiplexerFunc, []string{"sep1", "sep2"}, "final")

	go func() {
		err := conveyerPipeline.Run(ctx)
		if err != nil {
			log.Println("Pipeline error:", err)
		}
	}()

	inputs := []string{"hello", "world", "no decorator", "foo", "bar"}
	for _, data := range inputs {
		err := conveyerPipeline.Send("in", data)
		if err != nil {
			log.Println("Send error:", err)
		}
	}

	time.Sleep(1 * time.Second)

	for {
		val, err := conveyerPipeline.Recv("final")
		if err != nil {
			log.Println("Recv error:", err)

			break
		}

		if val == "undefined" {
			break
		}

		log.Println("Final:", val)
	}
}
