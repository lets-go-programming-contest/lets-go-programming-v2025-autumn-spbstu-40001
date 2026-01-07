package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Env string `yaml:"environment"`
	Log string `yaml:"log_level"`
}

func ParseConfig(data []byte) (*AppConfig, error) {
	var cfg AppConfig

	err := yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("config parse failed: %w", err)
	}

	return &cfg, nil
}

func (c *AppConfig) Display() {
	fmt.Printf("%s %s\n", c.Env, c.Log)
}
