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

	var envShort string

	switch cfg.Environment {
	case "production":
		envShort = "prod"
	case "development":
		envShort = "dev"
	default:
		envShort = cfg.Environment
	}

	fmt.Print(envShort, " ", cfg.LogLevel)
}
