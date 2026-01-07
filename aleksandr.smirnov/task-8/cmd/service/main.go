package main

import (
	"fmt"

	"aleksandr.smirnov/task-8/internal/config"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		fmt.Printf("Config error: %v\n", err)

		return
	}

	cfg.Display()
}
