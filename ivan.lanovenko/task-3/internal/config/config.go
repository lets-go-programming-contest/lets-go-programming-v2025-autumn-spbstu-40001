package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFilePath  string `yaml:"input-file"`
	OutputFilePath string `yaml:"output-file"`
}

func (config *Config) LoadFromFile(path string) {
	configFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configFile, config)
	if err != nil {
		panic(err)
	}
}
