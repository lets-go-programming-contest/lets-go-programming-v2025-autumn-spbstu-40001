package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (heap *MinHeap) Len() int {
	return len(*heap)
}

func (heap *MinHeap) Less(index1, index2 int) bool {
	return (*heap)[index1] < (*heap)[index2]
}

func (heap *MinHeap) Swap(index1, index2 int) {
	(*heap)[index1], (*heap)[index2] = (*heap)[index2], (*heap)[index1]
}

func (heap *MinHeap) Push(elem any) {
	val, ok := elem.(int)
	if !ok {
		return
	}

	*heap = append(*heap, val)
}

func (heap *MinHeap) Pop() any {
	old := *heap
	size := len(old)
	lastVal := old[size-1]
	*heap = old[0 : size-1]

	return lastVal
}

func RemoveMinElements(myHeap *MinHeap, countElemInHeap int) {
	for myHeap.Len() > countElemInHeap {
		heap.Pop(myHeap)
	}
}

func main() {
	var dishesCount int

	_, err := fmt.Scan(&dishesCount)
	if err != nil || dishesCount < 1 || dishesCount > 10000 {
		return
	}

	myHeap := &MinHeap{}
	heap.Init(myHeap)

	for range dishesCount {
		var dishRating int

		_, err = fmt.Scan(&dishRating)
		if err != nil || dishRating < -10000 || dishRating > 10000 {
			return
		}

		heap.Push(myHeap, dishRating)
	}

	var index int

	_, err = fmt.Scan(&index)
	if err != nil || index > dishesCount || index < 1 {
		return
	}

	RemoveMinElements(myHeap, index)

	if myHeap.Len() != index {
		return
	}

	needDish := heap.Pop(myHeap)
	fmt.Println(needDish)
}
