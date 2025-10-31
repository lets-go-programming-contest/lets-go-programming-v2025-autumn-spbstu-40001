package main

import (
	"flag"
	"fmt"

	"github.com/AlexeyFinaev02/task-3/internal/config"
	"github.com/AlexeyFinaev02/task-3/internal/jsonwriter"
	"github.com/AlexeyFinaev02/task-3/internal/xmlparser"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to yaml file")
	flag.Parse()
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(fmt.Sprintf("Error read config: %v", err))
	}
	valCurs, err := xmlparser.LoadCurrencies(cfg.InputFile)
	if err != nil {
		panic(fmt.Sprintf("Error load currencies: %v", err))
	}

	valCurs.SortCurrenciesByValueDesc()

	err = jsonwriter.SaveJSON(cfg.OutputFile, valCurs.Currencies)
	if err != nil {
		panic(fmt.Sprintf("Error saving JSON: %v", err))
	}

	fmt.Printf("Successfully processed %d currencies. Result saved to: %s\n",
		len(valCurs.Currencies), cfg.OutputFile)
}
