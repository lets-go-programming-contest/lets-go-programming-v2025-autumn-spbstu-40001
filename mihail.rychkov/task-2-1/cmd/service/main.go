package main

import "os"
import "fmt"
import "github.com/Rychmick/task-2-1/internal/mainproc"

func main() {
	var nDepartments int

	_, err := fmt.Scan(&nDepartments)
	if err != nil {
		fmt.Println("Failed to read departments count")
		fmt.Println(err)

		return
	}

	for range nDepartments {
		_, err := mainproc.ProcessDepartmentWishes(os.Stdin, os.Stdout);
		if (err != nil) {
			fmt.Println(err);
		}
	}
}
