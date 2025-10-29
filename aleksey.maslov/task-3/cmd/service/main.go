package main

import (
	"flag"

	"github.com/A1exMas1ov/task-3/internal/config"
	"github.com/A1exMas1ov/task-3/internal/currency"
	"github.com/A1exMas1ov/task-3/internal/jsonwriter"
	"github.com/A1exMas1ov/task-3/internal/xmlparser"
)

func main() {
	configPath := flag.String("config", "configs/config.yaml", "path to config file")
	flag.Parse()

	config, err := config.ReadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	var valCurs currency.ValCurs

	err = xmlparser.ParseXML(config.InputFile, &valCurs)
	if err != nil {
		panic(err)
	}

	valCurs.SortByValue()

	err = jsonwriter.SaveJSON(config.OutputFile, valCurs.Valutes)
	if err != nil {
		panic(err)
	}
}
