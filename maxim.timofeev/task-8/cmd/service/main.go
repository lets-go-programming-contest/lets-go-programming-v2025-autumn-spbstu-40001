package main

import (
	"fmt"

	"github.com/PigoDog/task-8/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	fmt.Print(cfg.Environment, " ", cfg.LogLevel)
}
