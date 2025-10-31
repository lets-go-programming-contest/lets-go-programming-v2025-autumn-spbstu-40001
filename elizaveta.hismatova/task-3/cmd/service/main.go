package main

import (
	"flag"
	"fmt"

	"github.com/LeeLisssa/task-3/internal/config"
	"github.com/LeeLisssa/task-3/internal/jsonwriter"
	"github.com/LeeLisssa/task-3/internal/types"
	"github.com/LeeLisssa/task-3/internal/xmlparser"
)

func main() {
	configPath := flag.String("config", "", "Path to configuration file")
	flag.Parse()

	if *configPath == "" {
		panic("config flag is required")
	}

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	var valCurs types.ValCurs
	if err := xmlparser.ParseXML(cfg.InputFile, &valCurs); err != nil {
		panic(fmt.Sprintf("Failed to parse XML: %v", err))
	}

	sortedCurrencies := valCurs.SortByValueDesc()

	if err := jsonwriter.WriteJSON(cfg.OutputFile, sortedCurrencies); err != nil {
		panic(fmt.Sprintf("Failed to write JSON: %v", err))
	}
}

