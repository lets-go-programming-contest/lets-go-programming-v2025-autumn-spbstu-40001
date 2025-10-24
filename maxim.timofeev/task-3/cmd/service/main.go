package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/PigoDog/task-3/internal/IOcurrency"
	"github.com/PigoDog/task-3/internal/config"
)

func main() {
	configPath := flag.String("config", "", "Path to YAML config")
	flag.Parse()

	if *configPath == "" {
		panic("flag --config is empty")
	}

	config, err := config.ReadConfig(*configPath)

	if err != nil {
		panic(fmt.Errorf("failed to read config: %v", err))
	}

	valutes, err := IOcurrency.LoadXML(config.InputFile)

	if err != nil {
		panic(fmt.Errorf("failed to read XML: %v", err))
	}

	sort.Slice(valutes, func(i, j int) bool {
		return valutes[i].Value > valutes[j].Value
	})

	if err = IOcurrency.SaveJSON(config.OutputFile, valutes); err != nil {
		panic(fmt.Errorf("failed to read JSON: %v", err))
	}
}
