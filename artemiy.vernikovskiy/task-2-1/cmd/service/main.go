package main

import (
	"fmt"
	"os"

	"github.com/Aapng-cmd/task-2-1/internal/commander"
	"github.com/Aapng-cmd/task-2-1/internal/models"
)

func main() {
	wrapper := models.Wrapper{Cin: os.Stdin, Cout: os.Stdout}

	err := commander.CalcForDepartment(&wrapper)
	if err != nil {
		fmt.Println(err)
	}
}
