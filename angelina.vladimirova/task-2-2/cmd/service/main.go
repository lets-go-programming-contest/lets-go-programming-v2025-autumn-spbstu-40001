package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	value, ok := x.(int)
	if !ok {
		panic("value is not an int")
	}

	*h = append(*h, value)
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	var (
		count     int
		positionK int
	)

	_, err := fmt.Scan(&count)
	if err != nil {
		fmt.Println("Invalid input")

		return
	}

	ratings := make([]int, count)
	for index := range ratings {
		_, err := fmt.Scan(&ratings[index])
		if err != nil {
			fmt.Println("Invalid input")

			return
		}
	}

	_, err = fmt.Scan(&positionK)
	if err != nil {
		fmt.Println("Invalid input")

		return
	}

	if positionK < 1 || positionK > count {
		fmt.Println("There is no such dish")

		return
	}

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	for _, rating := range ratings {
		heap.Push(maxHeap, rating)
	}

	var result int
	for range positionK {
		value := heap.Pop(maxHeap)
		intValue, ok := value.(int)
		if !ok {
			panic("unexpected type from heap")
		}
		result = intValue
	}

	fmt.Println(result)
}
