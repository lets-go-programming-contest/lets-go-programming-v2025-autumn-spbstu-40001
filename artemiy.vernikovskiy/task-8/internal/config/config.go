package config

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

// will this help? Because i thought i needed only tag-related build
//
//go:embed test_prod.yaml
var configFile []byte

type Config struct {
	Env      string `yaml:"environment"`
	LogLevel string `yaml:"log_level"`
}

func ParseConfig() (*Config, error) {
	var config Config

	err := yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, fmt.Errorf("no config from YAML: %w", err)
	}

	return &config, nil
}
