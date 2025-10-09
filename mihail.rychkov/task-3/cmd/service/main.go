package main

import (
	"cmp"
	"flag"
	"fmt"
	"os"
	"slices"

	"github.com/Rychmick/task-3/internal/config"
	"github.com/Rychmick/task-3/internal/currency"
)

const DefaultFileMode = os.FileMode(0o666)

func CompareValues(lhs, rhs currency.Currency) int {
	return -cmp.Compare(lhs.Value, rhs.Value)
}

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "config.yaml", "set YAML settings file")
	flag.Parse()

	settings, err := config.Parse(configPath)
	if err != nil {
		panic(err)
	}

	currencyList, err := currency.ParseXML(settings.InputFilePath)
	if err != nil {
		panic(err)
	}

	err = currency.Prepare(&currencyList)
	if err != nil {
		fmt.Println(err)

		return
	}

	slices.SortStableFunc(currencyList.Rates, CompareValues)

	err = currency.ForceWriteToJSON(&currencyList, settings.OutputFilePath, DefaultFileMode)
	if err != nil {
		fmt.Println(err)

		return
	}
}
