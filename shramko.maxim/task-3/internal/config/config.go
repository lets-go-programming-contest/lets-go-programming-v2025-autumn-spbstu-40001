package config

import (
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
		return nil, err
	}

	var settings AppSettings
	if err = yaml.Unmarshal(content, &settings); err != nil {
		return nil, err
	}

	return &settings, nil
}
