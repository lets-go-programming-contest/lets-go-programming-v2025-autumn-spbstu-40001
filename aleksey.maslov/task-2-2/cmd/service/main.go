package main

import (
	"container/heap"
	"fmt"

	"github.com/A1exMas1ov/task-2-2/internal/intheap"
)

func main() {
	var dishesCount int

	_, err := fmt.Scan(&dishesCount)
	if err != nil {
		fmt.Println("Invalid input", err)

		return
	}

	ratings := &intheap.IntHeap{}
	heap.Init(ratings)

	for range dishesCount {
		var rating int

		_, err = fmt.Scan(&rating)
		if err != nil {
			fmt.Println("Invalid input", err)

			return
		}

		heap.Push(ratings, rating)
	}

	var selectedDish int

	_, err = fmt.Scan(&selectedDish)
	if err != nil {
		fmt.Println("Invalid input", err)

		return
	}

	printSelectedDish(*ratings, selectedDish)
}

func printSelectedDish(ratings intheap.IntHeap, selectedDish int) {
	if selectedDish > ratings.Len() {
		fmt.Println("There is no such dish")

		return
	}

	for range ratings.Len() - selectedDish {
		heap.Pop(&ratings)
	}

	fmt.Println(heap.Pop(&ratings))
}
