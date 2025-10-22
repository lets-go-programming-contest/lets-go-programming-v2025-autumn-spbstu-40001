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

func ParseYaml(configPath string) (*Config, error) {
	result := &Config{InputFilePath: "", OutputFilePath: ""}

	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %w", err)
	}

	err = yaml.Unmarshal(file, &result)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal file: %w", err)
	}

	return result, nil
}
