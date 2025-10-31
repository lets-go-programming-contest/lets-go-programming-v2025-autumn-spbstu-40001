package config

import (
	"io"
	"os"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func GetConfigFile(configPath *string) (*os.File, error) {
	configFile, err := os.Open(*configPath)

	if err != nil && os.IsNotExist(err) {
		return nil, os.ErrNotExist
	}

	return configFile, nil
}

func GetConfigData(configFile *os.File) ([]byte, error) {
	configData, err := io.ReadAll(configFile)
	if err != nil {
		return configData, io.EOF
	}

	return configData, nil
}
