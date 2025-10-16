package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

var configPathFlag = flag.String("config", "config.yaml", "path to config file")

func main() {
	configFile, err := os.Open(*configPathFlag)
	if err != nil && os.IsNotExist(err) {
		panic("config file does not exist")
	}

	configData, err := io.ReadAll(configFile)
	if err != nil {
		panic("cannot read file")
	}

	var files Config

	err = yaml.Unmarshal([]byte(configData), &files)

	if err != nil {
		panic("cannot unmarshal data")
	}

	fmt.Println(files.InputFile)
	fmt.Println(files.OutputFile)
}
