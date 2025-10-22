package main

import (
	"flag"
	"os"

	"github.com/Tuc0Sa1amanka/task-3/internal/config"
	"github.com/Tuc0Sa1amanka/task-3/internal/jsonwriter"
	"github.com/Tuc0Sa1amanka/task-3/internal/valcurs"
)

func main() {
	configPath := flag.String("config", "", "Path to yaml file")
	flag.Parse()

	config := new(config.Config)
	config.LoadFromFile(*configPath)

	inputFile, err := os.ReadFile(config.InputFilePath)
	if err != nil {
		panic(err)
	}

	valCurs := new(valcurs.ValCurs)
	valCurs.ParseXML(inputFile)
	valCurs.SortByValueDown()

	jsonwriter.SaveToJSON(valCurs.Valutes, config.OutputFilePath)
}
