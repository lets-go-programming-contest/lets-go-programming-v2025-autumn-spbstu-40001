package main

import (
	"flag"
	"log"

	"github.com/netwite/task-3/internal/config"
	"github.com/netwite/task-3/internal/processor"
)

func main() {
	configPath := flag.String("config", "config/config.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	proc := processor.NewProcessor()

	if err := proc.Process(cfg.InputFile, cfg.OutputFile); err != nil {
		log.Fatalf("Failed to process data: %v", err)
	}
}
