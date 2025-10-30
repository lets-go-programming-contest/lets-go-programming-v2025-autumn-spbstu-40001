package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func New(path string) (Config, error) {
	var config Config
	data, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("cannot read config file: %w", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("unarshallig failed: %w", err)
	}
	return config, nil
}
