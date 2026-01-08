package main

import (
	"flag"
	"fmt"

	"github.com/A1exCRE/task-3/internal/config"
	"github.com/A1exCRE/task-3/internal/currency"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Config loaded\n")

	_, err = currency.LoadFromFile(cfg.Input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("XML loaded successfully")
}
