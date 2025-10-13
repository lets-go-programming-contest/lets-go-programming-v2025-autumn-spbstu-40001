package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (iHeap *IntHeap) Len() int {
	return len(*iHeap)
}

func (iHeap *IntHeap) Less(firstIndex, secondIndex int) bool {
	return (*iHeap)[firstIndex] < (*iHeap)[secondIndex]
}

func (iHeap *IntHeap) Swap(firstIndex, secondIndex int) {
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

func validateCountOfDishes(count int) bool {
	if count < 1 || count > 10000 {
		fmt.Println("Count of dishes out of allowed range")
		return false
	}
	return true
}

	_, err := fmt.Scan(&countOfDishes)
	if err != nil {
		fmt.Println("Invalid input of count of dishes")

		return
	}
	if !validateCountOfDishes {

		return
	}

	dishHeap := &IntHeap{}
	heap.Init(dishHeap)

	for range countOfDishes {
		var rating int

		_, err = fmt.Scan(&rating)
		if err != nil {
			fmt.Println("Invalid input of rating of dish")

			return
		}

		heap.Push(dishHeap, rating)
	}

	var numOfPreference int

	_, err = fmt.Scan(&numOfPreference)
	if err != nil {
		fmt.Println("Invalid input of num of preference")

		return
	}

	if numOfPreference < 1 || numOfPreference > countOfDishes {
		fmt.Println("Num of preference out of allowed range")

		return
	}

	removeMinUntil(dishHeap, numOfPreference)

	if dishHeap.Len() == numOfPreference && dishHeap.Len() > 0 {
		val := heap.Pop(dishHeap)
		if val == nil {
			fmt.Println("Unexpected nil from heap.Pop")

			return
		}
		got, ok := val.(int)

		if !ok {
			fmt.Println("Heap returned non-int value")

			return
		}
		
		fmt.Println(got)
	} else {
		fmt.Println("Heap size mismatch after trimming")
	}
}
