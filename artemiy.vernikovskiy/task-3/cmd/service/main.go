package main

import (
	"flag"
	"fmt"

	"github.com/Aapng-cmd/task-3/internal/files"
	"github.com/Aapng-cmd/task-3/internal/models"
	"github.com/Aapng-cmd/task-3/internal/sorts"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "", "Path to config file")

	flag.Parse()

	if configPath == "" {
		fmt.Println("Config file needed. Pass it with --config filename")
		panic("AAAAAAAAAAAAAAA")
	}

	inputFile, outputFile, err := files.ReadYAMLConfigFile(configPath)
	if err != nil {
		panic(err)
	}

	var valCurs models.ValCurs

	valCurs, err = files.ReadAndParseXML(inputFile)
	if err != nil {
		panic(err)
	}

	valCurs = sorts.SortDataByValue(valCurs)

	err = files.WriteDataToJSON(valCurs, outputFile)
	if err != nil {
		panic(err)
	}
}
