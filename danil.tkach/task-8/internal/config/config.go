package config

import (
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

var current Config

func init() {
	if err := yaml.Unmarshal(ActiveConfig, &current); err != nil {
		log.Fatalf("error parsing yaml: %v", err)
	}
}

func Get() Config {
	return current
}
