package main

import (
	"flag"

	"github.com/Elektrek/currency-processor/internal/currencyio"
	"github.com/Elektrek/currency-processor/internal/settings"
)

func main() {
	cfgPath := flag.String("config", "", "Configuration file path")
	flag.Parse()

	if *cfgPath == "" {
		panic("Configuration file path is required")
	}

	config, err := settings.LoadConfig(*cfgPath)
	if err != nil {
		panic(err)
	}

	var rates currencyio.ExchangeRates
	rates.LoadFromXML(config.Source)
	rates.SortByRate()

	if err := currencyio.WriteJSONFile(config.Result, rates.Currencies); err != nil {
		panic(err)
	}
}
