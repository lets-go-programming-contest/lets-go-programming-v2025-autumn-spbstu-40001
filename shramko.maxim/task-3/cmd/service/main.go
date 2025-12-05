package main

import (
	"flag"
	"log"

	"github.com/Elektrek/task-3/internal/config"
	"github.com/Elektrek/task-3/internal/processor"
)

const defaultConfig = "config.yaml"

func main() {
	configPath := flag.String("config", defaultConfig, "YAML configuration file path")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := processor.ProcessCurrencies(cfg); err != nil {
		log.Fatalf("Failed to process currencies: %v", err)
	}
}
