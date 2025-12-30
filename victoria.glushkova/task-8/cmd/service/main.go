package main

import (
	"fmt"
	"log"

	"victoria.glushkova/task-8/internal/config"
)

func main() {
	loader := config.NewLoader()

	cfg, err := config.GetConfig(loader)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	fmt.Printf("%s %s\n", cfg.Environment, cfg.LogLevel)
}
