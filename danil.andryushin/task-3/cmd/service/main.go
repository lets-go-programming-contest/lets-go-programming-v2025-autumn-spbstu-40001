package main

import (
	"fmt"
	"os"

	"github.com/atroxxxxxx/task-3/internal/parcing/config"
)

func main() {
	args := os.Args
	if len(args) < 3 || args[1] != "-config" {
		fmt.Println("invalid args")
	}
	data, err := config.New(args[2])
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
