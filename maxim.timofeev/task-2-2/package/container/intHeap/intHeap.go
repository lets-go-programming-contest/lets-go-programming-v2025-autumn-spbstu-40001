package intHeap

import "fmt"

type IntHeap []int

func (h *IntHeap) Len() int           { return len(*h) }
func (h *IntHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *IntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntHeap) Push(x any) {
	if val, check := x.(int); check {
		*h = append(*h, val)
	} else {
		fmt.Println("invalid type in intHeap push")
	}
}

func (h *IntHeap) Pop() any {
	old := *h
	length := len(old)
	top := old[length-1]
	*h = old[:length-1]

	return top
}
