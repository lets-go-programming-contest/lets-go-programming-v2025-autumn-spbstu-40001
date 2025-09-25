package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (slice *IntHeap) Len() int {
	return len(*slice)
}

func (slice *IntHeap) Less(i, j int) bool {
	return (*slice)[i] > (*slice)[j]
}

func (slice *IntHeap) Swap(i, j int) {
	(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
}

func (slice *IntHeap) Push(value any) {
	intValue, ok := value.(int)
	if !ok {
		return
	}

	*slice = append(*slice, intValue)
}

func (slice *IntHeap) Pop() any {
	result := (*slice)[len(*slice)-1]
	*slice = (*slice)[0 : len(*slice)-1]

	return result
}

func (slice *IntHeap) Top() int {
	return (*slice)[0]
}

func main() {
	var nDishes int

	_, err := fmt.Scan(&nDishes)
	if err != nil {
		fmt.Println("Failed to read dishes count")
		fmt.Println(err)

		return
	}

	dishesQueue := &IntHeap{}

	for range nDishes {
		var dishValue int

		_, err = fmt.Scan(&dishValue)
		if err != nil {
			fmt.Println("Failed to read dish value")
			fmt.Println(err)

			return
		}

		heap.Push(dishesQueue, dishValue)
	}

	var dishId int

	_, err = fmt.Scan(&dishId)
	if err != nil {
		fmt.Println("Failed to read priority number")
		fmt.Println(err)

		return
	}

	if (dishId > dishesQueue.Len()) || (dishId <= 0) {
		fmt.Println("Entered nonexistent priority number")

		return
	}

	for range dishId - 1 {
		heap.Pop(dishesQueue)
	}

	fmt.Println(dishesQueue.Top())
}
