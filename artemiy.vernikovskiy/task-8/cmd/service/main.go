package main

import (
	"fmt"

	"github.com/Aapng-cmd/task-8/internal/config"
)

func main() {
	config, err := config.ParseConfig()
	if err != nil {
		fmt.Println("Error: ", err)

		return
	}

	fmt.Print(config.Env, " ", config.LogLevel)
}
