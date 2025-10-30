package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/vikaglushkova/task-3/internal/currency"
	"github.com/vikaglushkova/task-3/internal/json"
	"github.com/vikaglushkova/task-3/internal/xml"
	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

var ErrConfigFieldsRequired = errors.New("config file must contain both input-file and output-file fields")

func ReadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("cannot open config file: %w", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		_ = file.Close()

		return nil, fmt.Errorf("cannot read config file: %w", err)
	}

	if err := file.Close(); err != nil {
		return nil, fmt.Errorf("cannot close config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("cannot parse YAML config: %w", err)
	}

	if config.InputFile == "" || config.OutputFile == "" {
		return nil, ErrConfigFieldsRequired
	}

	return &config, nil
}

func main() {
	configPath := flag.String("config", "", "Path to configuration file")
	flag.Parse()

	if *configPath == "" {
		panic("Config file path is required")
	}

	cfg, err := ReadConfig(*configPath)
	if err != nil {
		panic(fmt.Sprintf("Error reading config: %v", err))
	}

	valCurs, err := xml.ParseXMLFile(cfg.InputFile)
	if err != nil {
		panic(fmt.Sprintf("Error reading XML data: %v", err))
	}

	currencies := currency.ConvertAndSort(valCurs)

	err = json.WriteToFile(cfg.OutputFile, currencies)
	if err != nil {
		panic(fmt.Sprintf("Error saving results: %v", err))
	}

	fmt.Printf("Successfully processed %d currencies. Results saved to: %s\n", len(currencies), cfg.OutputFile)
}
