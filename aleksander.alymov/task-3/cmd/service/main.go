package main

import (
	"flag"
	"log"

	"github.com/netwite/task-3/internal/config"
	"github.com/netwite/task-3/internal/currency"
	"github.com/netwite/task-3/internal/json"
	"github.com/netwite/task-3/internal/processor"
	"github.com/netwite/task-3/internal/sorter"
	"github.com/netwite/task-3/internal/xml"
)

func main() {
	configPath := flag.String("config", "config/config.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	xmlLoader := xml.NewLoader()
	currencyConverter := currency.NewConverter()
	descSorter := sorter.NewDescendingSorter()
	jsonSaver := json.NewSaver()

	dataProcessor := processor.NewDataProcessor(
		xmlLoader,
		currencyConverter,
		descSorter,
		jsonSaver,
	)

	if err := dataProcessor.Process(cfg.InputFile, cfg.OutputFile); err != nil {
		log.Fatalf("Failed to process data: %v", err)
	}
}
