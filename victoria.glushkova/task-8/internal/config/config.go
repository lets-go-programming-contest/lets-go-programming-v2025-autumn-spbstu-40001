package config

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"
)

var (
	ErrEnvironmentRequired = errors.New("environment field is required")
	ErrLogLevelRequired    = errors.New("log_level field is required")
)

type Config struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

type Loader interface {
	Load() ([]byte, error)
}

func Load(data []byte) (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	if cfg.Environment == "" {
		return nil, ErrEnvironmentRequired
	}

	if cfg.LogLevel == "" {
		return nil, ErrLogLevelRequired
	}

	return &cfg, nil
}

func GetConfig(loader Loader) (*Config, error) {
	data, err := loader.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return Load(data)
}
