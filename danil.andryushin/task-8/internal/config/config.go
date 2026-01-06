package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

func LoadConfig() (Config, error) {
	var conf Config

	err := yaml.Unmarshal(configData, conf)
	if err != nil {
		return conf, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return conf, nil
}
