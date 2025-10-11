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
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	v, ok := x.(int)
	if ok {
		*h = append(*h, v)
	else {
		panic("IntHeap: Push received non-int value")
	}
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

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

	var topCount int
	if _, err := fmt.Scan(&topCount); err != nil {
		fmt.Println("Error reading k:", err)

		return
	}

	heapInt := &IntHeap{}
	heap.Init(heapInt)

	for _, num := range arr {
		if heapInt.Len() < topCount {
			heap.Push(heapInt, num)
		} else if num > (*heapInt)[0] {
			heap.Pop(heapInt)
			heap.Push(heapInt, num)
		}
	}

	fmt.Println((*heapInt)[0])
}
