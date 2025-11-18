package main

import (
	"flag"
	"os"

	"github.com/Tuc0Sa1amanka/task-3/internal/config"
	"github.com/Tuc0Sa1amanka/task-3/internal/jsonwriter"
	"github.com/Tuc0Sa1amanka/task-3/internal/valcurs"
	"github.com/Tuc0Sa1amanka/task-3/internal/xmlparser"
)

func main() {
	configPath := flag.String("config", "example/config.yaml", "Path to yaml file")
	flag.Parse()

	config := new(config.Config)

	if err := config.LoadFromFile(*configPath); err != nil {
		panic(err)
	}

	inputFile, err := os.ReadFile(config.InputFilePath)
	if err != nil {
		panic(err)
	}

	valCurs, err := xmlparser.ParseXML[valcurs.ValCurs](inputFile)
	if err != nil {
		panic(err)
	}

	valCurs.SortByValueDown()

	if err := jsonwriter.SaveToJSON(valCurs.Valutes, config.OutputFilePath); err != nil {
		panic(err)
	}
}
