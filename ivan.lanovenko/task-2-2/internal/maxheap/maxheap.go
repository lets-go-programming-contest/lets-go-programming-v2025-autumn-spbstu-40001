package maxheap

type IntHeap []int

func (heap *IntHeap) Len() int {
	return len(*heap)
}

func (heap *IntHeap) Less(firstIndex, secondIndex int) bool {
	return (*heap)[firstIndex] > (*heap)[secondIndex]
}

func (heap *IntHeap) Swap(firstIndex, secondIndex int) {
	(*heap)[firstIndex], (*heap)[secondIndex] = (*heap)[secondIndex], (*heap)[firstIndex]
}

func (heap *IntHeap) Push(value any) {
	num, ok := value.(int)
	if !ok {
		panic("IntHeap.Push: expected int")
	}

	*heap = append(*heap, num)
}

func (heap *IntHeap) Pop() any {
	length := len(*heap)
	if length == 0 {
		return nil
	}

	old := *heap
	lastElement := old[length-1]
	*heap = old[0 : length-1]

	return lastElement
}
