package intheap

import "errors"

var ErrEmptyHeap error = errors.New("can't pop - heap is empty")

type IntHeap []int //nolint:recvcheck

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	val, ok := x.(int)

	if !ok {
		panic("can't convert to int")
	}

	*h = append(*h, val)
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	if n == 0 {
		return nil
	}
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}
