package main

import (
	"container/heap"
	"fmt"

	"github.com/Aapng-cmd/task-2-2/internal/my_heap"
)

func main() {
	var (
		nDishesNumber, actualDish, wanting int
		workHeap      = &my_heap.Heap{}
	)

	_, err := fmt.Scan(&nDishesNumber)
	if err != nil {
		fmt.Println("Invalid dishes count")

		return
	}

	heap.Init(workHeap)

	for range nDishesNumber {
		_, err = fmt.Scan(&actualDish)
		if err != nil {
			fmt.Println("invalid input")

			return
		}

		heap.Push(workHeap, actualDish)
	}

	_, err = fmt.Scan(&wanting)
	if err != nil {
		fmt.Println("invalid input")

		return
	}

	for range wanting - 1 {
		if workHeap.Len() == 0 {
			fmt.Println("no dish for you")
		}

		heap.Pop(workHeap)
	}

	fmt.Println(heap.Pop(workHeap))
}
