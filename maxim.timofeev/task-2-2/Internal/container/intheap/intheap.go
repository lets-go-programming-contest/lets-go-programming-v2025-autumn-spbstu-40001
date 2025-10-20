package intheap

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	if i < 0 && i >= h.Len() || j < 0 && j >= h.Len() {
		panic("index out of range(Less)")
	}

	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	if i < 0 && i >= h.Len() || j < 0 && j >= h.Len() {
		panic("index out of range(Swap)")
	}

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

	return top
}
