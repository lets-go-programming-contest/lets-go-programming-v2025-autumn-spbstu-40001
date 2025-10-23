package maxheap

type IntHeap []int

func (heap *IntHeap) Len() int {
	return len(*heap)
}

func (heap *IntHeap) checkIndexes(indexes ...int) {
	for _, index := range indexes {
		if index < 0 || index >= len(*heap) {
			panic("index out of range")
		}
	}
}

func (heap *IntHeap) Less(firstIndex, secondIndex int) bool {
	heap.checkIndexes(firstIndex, secondIndex)

	return (*heap)[firstIndex] > (*heap)[secondIndex]
}

func (heap *IntHeap) Swap(firstIndex, secondIndex int) {
	heap.checkIndexes(firstIndex, secondIndex)
	(*heap)[firstIndex], (*heap)[secondIndex] = (*heap)[secondIndex], (*heap)[firstIndex]
}

func (heap *IntHeap) Push(value any) {
	num, ok := value.(int)
	if !ok {
		panic("failed to cast new value to int")
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
