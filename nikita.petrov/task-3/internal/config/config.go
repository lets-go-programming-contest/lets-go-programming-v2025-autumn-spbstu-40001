package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func GetConfigInfo(configPath *string) (Config, error) {
	var configInfo Config

	configFile, err := os.Open(*configPath)
	if err != nil {
		return configInfo, fmt.Errorf("can't get config file descriptor: %w", err)
	}

	configData, err := io.ReadAll(configFile)
	if err != nil {
		return configInfo, fmt.Errorf("can't get config data: %w", err)
	}

	err = yaml.Unmarshal(configData, &configInfo)
	if err != nil {
		return configInfo, fmt.Errorf("can't unmarshal config data: %w", err)
	}

	return configInfo, nil
}
