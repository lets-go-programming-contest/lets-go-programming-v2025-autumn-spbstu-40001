package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"

	"github.com/vikaglushkova/task-3/internal/config"
	"github.com/vikaglushkova/task-3/internal/currency"
	"github.com/vikaglushkova/task-3/internal/json"
	"github.com/vikaglushkova/task-3/internal/xmlparser"
)

const (
	defaultConfigPath = "config.yaml"
	dirPermissions    = 0o755
)

type ValCurs struct {
	XMLName xml.Name            `xml:"ValCurs"`
	Valutes []currency.Currency `xml:"Valute"`
}

func main() {
	configPath := flag.String("config", defaultConfigPath, "Path to configuration file")
	flag.Parse()

	cfg, err := config.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	valCurs, err := xmlparser.ParseXMLFile[ValCurs](cfg.InputFile)
	if err != nil {
		log.Fatalf("Error reading XML data: %v", err)
	}

	currencies := currency.ConvertAndSort(valCurs.Valutes)

	err = json.WriteToFile(cfg.OutputFile, currencies, dirPermissions)
	if err != nil {
		log.Fatalf("Error saving results: %v", err)
	}

	fmt.Printf("Successfully processed %d currencies. Results saved to: %s\n", len(currencies), cfg.OutputFile)
}
