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
	var nDishes int

	_, err := fmt.Scan(&nDishes)
	if err != nil {
		fmt.Println("Failed to read dishes count")
		fmt.Println(err)

		return
	}

	dishesQueue := heap.New[int](greater[int])

	for range nDishes {
		var dishValue int

		_, err = fmt.Scan(&dishValue)
		if err != nil {
			fmt.Println("Failed to read dish value")
			fmt.Println(err)

			return
		}

		dishesQueue.Push(dishValue)
	}

	var dishID int

	_, err = fmt.Scan(&dishID)
	if err != nil {
		fmt.Println("Failed to read priority number")
		fmt.Println(err)

		return
	}

	if (dishID > dishesQueue.Len()) || (dishID <= 0) {
		fmt.Println("Entered nonexistent priority number")

		return
	}

	for range dishID - 1 {
		_, _ = dishesQueue.Pop()
	}

	result, _ := dishesQueue.Top()
	fmt.Println(result)
}
