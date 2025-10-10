package myheap

import (
    "errors"
)

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
		return errors.New("failed to push into heap")
	}

	*heap = append(*heap, number)
	return nil
}

func (heap *Heap) Pop() any {
	oldHeap := *heap
	n := len(oldHeap)
	if (n == 0) {
	    return nil
	}
	// совсем забыл об этом
	x := oldHeap[n-1]
	*heap = oldHeap[0 : n-1]

	return x
}
