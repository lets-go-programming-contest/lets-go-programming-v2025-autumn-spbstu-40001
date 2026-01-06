package main

import (
	"fmt"

	"github.com/atroxxxxxx/task-/internal/config"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	fmt.Print(conf.Environment, " ", conf.LogLevel)
}
