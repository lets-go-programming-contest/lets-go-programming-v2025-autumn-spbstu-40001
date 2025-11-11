package main

import (
	"flag"

	"github.com/PigoDog/task-3/internal/config"
	"github.com/PigoDog/task-3/internal/iocurrency"
	"github.com/PigoDog/task-3/internal/json"
	"github.com/PigoDog/task-3/internal/xml"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to YAML config")
	flag.Parse()

	config, err := config.ReadConfig(*configPath)
	if err != nil {
		panic(err.Error())
	}

	var valutes iocurrency.ValCurs

	xml.ReadXML(config.InputFile, valutes)
	valutes.Sort()

	if err := json.SaveJSON(config.OutputFile, valutes.Valutes); err != nil {
		panic(err.Error())
	}
}
