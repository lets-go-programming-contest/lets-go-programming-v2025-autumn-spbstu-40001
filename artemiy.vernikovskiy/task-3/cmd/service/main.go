// Package main is the entry point for the currency data processing application.
// It reads configuration, parses XML data, sorts it, and writes the result to JSON.
package main

import (
	"flag"
	"log"

	"github.com/Aapng-cmd/task-3/internal/files"
	"github.com/Aapng-cmd/task-3/internal/models"
	"github.com/Aapng-cmd/task-3/internal/sorts"
)

// main parses command-line flags, reads configuration, processes currency data, and handles errors.
func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "", "Path to config file")

	flag.Parse()

	if configPath == "" {
		log.Fatal("Config file path is required. Use --config to specify the path.")
	}

	inputFile, outputFile, err := files.ReadYAMLConfigFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var valCurs models.ValCurs

	valCurs, err = files.ReadAndParseXML(inputFile)
	if err != nil {
		log.Fatalf("Failed to read and parse XML file: %v", err)
	}

	valCurs = sorts.SortDataByValue(valCurs)

	err = files.WriteDataToJSON(valCurs, outputFile)
	if err != nil {
		log.Fatalf("Failed to write JSON file: %v", err)
	}
}
