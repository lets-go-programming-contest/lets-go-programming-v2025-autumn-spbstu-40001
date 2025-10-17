package main

import (
	"cmp"
	"flag"
	"os"
	"slices"

	"github.com/Rychmick/task-3/internal/config"
	"github.com/Rychmick/task-3/internal/currency"
	"github.com/Rychmick/task-3/internal/json"
	"github.com/Rychmick/task-3/internal/xml"
)

const defaultFileMode = os.FileMode(0o666)

func compareValues(lhs, rhs currency.Currency) int {
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

	var currencyList currency.Rates

	err = xml.ParseFile(settings.InputFilePath, &currencyList)
	if err != nil {
		panic(err)
	}

	slices.SortStableFunc(currencyList.Data, compareValues)

	err = json.WriteFile(settings.OutputFilePath, currencyList.Data, defaultFileMode)
	if err != nil {
		panic(err)
	}
}
