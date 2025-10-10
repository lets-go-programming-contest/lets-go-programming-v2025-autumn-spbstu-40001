package main

import (
	"container/heap"
	"fmt"

	"github.com/atroxxxxxx/task-2-2/internal/intheap"
)

func Greater(a, b int) bool {
	return a > b
}

func main() {
	var nDishes uint

	_, err := fmt.Scan(&nDishes)

	dishes := intheap.New(make([]int, nDishes), Greater)

	heap.Init(&dishes)

	if err != nil {
		fmt.Println("failed to read dishes count:", err)

		return
	}

	if nDishes == 0 {
		fmt.Println("invalid dishes count")

		return
	}

	var dishRating int
	for range nDishes {
		_, err = fmt.Scan(&dishRating)
		if err != nil {
			fmt.Println("wrong dish rating format:", err)

			return
		}

		heap.Push(&dishes, dishRating)
	}

	var dishNumber uint

	_, err = fmt.Scan(&dishNumber)
	if err != nil {
		fmt.Println("failed to read dish number:", err)

		return
	}

	if dishNumber == 0 || dishNumber > nDishes {
		fmt.Println("invalid dish number")

		return
	}

	for range dishNumber - 1 {
		heap.Pop(&dishes)
	}

	fmt.Println(heap.Pop(&dishes))
}
