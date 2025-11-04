package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type FilePathConfig struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func ReadConfig(path string) (*FilePathConfig, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		panic("Config file does not exist: " + path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		panic("Cannot read config file: " + err.Error())
	}

	var config FilePathConfig

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic("Invalid YAML format: " + err.Error())
	}

	return &config, nil
}
