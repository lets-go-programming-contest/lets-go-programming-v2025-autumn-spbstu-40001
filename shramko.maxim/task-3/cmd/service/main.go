package main

import (
	"flag"

	"github.com/Elektrek/task-3/internal/config"
	"github.com/Elektrek/task-3/internal/iocurrency"
)

func main() {
	configFile := flag.String("config", "", "YAML configuration file path")
	flag.Parse()

	if *configFile == "" {
		panic("Configuration file not specified")
	}

	settings, err := config.LoadSettings(*configFile)
	if err != nil {
		panic(err)
	}

	var currencyData iocurrency.CurrencyList

	currencyData.ParseXML(settings.InputPath)
	currencyData.OrderByValue()

	if err = iocurrency.ExportJSON(settings.OutputPath, currencyData.Items); err != nil {
		panic(err)
	}
}
