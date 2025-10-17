package myheap

type Heap []int

func (heap *Heap) Len() int {
	return len(*heap)
}

func (heap *Heap) Less(firstIndex, secondIndex int) bool {
	return (*heap)[firstIndex] > (*heap)[secondIndex]
}

func (heap *Heap) Swap(firstIndex, secondIndex int) {
	(*heap)[firstIndex], (*heap)[secondIndex] = (*heap)[secondIndex], (*heap)[firstIndex]
}

func (heap *Heap) Push(inter interface{}) {
	number, ok := inter.(int)
	if !ok {
		return
	}

	*heap = append(*heap, number)
}

func (heap *Heap) Pop() any {
	oldHeap := *heap

	nLengthOfHeap := len(oldHeap)
	if nLengthOfHeap == 0 {
		return nil
	}

	x := oldHeap[nLengthOfHeap-1]
	*heap = oldHeap[0 : nLengthOfHeap-1]

	return x
}
