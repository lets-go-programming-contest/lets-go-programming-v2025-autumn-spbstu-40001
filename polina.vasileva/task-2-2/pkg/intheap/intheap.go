package intheap

import "fmt"

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	value, ok := x.(int)
	if !ok {
		panic(fmt.Sprintf("IntHeap.Push: expected int, got %T", x))
	}

	*h = append(*h, value)
}

func (h *IntHeap) Pop() any {
	old := *h

	lenHeap := len(old)
	if lenHeap == 0 {
		panic("pop from empty heap")
	}

	x := old[lenHeap-1]
	*h = old[0 : lenHeap-1]

	return x
}
