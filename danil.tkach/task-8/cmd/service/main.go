package main

import (
	"fmt"
	"log"

	"github.com/Danil3352/task-8/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s %s\n", cfg.Environment, cfg.LogLevel)
}
