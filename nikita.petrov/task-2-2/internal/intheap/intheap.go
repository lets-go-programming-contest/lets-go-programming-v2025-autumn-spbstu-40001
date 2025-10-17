package intheap

import "errors"

var ErrEmptyHeap error = errors.New("can't pop - heap is empty")

type IntHeap []int //nolint:recvcheck

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h IntHeap) Less(i, j int) bool {
	if i < 0 && i >= h.Len() || j < 0 && j >= h.Len() {
		return false
	}

	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	if i < 0 && i >= h.Len() || j < 0 && j >= h.Len() {
		return
	}

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
	if len(*h) == 0 {
		return -1
	}

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}
