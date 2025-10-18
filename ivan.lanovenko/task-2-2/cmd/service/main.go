package main

import (
	"container/heap"
	"fmt"

	"github.com/Tuc0Sa1amanka/task-2-2/internal/maxheap"
)

func main() {
	var numberOfDishes int

	if _, err := fmt.Scan(&numberOfDishes); err != nil {
		fmt.Println("Failed to read number of dishes: ", err)

		return
	}

	heapOfRatings := &maxheap.IntHeap{}
	heap.Init(heapOfRatings)

	for range numberOfDishes {
		var current int

		if _, err := fmt.Scan(&current); err != nil {
			fmt.Println("Failed to read current rating: ", err)

			return
		}

		heap.Push(heapOfRatings, current)
	}

	var sequenceNumber int

	_, err := fmt.Scan(&sequenceNumber)
	if err != nil {
		fmt.Println("Failed to read sequenceNumber: ", err)

		return
	}

	if sequenceNumber > numberOfDishes {
		fmt.Println("The priority sequence number should not be more than the number of dishes")

		return
	}

	var value int

	for range sequenceNumber {
		num, ok := heap.Pop(heapOfRatings).(int)
		if !ok {
			return
		}

		value = num
	}

	fmt.Println(value)
}
