package main

import (
	"flag"
	"os"

	"task-3/internal/config"
	"task-3/internal/exporter"
	"task-3/internal/model"
	"task-3/internal/xmlparser"
)

func main() {
	defaultConfig := "config.yaml"
	if envConfig := os.Getenv("APP_CONFIG"); envConfig != "" {
		defaultConfig = envConfig
	}

	configFile := flag.String("config", defaultConfig, "YAML configuration file path")
	flag.Parse()

	settings, err := config.LoadSettings(*configFile)
	if err != nil {
		panic(err)
	}

	currencyData, err := xmlparser.ParseXML(settings.InputPath)
	if err != nil {
		panic(err)
	}

	currencyData.OrderByValue()

	if err = exporter.ExportJSON(settings.OutputPath, currencyData.Items, 0o755); err != nil {
		panic(err)
	}
}
