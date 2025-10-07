package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (iHeap IntHeap) Len() int {
	return len(iHeap)
}

func (iHeap IntHeap) Less(firstIndex, secondIndex int) bool {
	return iHeap[firstIndex] < iHeap[secondIndex]
}

func (iHeap IntHeap) Swap(firstIndex, secondIndex int) {
	iHeap[firstIndex], iHeap[secondIndex] = iHeap[secondIndex], iHeap[firstIndex]
}

func (iHeap *IntHeap) Push(x any) {
	value, ok := x.(int)
	if !ok {

		return
	}
	*iHeap = append(*iHeap, value)
}

func (iHeap *IntHeap) Pop() any {
	olhH := *iHeap
	length := len(olhH)
	first := olhH[length-1]
	*iHeap = olhH[:length-1]

	return first
}

func main() {
	var countOfDishes int

	_, err := fmt.Scan(&countOfDishes)
	if err != nil || 10000 < countOfDishes || countOfDishes < 1 {
		fmt.Println("Invalid input")
	}

	h := &IntHeap{}
	heap.Init(h)

	for range countOfDishes {
		var rating int
		_, err = fmt.Scan(&rating)
		if err != nil || rating < -10000 || rating > 10000 {
			fmt.Println("Invalid input")
			return
		}
		heap.Push(h, rating)
	}

	var numOfPreference int

	_, err = fmt.Scan(&numOfPreference)
	if err != nil || numOfPreference < 1 || numOfPreference > countOfDishes {
		fmt.Println("Invalid input")
		return
	}
	for h.Len() > numOfPreference {
		heap.Pop(h)
	}
	if h.Len() == numOfPreference && h.Len() > 0 {
		fmt.Println((*h)[0])
	} else {
		fmt.Println("Invalid input")
	}

}
