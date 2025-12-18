package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/vikaglushkova/task-3/internal/config"
	"github.com/vikaglushkova/task-3/internal/currency"
	"github.com/vikaglushkova/task-3/internal/json"
)

const (
	defaultConfigPath = "config.yaml"
	dirPermissions    = 0o755
)

func main() {
	configPath := flag.String("config", defaultConfigPath, "Path to configuration file")
	flag.Parse()

	cfg, err := config.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	currencies, err := currency.ParseFromXMLFile(cfg.InputFile)
	if err != nil {
		log.Fatalf("Error parsing XML: %v", err)
	}

	sortedCurrencies := currency.ConvertAndSort(currencies)

	err = json.WriteCurrencyRateToFile(&sortedCurrencies, cfg.OutputFile, dirPermissions)
	if err != nil {
		log.Fatalf("Error saving JSON: %v", err)
	}

	fmt.Printf("Successfully processed %d currencies. Results saved to: %s\n", len(sortedCurrencies), cfg.OutputFile)
}
