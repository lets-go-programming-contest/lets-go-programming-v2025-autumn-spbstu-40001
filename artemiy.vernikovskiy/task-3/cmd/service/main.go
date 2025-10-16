package main

import (
	"flag"
	"fmt"

	"github.com/Aapng-cmd/task-3/internal/funcs"
	"github.com/Aapng-cmd/task-3/internal/models"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "", "Path to config file")

	flag.Parse()

	if configPath == "" {
		fmt.Println("Config file needed. Pass it with --config filename")
		panic("AAAAAAAAAAAAAAA")

	}

	inputFile, outputFile, err := funcs.ReadYAMLConfigFile(configPath)
	if err != nil {
		// fmt.Println(err)
		panic(err)
	}

	var valCurs models.ValCurs

	valCurs, err = funcs.ReadAndParseXML(inputFile)
	if err != nil {
		// fmt.Println(err)
		panic(err)
	}

	err = funcs.WriteDataToJSON(valCurs, outputFile)
	if err != nil {
		// fmt.Println(err)
		panic(err)
	}
}
