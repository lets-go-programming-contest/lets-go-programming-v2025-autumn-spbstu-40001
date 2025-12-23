package config

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"
)

var errEmptyConfig = errors.New("empty config")

type Config struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

func Load() (*Config, error) {
	if len(rawConfig) == 0 {
		return nil, errEmptyConfig
	}

	cfg := new(Config)
	if err := yaml.Unmarshal(rawConfig, cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML config: %w", err)
	}

	return cfg, nil
}
