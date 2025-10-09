package main

import (
	"container/heap"
	"fmt"

	"github.com/A1exMas1ov/task-2-2/internal/intheap"
)

func main() {
	var dishesCount int

	_, err := fmt.Scanln(&dishesCount)
	if err != nil {
		fmt.Println("Invalid input", err)

		return
	}

	ratings := &intheap.IntHeap{}

	for range dishesCount {
		var rating int

		_, err = fmt.Scanln(&rating)
		if err != nil {
			fmt.Println("Invalid input", err)

			return
		}
		heap.Push(ratings, rating)
	}

	var selectedDish int

	_, err = fmt.Scanln(&selectedDish)
	if err != nil {
		fmt.Println("Invalid input", err)

		return
	}

	/*result := //func
	fmt.println(result)*/
}
