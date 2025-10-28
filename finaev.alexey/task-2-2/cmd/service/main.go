package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (a *MaxHeap) Len() int {
	return len(*a)
}

func (a *MaxHeap) Less(i, j int) bool {
	if i < 0 || j < 0 || i >= a.Len() || j >= a.Len() {
		panic("index out of range")
	}

	return (*a)[i] > (*a)[j]
}

func (a *MaxHeap) Swap(i, j int) {
	if i < 0 || j < 0 || i >= a.Len() || j >= a.Len() {
		panic("index out of range")
	}

	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}

func (a *MaxHeap) Push(x any) {
	num, noErr := x.(int)
	if noErr {
		*a = append(*a, num)
	} else {
		panic("not correct num")
	}
}

func (a *MaxHeap) Pop() any {
	if a.Len() == 0 {
		return nil
	}

	old := *a
	n := len(old)
	x := old[n-1]
	*a = old[0 : n-1]

	return x
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic recovered:", err)
		}
	}()

	var numberDishes int

	_, err := fmt.Scan(&numberDishes)
	if err != nil {
		fmt.Println("Invalid number of dish input :", err)

		return
	}

	preferences := &MaxHeap{}
	heap.Init(preferences)

	for range numberDishes {
		var preference int

		_, err = fmt.Scan(&preference)
		if err != nil {
			fmt.Println("Invalid preference array input:", err)

			return
		}

		heap.Push(preferences, preference)
	}

	var dishCount int

	_, err = fmt.Scan(&dishCount)
	if err != nil {
		fmt.Println("Invalid preference dish number input:", err)

		return
	}

	if dishCount > preferences.Len() {
		fmt.Println("Fewer dishes than preference number")

		return
	}

	for range dishCount - 1 {
		heap.Pop(preferences)
	}

	result := heap.Pop(preferences)
	if result == nil {
		fmt.Println("Preference dish is missing")

		return
	}

	fmt.Println(result)
}
