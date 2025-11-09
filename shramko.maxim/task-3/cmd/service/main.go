package main

import (
	"flag"
	"log"
	"os"

	"github.com/Elektrek/task-3/internal/config"
	"github.com/Elektrek/task-3/internal/processor"
)

func main() {
	defaultConfig := "config.yaml"
	if envConfig := os.Getenv("APP_CONFIG"); envConfig != "" {
		defaultConfig = envConfig
	}

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
