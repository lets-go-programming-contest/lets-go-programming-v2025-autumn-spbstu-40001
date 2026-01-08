package main

import (
	"container/heap"
	"errors"
	"fmt"
)

var (
	errIndexNotPos     = errors.New("index must be positive")
	errIndexOutOfRange = errors.New("index out of range")
)

type MinHeap []int

func (heap *MinHeap) checkIndexBounds(index int) {
	length := heap.Len()

	if index < 0 || index >= length {
		panic(fmt.Sprintf("heap index out of range [%d] with length %d", index, length))
	}
}

func (heap *MinHeap) Len() int {
	return len(*heap)
}

func (heap *MinHeap) Less(index1, index2 int) bool {
	heap.checkIndexBounds(index1)
	heap.checkIndexBounds(index2)

	return (*heap)[index1] < (*heap)[index2]
}

func (heap *MinHeap) Swap(index1, index2 int) {
	heap.checkIndexBounds(index1)
	heap.checkIndexBounds(index2)
	(*heap)[index1], (*heap)[index2] = (*heap)[index2], (*heap)[index1]
}

func (heap *MinHeap) Push(elem any) {
	val, ok := elem.(int)
	if !ok {
		panic(fmt.Sprintf("Push: expected int, but got %T", elem))
	}

	*heap = append(*heap, val)
}

func (heap *MinHeap) Pop() any {
	old := *heap

	size := len(old)
	if size == 0 {
		panic("Pop called on empty MinHeap")
	}

	lastVal := old[size-1]
	*heap = old[0 : size-1]

	return lastVal
}

func readAndValidateIndex(dishesCount int) (int, error) {
	var index int

	_, err := fmt.Scan(&index)
	if err != nil {
		return 0, fmt.Errorf("failed to read index of needed dish: %w", err)
	}

	if index <= 0 {
		return 0, fmt.Errorf("%w: got %d", errIndexNotPos, index)
	}

	if index > dishesCount {
		return 0, fmt.Errorf("%w: index %d, dish count %d", errIndexOutOfRange, index, dishesCount)
	}

	return index, nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error: %v\n", r)
		}
	}()

	var dishesCount int

	_, err := fmt.Scan(&dishesCount)
	if err != nil {
		fmt.Println("Failed to read count of dishes:", err)

		return
	}

	myHeap := &MinHeap{}
	heap.Init(myHeap)

	for range dishesCount {
		var dishRating int

		_, err = fmt.Scan(&dishRating)
		if err != nil {
			fmt.Println("Failed to read rating of dishes:", err)

			return
		}

		heap.Push(myHeap, dishRating)
	}

	index, err := readAndValidateIndex(dishesCount)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	for myHeap.Len() > index {
		heap.Pop(myHeap)
	}

	needDish, ok := heap.Pop(myHeap).(int)

	if !ok {
		fmt.Println("Cannot convert to int")

		return
	}

	fmt.Println(needDish)
}
