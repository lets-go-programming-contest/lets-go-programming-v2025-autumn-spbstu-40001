package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Settings struct {
	InputFilePath  string `yaml:"input-file"`
	OutputFilePath string `yaml:"output-file"`
}

func Parse(configPath string) (Settings, error) {
	var result Settings

	yamlFileData, err := os.ReadFile(configPath)
	if err != nil {
		return result, fmt.Errorf("failed to read config file: %w", err)
	}

	err = yaml.Unmarshal(yamlFileData, &result)
	if err != nil {
		return result, fmt.Errorf("failed to parse config file: %w", err)
	}

	return result, nil
}
