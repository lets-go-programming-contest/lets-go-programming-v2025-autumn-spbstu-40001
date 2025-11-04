package main

import (
	"flag"

	"github.com/A1exMas1ov/task-3/internal/config"
	"github.com/A1exMas1ov/task-3/internal/jsonwriter"
	"github.com/A1exMas1ov/task-3/internal/models"
	"github.com/A1exMas1ov/task-3/internal/xmlparser"
)

const dirPermission = 0o755

func main() {
	configPath := flag.String("config", "config/config.yaml", "path to config file")
	flag.Parse()

	config, err := config.ReadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	var valCurs models.ValCurs

	err = xmlparser.ParseXML(config.InputFile, &valCurs)
	if err != nil {
		panic(err)
	}

	valCurs.SortByValue()

	err = jsonwriter.SaveJSON(config.OutputFile, valCurs.Valutes, dirPermission)
	if err != nil {
		panic(err)
	}
}
