package intheap

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	if i >= h.Len() || j >= h.Len() {
		panic("index out of range in intheap")
	}

	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	if i >= h.Len() || j >= h.Len() {
		panic("index out of range in intheap")
	}

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	value, ok := x.(int)
	if ok {
		*h = append(*h, value)
	} else {
		panic("Value is not an int")
	}
}

func (h *IntHeap) Pop() any {
	if h.Len() == 0 {
		return nil
	}

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}
