package main

import (
	"flag"
	"log"

	"github.com/netwite/task-3/internal/config"
	"github.com/netwite/task-3/internal/currency"
)

func main() {
	configPath := flag.String("config", "internal/config/config.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := currency.ProcessValutes(cfg.InputFile, cfg.OutputFile); err != nil {
		log.Fatalf("Failed to process data: %v", err)
	}

}
