package config

import (
	"errors"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

func Load() (*Config, error) {
	if len(rawConfig) == 0 {
		return nil, errors.New("empty config")
	}

	cfg := new(Config)
	if err := yaml.Unmarshal(rawConfig, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
