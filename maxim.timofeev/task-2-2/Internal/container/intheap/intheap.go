package intheap

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

func (h *IntHeap) Push(x any) {
	val, ok := x.(int)
	if !ok {
		panic("IntHeap: invalid type, expected int")
	}

	*h = append(*h, val)
}

func (h *IntHeap) Pop() any {
	old := *h
	length := len(old)

	if length == 0 {
		return nil
	}

	top := old[length-1]
	*h = old[:length-1]
	fmt.Println(old)

	return top
}

func (h *IntHeap) Top() any {
	if len(*h) == 0 {
		panic("top heap is empty")
	}

	return (*h)[0]
}

func (h *IntHeap) CalculatePriority(priority int) any {
	for h.Len() > priority {
		heap.Pop(h)
	}

	return h.Top()
}
