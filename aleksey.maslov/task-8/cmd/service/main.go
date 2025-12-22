package main

import (
	"fmt"

	"github.com/A1exMas1ov/task-8/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("when loading config:", err)
		return
	}

	fmt.Print(cfg.Environment, " ", cfg.LogLevel)
}
