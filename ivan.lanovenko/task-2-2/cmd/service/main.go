package main

import (
	"container/heap"
	"fmt"

	"github.com/Tuc0Sa1amanka/task-2-2/internal/maxheap"
)

func main() {
	var numberOfDishes int

	if _, err := fmt.Scan(&numberOfDishes); err != nil {
		fmt.Println("failed to read number of dishes: ", err)

		return
	}

	heapOfRatings := &maxheap.IntHeap{}
	heap.Init(heapOfRatings)

	for range numberOfDishes {
		var current int

		if _, err := fmt.Scan(&current); err != nil {
			fmt.Println("failed to read current rating: ", err)

			return
		}

		heap.Push(heapOfRatings, current)
	}

	var sequenceNumber int

	_, err := fmt.Scan(&sequenceNumber)
	if err != nil {
		fmt.Println("failed to read sequenceNumber: ", err)

		return
	}

	if sequenceNumber > numberOfDishes {
		fmt.Println("the priority sequence number should not be more than the number of dishes")

		return
	}

	if heapOfRatings.Len() == 0 {
		fmt.Println("there is no dishes")

		return
	}

	for range sequenceNumber - 1 {
		heap.Pop(heapOfRatings)
	}

	fmt.Println(heap.Pop(heapOfRatings))
}
