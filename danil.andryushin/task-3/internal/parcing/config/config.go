package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

var ErrInvalidConfig = errors.New("invalid config format (expected .yaml)")

func New(path string) (Config, error) {
	var config = Config{}
	if len(path) < 5 || path[len(path)-5:] != ".yaml" {
		return config, ErrInvalidConfig
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}
