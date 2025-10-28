package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml: "input-file"`
	OutputFile string `yaml : "output-file"`
}

func LoadConfig(path string) (*Config, error) {
	val, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(val, &config)
	if err != nil {
		return nil, err
	}

	if config.InputFile == "" || config.OutputFile == "" {
		return nil, err
	}

	return &config, nil
}
