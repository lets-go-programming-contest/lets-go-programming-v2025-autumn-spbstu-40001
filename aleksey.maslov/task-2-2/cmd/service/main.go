package main

import (
	"container/heap"
	"fmt"

	"github.com/A1exMas1ov/task-2-2/internal/intheap"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Panic recovered:", r)
		}
	}()

	var dishesCount int

	_, err := fmt.Scan(&dishesCount)
	if err != nil {
		fmt.Println("Failed to read dishes count", err)

		return
	}

	ratings := &intheap.IntHeap{}
	heap.Init(ratings)

	for range dishesCount {
		var rating int

		_, err = fmt.Scan(&rating)
		if err != nil {
			fmt.Println("Failed to read rating", err)

			return
		}

		heap.Push(ratings, rating)
	}

	var selectedDish int

	_, err = fmt.Scan(&selectedDish)
	if err != nil {
		fmt.Println("Failed to read selected dish", err)

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

	result := heap.Pop(&ratings)
	if result == nil {
		fmt.Println("There is no such dish")

		return
	}

	fmt.Println(result)
}
