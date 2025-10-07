package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (iHeap *IntHeap) Len() int {
	return len(*iHeap)
}

func (iHeap *IntHeap) Less(firstIndex, secondIndex int) bool {
	return (*iHeap)[firstIndex] < (*iHeap)[secondIndex]
}

func (iHeap *IntHeap) Swap(firstIndex, secondIndex int) {
	(*iHeap)[firstIndex], (*iHeap)[secondIndex] = (*iHeap)[secondIndex], (*iHeap)[firstIndex]
}

func (iHeap *IntHeap) Push(x any) {
	value, ok := x.(int)
	if !ok {
		fmt.Println("Invalid type")

		return
	}
	*iHeap = append(*iHeap, value)
}

func (iHeap *IntHeap) Pop() any {
	olhH := *iHeap
	length := len(olhH)
	last := olhH[length-1]
	*iHeap = olhH[:length-1]

	return last
}

func removeMinUntil(dishHeap *IntHeap, numOfPreference int) {
	for dishHeap.Len() > numOfPreference {
		heap.Pop(dishHeap)
	}
}

func main() {
	var countOfDishes int

	_, err := fmt.Scan(&countOfDishes)
	if err != nil || countOfDishes < 1 || countOfDishes > 10000 {
		fmt.Println("Invalid input")

		return
	}

	dishHeap := &IntHeap{}
	heap.Init(dishHeap)

	for range countOfDishes {
		var rating int
		_, err = fmt.Scan(&rating)
		if err != nil || rating < -10000 || rating > 10000 {
			fmt.Println("Invalid input")

			return
		}
		heap.Push(dishHeap, rating)
	}

	var numOfPreference int

	_, err = fmt.Scan(&numOfPreference)
	if err != nil || numOfPreference < 1 || numOfPreference > countOfDishes {
		fmt.Println("Invalid input")

		return
	}

	removeMinUntil(dishHeap, numOfPreference)

	if dishHeap.Len() == numOfPreference && dishHeap.Len() > 0 {
		fmt.Println((*dishHeap)[0])
	} else {
		fmt.Println("Invalid input")
	}

}
