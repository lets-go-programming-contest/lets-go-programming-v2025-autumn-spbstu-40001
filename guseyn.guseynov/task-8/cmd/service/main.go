package main

import (
	"fmt"

	"github.com/GuseynovGuseynGG/task-8/pkg/config"
)

func main() {
	cfg := config.GetConfig()
	fmt.Print(cfg.Environment + " " + cfg.LogLevel)
}
