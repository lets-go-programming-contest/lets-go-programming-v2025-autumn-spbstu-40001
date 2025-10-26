package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func readConfig(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse YAML config: %w", err)
	}

	return &config, nil
}

func parseXML(path string) (//data, error) {}

func saveJSON(path string, //data) error {}



func main() {
	configPath := flag.String("config", "config.yaml", "path to config")
	flag.Parse()

	config, err := readConfig(*configPath)
	if err != nil {
		panic(err)
	}

	data, err = parseXML(config.InputFile)
	if err != nil {
		panic(err)
	}
//sort
	err = saveJSON(config.OutputFile, //data)
	if err != nil {
		panic(err)
	}
}
