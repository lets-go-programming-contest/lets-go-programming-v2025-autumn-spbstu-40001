package main

import (
	"flag"
	"fmt"

	"github.com/Danil3352/task-3/internal/config"
	"github.com/Danil3352/task-3/internal/currency"
)

func main() {
	configPath := flag.String("config", "internal/config/config.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.NewConfig(*configPath)
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	if err := currency.Process(cfg.InputFile, cfg.OutputFile); err != nil {
		panic(fmt.Errorf("failed to process fata: %w", err))
	}
}
