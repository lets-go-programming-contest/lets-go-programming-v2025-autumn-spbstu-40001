package main

import (
	"fmt"

	"github.com/Danil3352/task-8/internal/config"
)

func main() {
	cfg := config.Get()
	fmt.Printf("%s %s\n", cfg.Environment, cfg.LogLevel)
}
