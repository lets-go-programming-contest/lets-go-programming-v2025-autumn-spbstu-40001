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

// Global constants for file permissions will be here and will not be moved anywhere else.
const (
	// DirPerm defines the permissions for created directories.
	DirPerm = 0o750 // rwxr-x---
	// FilePerm defines the permissions for created files.
	FilePerm = 0o600 // rw-------
)

// main parses command-line flags, reads configuration, processes currency data, and handles errors.
func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "config.yaml", "Path to config file")

	flag.Parse()

	if configPath == "" {
		log.Fatal("Config file path is required. Use --config to specify the path.")
	}

	var settings models.Settings

	settings, err := files.ReadYAMLConfigFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var valCurs models.ValCurs

	valCurs, err = files.ReadAndParseXML(settings.InputFileSetting)
	if err != nil {
		log.Fatalf("Failed to read and parse XML file: %v", err)
	}

	valCurs = sorts.SortDataByValue(valCurs)

	err = files.WriteDataToJSON(valCurs, settings.OutputFileSetting, DirPerm, FilePerm)
	if err != nil {
		log.Fatalf("Failed to write JSON file: %v", err)
	}
}
