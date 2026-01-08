package main

import (
	"flag"
	"fmt"

	"github.com/A1exCRE/task-3/internal/config"
	"github.com/A1exCRE/task-3/internal/currency"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		fmt.Printf("Config error: %v\n", err)
		return
	}

	fmt.Printf("Loading XML from: %s\n", cfg.Input)

	bankData, err := currency.LoadFromFile(cfg.Input)
	if err != nil {
		fmt.Printf("XML error: %v\n", err)
		return
	}

	converted, err := currency.ConvertValues(bankData)
	if err != nil {
		fmt.Printf("Conversion error: %v\n", err)
		return
	}
	converted.SortDesc()

	if err := converted.WriteJSONFile(cfg.Output); err != nil {
		fmt.Printf("Save error: %v\n", err)
		return
	}

	fmt.Printf("Success! Saved %d currencies to: %s\n", len(converted), cfg.Output)
}
