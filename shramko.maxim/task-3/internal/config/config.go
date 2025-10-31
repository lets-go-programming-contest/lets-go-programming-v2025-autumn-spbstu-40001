package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type AppSettings struct {
	InputPath  string `yaml:"input-file"`
	OutputPath string `yaml:"output-file"`
}

func LoadSettings(filename string) (*AppSettings, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var settings AppSettings
	if err = yaml.Unmarshal(content, &settings); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return &settings, nil
}
