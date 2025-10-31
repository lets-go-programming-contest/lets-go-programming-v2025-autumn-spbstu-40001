package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var ErrInvalidConfig = errors.New("input-file and output-file must be specified")

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfig(path string) (*Config, error) {
	val, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(val, &config)

	if err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	if config.InputFile == "" || config.OutputFile == "" {
		return nil, ErrInvalidConfig
	}

	return &config, nil
}
