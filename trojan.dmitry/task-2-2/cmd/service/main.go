package main

import (
	"container/heap"
	"errors"
	"fmt"
)

var ErrOfScan = errors.New("scan failed")

type IntHeap []int

func (iHeap *IntHeap) Len() int {
	return len(*iHeap)
}

func (iHeap *IntHeap) Less(firstIndex, secondIndex int) bool {
	if firstIndex < 0 || secondIndex < 0 || firstIndex >= len(*iHeap) || secondIndex >= len(*iHeap) {
		panic("index out of range in less")
	}

	return (*iHeap)[firstIndex] < (*iHeap)[secondIndex]
}

func (iHeap *IntHeap) Swap(firstIndex, secondIndex int) {
	if firstIndex < 0 || secondIndex < 0 || firstIndex >= len(*iHeap) || secondIndex >= len(*iHeap) {
		panic("index out of range in swap")
	}

	(*iHeap)[firstIndex], (*iHeap)[secondIndex] = (*iHeap)[secondIndex], (*iHeap)[firstIndex]
}

func (iHeap *IntHeap) Push(x any) {
	value, ok := x.(int)
	if !ok {
		panic("heap: Push of non-int value")
	}

	*iHeap = append(*iHeap, value)
}

func (iHeap *IntHeap) Pop() any {
	olhH := *iHeap
	length := len(olhH)

	if length == 0 {
		return nil
	}

	last := olhH[length-1]
	*iHeap = olhH[:length-1]

	return last
}

func removeMinUntil(dishHeap *IntHeap, numOfPreference int) {
	for dishHeap.Len() > numOfPreference {
		heap.Pop(dishHeap)
	}
}

func readInt() (int, error) {
	var val int

	_, err := fmt.Scan(&val)
	if err != nil {
		return 0, ErrOfScan
	}

	return val, nil
}

func main() {
	countOfDishes, err := readInt()
	if err != nil {
		fmt.Println("Invalid input of count of dishes", err)

		return
	}

	dishHeap := &IntHeap{}
	heap.Init(dishHeap)

	for range countOfDishes {
		rating, err := readInt()
		if err != nil {
			fmt.Println("Invalid input of rating of dish", err)

			return
		}

		heap.Push(dishHeap, rating)
	}

	numOfPreference, err := readInt()
	if err != nil {
		fmt.Println("Invalid input of num of preference", err)

		return
	}

	if numOfPreference < 1 || numOfPreference > countOfDishes {
		fmt.Println("Num of preference out of allowed range")

		return
	}

	removeMinUntil(dishHeap, numOfPreference)

	if dishHeap.Len() != numOfPreference || dishHeap.Len() == 0 {
		fmt.Println("Heap size mismatch after trimming")

		return
	}

	val := heap.Pop(dishHeap)
	got, ok := val.(int)

	if !ok {
		fmt.Println("Heap returned non-int value")

		return
	}

	fmt.Println(got)
}
