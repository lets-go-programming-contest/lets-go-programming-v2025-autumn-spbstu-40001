package config

import (
	"errors"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

var ErrConfigFieldsRequired = errors.New("config file must contain both input-file and output-file fields")

func ReadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("cannot open config file: %w", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		_ = file.Close()

		return nil, fmt.Errorf("cannot read config file: %w", err)
	}

	if err := file.Close(); err != nil {
		return nil, fmt.Errorf("cannot close config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("cannot parse YAML config: %w", err)
	}

	if config.InputFile == "" || config.OutputFile == "" {
		return nil, ErrConfigFieldsRequired
	}

	return &config, nil
}
