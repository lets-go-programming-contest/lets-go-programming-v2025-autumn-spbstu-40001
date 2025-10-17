package main

import (
	"cmp"
	"flag"
	"os"
	"slices"

	"github.com/Rychmick/task-3/internal/config"
	"github.com/Rychmick/task-3/internal/currency"
)

func CompareValues(lhs, rhs currency.Currency) int {
	return -cmp.Compare(lhs.Value, rhs.Value)
}

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "config.yaml", "path to config file")
	flag.Parse()

	settings, err := config.Parse(configPath)
	if err != nil {
		panic(err)
	}

	currencyList, err := currency.ParseXML(settings.InputFilePath)
	if err != nil {
		panic(err)
	}

	slices.SortStableFunc(currencyList.Rates, CompareValues)

	err = currency.WriteToJSON(&currencyList, settings.OutputFilePath, os.FileMode(0o666))
	if err != nil {
		panic(err)
	}
}
