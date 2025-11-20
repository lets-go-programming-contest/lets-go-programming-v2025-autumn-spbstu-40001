package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/vikaglushkova/task-3/internal/config"
	"github.com/vikaglushkova/task-3/internal/currency"
	"github.com/vikaglushkova/task-3/internal/json"
	"github.com/vikaglushkova/task-3/internal/xml"
)

const defaultConfigPath = "config.yaml"

func main() {
	configPath := flag.String("config", defaultConfigPath, "Path to configuration file")
	flag.Parse()

	cfg, err := config.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	valCurs, err := xml.ParseXMLFile(cfg.InputFile)
	if err != nil {
		log.Fatalf("Error reading XML data: %v", err)
	}

	currencies := currency.ConvertAndSort(valCurs)

	err = json.WriteToFile(cfg.OutputFile, currencies, 0o755)
	if err != nil {
		log.Fatalf("Error saving results: %v", err)
	}

	fmt.Printf("Successfully processed %d currencies. Results saved to: %s\n", len(currencies), cfg.OutputFile)
}
