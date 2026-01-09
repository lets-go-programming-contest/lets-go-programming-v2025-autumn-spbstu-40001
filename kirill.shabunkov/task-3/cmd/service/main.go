package main

import (
	"flag"

	"github.com/Kirill2155/task-3/internal/config"
	"github.com/Kirill2155/task-3/internal/json"
	"github.com/Kirill2155/task-3/internal/xml"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to yaml file")
	flag.Parse()

	config, err := config.ReadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	valCurs, err := xml.ParserXML(config.InputFile)
	if err != nil {
		panic(err)
	}

	valCurs.SortByValue()

	err = json.SaveJSON(config.OutputFile, valCurs.Valutes)
	if err != nil {
		panic(err)
	}
}
