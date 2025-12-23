package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env    string `yaml:"environment"`
	LogLvl string `yaml:"log_level"`
}

func GetActive() (Config, error) {
	var result Config

	err := yaml.Unmarshal(activeConfigRaw, &result)
	if err != nil {
		return result, fmt.Errorf("failed to parse active config: %w", err)
	}

	return result, nil
}
