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
	oldH := *iHeap
	n := len(oldH)

	if n == 0 {
		return nil
	}

	last := oldH[n-1]
	*iHeap = oldH[:n-1]

	return last
}

func removeMinUntil(dh *IntHeap, k int) {
	for dh.Len() > k {
		heap.Pop(dh)
	}
}

func validateCountOfDishes(count int) bool {
	if count < 1 || count > 10000 {
		fmt.Println("Count of dishes out of allowed range")

		return false
	}

	return true
}

func readCountOfDishes() (int, error) {
	var count int
	_, err := fmt.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func readRatingsToHeap(count int) (*IntHeap, error) {
	h := &IntHeap{}
	heap.Init(h)

	for i := 0; i < count; i++ {
		var rating int
		_, err := fmt.Scan(&rating)
		if err != nil {
			return nil, err
		}
		heap.Push(h, rating)
	}

	return h, nil
}

func readNumOfPreference(max int) (int, error) {
	var k int
	_, err := fmt.Scan(&k)
	if err != nil {
		return 0, err
	}
	if k < 1 || k > max {
		return 0, fmt.Errorf("num of preference out of allowed range")
	}
	return k, nil
}

func main() {
	countOfDishes, err := readCountOfDishes()
	if err != nil {
		fmt.Println("Invalid input of count of dishes")

		return
	}

	if !validateCountOfDishes(countOfDishes) {
		return
	}

	dishHeap, err := readRatingsToHeap(countOfDishes)
	if err != nil {
		fmt.Println("Invalid input of rating of dish")

		return
	}

	numOfPreference, err := readNumOfPreference(countOfDishes)
	if err != nil {
		fmt.Println("Invalid input of num of preference")

		return
	}

	removeMinUntil(dishHeap, numOfPreference)

	if dishHeap.Len() != numOfPreference || dishHeap.Len() == 0 {
		fmt.Println("Heap size mismatch after trimming")

		return
	}

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
}
