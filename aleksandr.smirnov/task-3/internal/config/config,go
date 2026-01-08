package config

import (
	"fmt"
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
	defer file.Close()

	var cfg AppConfig
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("cannot decode yaml: %w", err)
	}

	return &cfg, nil
}
