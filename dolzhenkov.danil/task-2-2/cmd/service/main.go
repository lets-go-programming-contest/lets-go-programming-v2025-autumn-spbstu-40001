package main

import (
	"container/heap"
	"errors"
	"fmt"
)

var ErrOfScan = errors.New("scan failed")

type MaxHeap []int

func (iHeap MaxHeap) Len() int {
	return len(iHeap)
}

func (iHeap MaxHeap) Less(firstIndex, secondIndex int) bool {
	if firstIndex < 0 || secondIndex < 0 || firstIndex >= len(iHeap) || secondIndex >= len(iHeap) {
		panic("index out of range in less")
	}

	return iHeap[firstIndex] > iHeap[secondIndex]
}

func (iHeap MaxHeap) Swap(firstIndex, secondIndex int) {
	if firstIndex < 0 || secondIndex < 0 || firstIndex >= len(iHeap) || secondIndex >= len(iHeap) {
		panic("index out of range in swap")
	}

	iHeap[firstIndex], iHeap[secondIndex] = iHeap[secondIndex], iHeap[firstIndex]
}

func (iHeap *MaxHeap) Push(x any) {
	value, ok := x.(int)
	if !ok {
		panic("heap: Push of non-int value")
	}

	*iHeap = append(*iHeap, value)
}

func (iHeap *MaxHeap) Pop() any {
	olhH := *iHeap
	length := len(olhH)

	if length == 0 {
		return nil
	}

	last := olhH[length-1]
	*iHeap = olhH[:length-1]

	return last
}

func main() {
	var countOfDishes int

	_, err := fmt.Scan(&countOfDishes)
	if err != nil {
		fmt.Println("Invalid input of count of dishes:", err)
		return
	}

	if countOfDishes < 1 || countOfDishes > 10000 {
		fmt.Println("Count of dishes must be between 1 and 10000")
		return
	}

	dishHeap := &MaxHeap{}
	heap.Init(dishHeap)

	for range countOfDishes {
		var rating int

		_, err := fmt.Scan(&rating)
		if err != nil {
			fmt.Println("Invalid input of rating of dish:", err)
			return
		}

		if rating < -10000 || rating > 10000 {
			fmt.Println("Rating must be between -10000 and 10000")
			return
		}

		heap.Push(dishHeap, rating)
	}

	var numOfPreference int

	_, err = fmt.Scan(&numOfPreference)
	if err != nil {
		fmt.Println("Invalid input of num of preference:", err)
		return
	}

	if numOfPreference < 1 || numOfPreference > countOfDishes {
		fmt.Println("Num of preference must be between 1 and", countOfDishes)
		return
	}

	for range numOfPreference - 1 {
		heap.Pop(dishHeap)
	}

	kthLargestRaw := heap.Pop(dishHeap)
	kthLargest, ok := kthLargestRaw.(int)
	if !ok {
		fmt.Println("Type assertion failed")
		return
	}

	fmt.Println(kthLargest)
}
