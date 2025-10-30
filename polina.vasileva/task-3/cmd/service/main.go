package main

import (
	"flag"
	"sort"

	"polina.vasileva/task-3/internal/config"
	"polina.vasileva/task-3/internal/currency"
	"polina.vasileva/task-3/internal/json"
	"polina.vasileva/task-3/internal/xml"
)

const (
	dirPermissions  = 0o755
	filePermissions = 0o600
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config")
	flag.Parse()

	config, err := config.ParseYaml(*configPath)
	if err != nil {
		panic(err)
	}

	currencyList := currency.Rates{Data: []currency.Currency{}}

	err = xml.ParseXML(config.InputFilePath, &currencyList)
	if err != nil {
		panic(err)
	}

	sort.Slice(currencyList.Data, func(i, j int) bool {
		return currencyList.Data[i].Value > currencyList.Data[j].Value
	})

	err = json.ParseJSON(config.OutputFilePath, currencyList.Data, dirPermissions, filePermissions)
	if err != nil {
		panic(err)
	}
}
