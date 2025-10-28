package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	if i < 0 || i >= len(*h) || j < 0 || j >= len(*h) {
		return false
	}

	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	if i < 0 || i >= len(*h) || j < 0 || j >= len(*h) {
		return
	}

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	v, ok := x.(int)
	if !ok {
		panic("IntHeap: Push received non-int value")
	}

	*h = append(*h, v)
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	length := len(old)

	if length == 0 {
		return nil
	}

	x := old[length-1]
	*h = old[:length-1]

	return x
}

func main() {
	var dishCount int
	if _, err := fmt.Scan(&dishCount); err != nil {
		fmt.Println("Error reading number of dishes:", err)

		return
	}

	arr := make([]int, dishCount)
	for i := range arr {
		if _, err := fmt.Scan(&arr[i]); err != nil {
			fmt.Println("Error reading dish rating:", err)

			return
		}
	}

	var kCount int
	if _, err := fmt.Scan(&kCount); err != nil {
		fmt.Println("Error reading k:", err)

		return
	}

	if kCount <= 0 || kCount > dishCount {
		fmt.Println("Invalid k value")

		return
	}

	heapInt := &IntHeap{}
	heap.Init(heapInt)

	for i := 0; i < kCount; i++ {
		heap.Push(heapInt, arr[i])
	}

	for i := kCount; i < len(arr); i++ {
		if arr[i] > (*heapInt)[0] {
			heap.Pop(heapInt)
			heap.Push(heapInt, arr[i])
		}
	}

	popped := heap.Pop(heapInt)
	result, ok := popped.(int)

	if !ok {
		fmt.Println("Invalid result type")

		return
	}

	fmt.Println(result)
}
