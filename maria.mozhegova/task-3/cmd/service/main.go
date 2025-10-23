package main

import (
	"flag"

	"github.com/mipxe/task-3/internal/config"
	"github.com/mipxe/task-3/internal/currency"
)

func main() {
	configPath := flag.String("config", "", "path to yaml file")
	flag.Parse()

	if *configPath == "" {
		panic("config flag is required")
	}

	config, err := config.ReadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	valCurs, err := currency.ReadValCurs(config.InputFile)
	if err != nil {
		panic(err)
	}

	valCurs.SortByValueDesc()

	err = currency.WriteToJSON(valCurs, config.OutputFile)
	if err != nil {
		panic(err)
	}
}
