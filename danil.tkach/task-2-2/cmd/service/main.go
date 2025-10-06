package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (heap MinHeap) Len() int {
	return len(heap)
}

func (heap MinHeap) Less(index1, index2 int) bool {
	return heap[index1] < heap[index2]
}

func (heap MinHeap) Swap(index1, index2 int) {
	heap[index1], heap[index2] = heap[index2], heap[index1]
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

func main() {
	var dishesCount int

	_, err := fmt.Scanln(&dishesCount)
	if err != nil {
		return
	}

	myHeap := &MinHeap{}
	heap.Init(myHeap)

	var index int

	_, err = fmt.Scanln(&index)
	if err != nil {
		return
	}

}
