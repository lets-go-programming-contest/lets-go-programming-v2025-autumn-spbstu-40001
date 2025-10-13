package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (a *MaxHeap) Len() int {

	return len(*a)
}

func (a MaxHeap) Less(i, j int) bool {

	return a[i] > a[j]
}

func (a *MaxHeap) Swap(i, j int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}

func (a *MaxHeap) Push(x any) {
	*a = append(*a, x.(int))
}

func (a *MaxHeap) Pop() any {
	old := *a
	n := len(old)
	x := old[n-1]
	*a = old[0 : n-1]

	return x
}

func main() {
	var numberDishes int

	_, err := fmt.Scan(&numberDishes)
	if err != nil {
		fmt.Println("Invalid input:", err)

		return
	}

	preferences := &MaxHeap{}
	heap.Init(preferences)

	for range numberDishes {
		var preference int

		_, err = fmt.Scan(&preference)
		if err != nil {
			fmt.Println("Invalid input:", err)

			return
		}

		heap.Push(preferences, preference)
	}

	var dishCount int

	_, err = fmt.Scan(&dishCount)
	if err != nil {
		fmt.Println("Invalid input:", err)

		return
	}

	if dishCount > preferences.Len() {
		fmt.Println("Fewer dishes than preference number")

		return
	}

	for range dishCount - 1 {
		heap.Pop(preferences)
	}

	fmt.Println(heap.Pop(preferences))
}
