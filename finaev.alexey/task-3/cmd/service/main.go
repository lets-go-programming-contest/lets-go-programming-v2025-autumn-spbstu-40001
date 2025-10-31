package main

import (
	"flag"

	"github.com/AlexeyFinaev02/task-3/internal/config"
	"github.com/AlexeyFinaev02/task-3/internal/jsonwriter"
	"github.com/AlexeyFinaev02/task-3/internal/valcurs"
	"github.com/AlexeyFinaev02/task-3/internal/xmlparser"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to yaml file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err.Error())
	}

	var curs valcurs.Currency

	err = xmlparser.LoadCurrencies(cfg.InputFile, &curs)
	if err != nil {
		panic(err.Error())
	}

	curs.SortCurrenciesByValueDesc()

	err = jsonwriter.SaveJSON(cfg.OutputFile, curs.Currencies)
	if err != nil {
		panic(err.Error())
	}
}
