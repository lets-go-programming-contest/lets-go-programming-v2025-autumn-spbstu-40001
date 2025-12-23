package main

import (
	"fmt"

	"github.com/Rychmick/task-8/internal/config"
)

func main() {
	cfg, err := config.GetActive()
	if err != nil {
		fmt.Println("failed to load config: %w")

		return
	}

	fmt.Print(cfg.Env, " ", cfg.LogLvl)
}
