package main

import (
	"flag"
	"sort"

	"github.com/verticalochka/task-3/internal/config"
	"github.com/verticalochka/task-3/internal/currency"
	"github.com/verticalochka/task-3/internal/json"
	"github.com/verticalochka/task-3/internal/xml"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config")
	flag.Parse()

	cfg, err := config.ParseYaml(*configPath)
	if err != nil {
		panic(err)
	}

	var currencyList currency.Rates

	err = xml.ParseXML(cfg.InputFilePath, &currencyList)
	if err != nil {
		panic(err)
	}

	sort.Slice(currencyList.Data, func(i, j int) bool {
		return currencyList.Data[i].Value > currencyList.Data[j].Value
	})

	err = json.ParseJSON(cfg.OutputFilePath, currencyList.Data, 0o755, 0o600)
	if err != nil {
		panic(err)
	}
}
