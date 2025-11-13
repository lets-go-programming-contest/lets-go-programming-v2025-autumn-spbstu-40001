package main

import (
	"flag"
	"fmt"

	"github.com/DimasFantomasA/task-3/internal/config"
	"github.com/DimasFantomasA/task-3/internal/currency"
)

func main() {
	path := flag.String("config", "config.yaml", "path to yaml config file (default: config.yaml)")
	flag.Parse()

	config, err := config.LoadConfig(*path)
	if err != nil {
		panic(fmt.Errorf("load config: %w", err))
	}

	err = currency.Process(config.InputFile, config.OutputFile)
	if err != nil {
		panic(fmt.Errorf("process currency: %w", err))
	}

	fmt.Println("Processing completed successfully")
}
