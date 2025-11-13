package main

import (
	"flag"
	"os"
	"sort"

	"github.com/Nekich06/task-3/internal/config"
	"github.com/Nekich06/task-3/internal/currencyrate"
	"github.com/Nekich06/task-3/internal/jsonparser"
	"github.com/Nekich06/task-3/internal/valutessorter"
	"github.com/Nekich06/task-3/internal/xmlparser"
	_ "github.com/paulrosania/go-charset/data"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")

	flag.Parse()

	configInfo, err := config.GetConfigInfo(configPath)
	if err != nil {
		panic(err)
	}

	var CBCurrencyRate currencyrate.CurrencyRate

	err = xmlparser.WriteInfoFromInputFileToCurrRate(configInfo.InputFile, &CBCurrencyRate)
	if err != nil {
		panic(err)
	}

	sort.Sort(valutessorter.ByValue(CBCurrencyRate))

	err = jsonparser.WriteInfoFromCurrRateToOutputFile(&CBCurrencyRate.Valutes, configInfo.OutputFile, os.FileMode(0o777))
	if err != nil {
		panic(err)
	}
}
