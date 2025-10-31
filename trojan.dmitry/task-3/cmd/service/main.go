package main

import (
	"flag"
	"fmt"

	"github.com/DimasFantomasA/task-3/internal/config"
	"github.com/DimasFantomasA/task-3/internal/currency"
)

func main() {
	path := flag.String("config", "", "path to yaml config file")
	flag.Parse()

	if path == nil || *path == "" {
		panic("config flag error")
	}

	config, err := config.LoadConfig(*path)
	if err != nil {
		panic(fmt.Errorf("load config: %w", err))
	}

	err = currency.Process(config.InputFile, config.OutputFile)
	if err != nil {
		panic(fmt.Errorf("process currency: %w", err))
	}

	fmt.Println("End")
}
