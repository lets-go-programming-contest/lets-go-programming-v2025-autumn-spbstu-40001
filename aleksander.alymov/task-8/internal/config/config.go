package config

import "gopkg.in/yaml.v3"

type Config struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

func Load() (*Config, error) {
	var cfg Config
	err := yaml.Unmarshal(ConfigFile, &cfg)
	return &cfg, err
}
