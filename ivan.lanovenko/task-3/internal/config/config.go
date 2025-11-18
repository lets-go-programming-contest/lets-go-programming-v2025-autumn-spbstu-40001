package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFilePath  string `yaml:"input-file"`
	OutputFilePath string `yaml:"output-file"`
}

func (config *Config) LoadFromFile(path string) error {
	configFile, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read config file: %w", err)
	}

	err = yaml.Unmarshal(configFile, config)
	if err != nil {
		return fmt.Errorf("unmarshal YAML: %w", err)
	}

	return nil
}
