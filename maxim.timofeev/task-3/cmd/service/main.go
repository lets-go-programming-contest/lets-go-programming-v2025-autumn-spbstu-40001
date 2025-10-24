package main

import (
	"flag"

	"github.com/PigoDog/task-3/internal/config"
	"github.com/PigoDog/task-3/internal/iocurrency"
)

func main() {
	configPath := flag.String("config", "", "Path to YAML config")
	flag.Parse()

	if *configPath == "" {
		panic("flag --config is empty")
	}

	config, err := config.ReadConfig(*configPath)
	if err != nil {
		panic(err.Error())
	}

	var valutes iocurrency.ValCurs

	valutes.ReadXML(config.InputFile)
	valutes.Sort()

	if err = iocurrency.SaveJSON(config.OutputFile, valutes.Valutes); err != nil {
		panic(err.Error())
	}
}
