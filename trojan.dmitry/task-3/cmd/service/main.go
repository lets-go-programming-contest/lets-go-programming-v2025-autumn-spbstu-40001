package main

import (
	"flag"
	"fmt"

	"github.com/DimasFantomasA/internal/config"
	"github.com/DimasFantomasA/internal/currency"
	"github.com/DimasFantomasA/task-3/internal/currency"
	"github.com/go-delve/delve/pkg/config"
)

func main() {
	path := flag.String("config", "", "path to yaml config file")
	flag.Parse()

	if path == nil || *path == "" {
		panic("config flag error")
	}

	config, err := config.LoadConfig(*path)
	if err != nil {
		panic(fmt.Errorf("Error:", err))
	}

	err = currency.Process(config.InputFile, config.OutputFile)
	if err != nil {
		panic(fmt.Errorf("Error:", err))
	}

	fmt.Println("End")
}
