package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Input  string `yaml:"input-file"`
	Output string `yaml:"output-file"`
}

func Load(path string) (*AppConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open config: %w", err)
	}

	cfg, decodeErr := decodeConfig(file)

	if closeErr := file.Close(); decodeErr != nil {
		if closeErr != nil {
			return nil, fmt.Errorf("failed to load config: %w", decodeErr)
		}
		return nil, fmt.Errorf("failed to load config: %w", decodeErr)
	} else if closeErr != nil {
		return nil, fmt.Errorf("failed to close config file: %w", closeErr)
	}

	return cfg, nil
}

func decodeConfig(r io.Reader) (*AppConfig, error) {
	var cfg AppConfig
	decoder := yaml.NewDecoder(r)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("cannot decode yaml: %w", err)
	}
	return &cfg, nil
}
