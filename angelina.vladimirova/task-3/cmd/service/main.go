package main

import (
	"flag"
	"sort"

	"task-3/internal/config"
	"task-3/internal/currency"
	"task-3/internal/json"
	"task-3/internal/xml"
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

	err = json.ParseJSON(cfg.OutputFilePath, currencyList.Data)
	if err != nil {
		panic(err)
	}
}
