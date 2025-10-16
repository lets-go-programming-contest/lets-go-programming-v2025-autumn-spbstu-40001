package main

import (
	"container/heap"
	"fmt"

	"github.com/GuseynovGuseynGG/task-2-2/internal/intheap"
)

func main() {
	var dishesCount int
	if _, err := fmt.Scan(&dishesCount); err != nil {
		fmt.Println("Invalid input")
		return
	}

	ratings := &intheap.IntHeap{}
	heap.Init(ratings)

	for i := 0; i < dishesCount; i++ {
		var rating int
		if _, err := fmt.Scan(&rating); err != nil {
			fmt.Println("Invalid input")
			return
		}
		heap.Push(ratings, rating)
	}

	var selectedDish int
	if _, err := fmt.Scan(&selectedDish); err != nil {
		fmt.Println("Invalid input")
		return
	}

	if selectedDish > dishesCount || selectedDish <= 0 {
		fmt.Println("There is no such dish")
		return
	}

	for i := 0; i < selectedDish-1; i++ {
		heap.Pop(ratings)
	}

	fmt.Println(heap.Pop(ratings))
}
