package main

import (
	"cmp"
	"fmt"

	heap "github.com/Rychmick/task-2-2/pkg/primheap"
)

func greater[T cmp.Ordered](lhs, rhs T) bool {
	return lhs > rhs
}

func main() {
	var nDishes uint

	_, err := fmt.Scan(&nDishes)
	if err != nil {
		fmt.Println("Failed to read dishes count:", err)

		return
	}

	dishesQueue := heap.New[int](greater[int])

	for range nDishes {
		var dishValue int

		_, err = fmt.Scan(&dishValue)
		if err != nil {
			fmt.Println("Failed to read dish value:", err)

			return
		}

		dishesQueue.Push(dishValue)
	}

	var dishID int

	_, err = fmt.Scan(&dishID)
	if err != nil {
		fmt.Println("Failed to read priority number:", err)

		return
	}

	result, err := dishesQueue.PopN(dishID)
	if err != nil {
		fmt.Println("Failed to get priority with entered number:", err)

		return
	}

	fmt.Println(result)
}
