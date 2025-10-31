package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfig(configPath string) (*Config, error) {
	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("failed parse YAML: %w", err)
	}

	if config.InputFile == "" {
		return nil, fmt.Errorf("input-file must be filled in")
	}
	if config.OutputFile == "" {
		return nil, fmt.Errorf("output-file must be filled in")
	}

	return &config, nil
}
