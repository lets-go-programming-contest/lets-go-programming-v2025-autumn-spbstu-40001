package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"

	"github.com/Rychmick/task-3/internal/config"
	"github.com/Rychmick/task-3/internal/currency"
)

const (
	RequiredArgsCount = 3
	RwOwnerROthers    = os.FileMode(0o644)
)

func LessValue(lhs, rhs currency.Currency) int {
	return cmp.Compare(lhs.Value, rhs.Value)
}

func main() {
	args := os.Args

	if (len(args) != RequiredArgsCount) || (args[1] != "--config") {
		fmt.Println("missing --config <file> in cmd args")

		return
	}

	settings, err := config.Parse(args[2])
	if err != nil {
		fmt.Println(err)

		return
	}

	currencyList, err := currency.ParseXML(settings.InputFilePath)
	if err != nil {
		fmt.Println(err)

		return
	}

	err = currency.Prepare(&currencyList)
	if err != nil {
		fmt.Println(err)

		return
	}

	slices.SortStableFunc(currencyList.Rates, LessValue)

	err = currency.ForceWriteAsJson(&currencyList, settings.OutputFilePath, RwOwnerROthers)
	if err != nil {
		fmt.Println(err)

		return
	}
}
