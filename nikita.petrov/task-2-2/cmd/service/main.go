package main

import (
	"container/heap"
	"fmt"

	"github.com/Nekich06/task-2-2/internal/intheap"
)

func main() {
	var dishesNumber int

	_, err := fmt.Scan(&dishesNumber)
	if err != nil {
		fmt.Println("scan dishes number error", err)

		return
	}

	ratingList := &intheap.IntHeap{}

	for range dishesNumber {
		var dishRating int

		_, err = fmt.Scan(&dishRating)
		if err != nil {
			fmt.Println("scan dish rating error", err)

			return
		}

		heap.Push(ratingList, dishRating)
	}

	var wishedDishNumber int

	_, err = fmt.Scan(&wishedDishNumber)
	if err != nil {
		fmt.Println("scan wished dish number error", err)

		return
	}

	if wishedDishNumber > dishesNumber {
		fmt.Println("invalid wished dish number value")

		return
	}

	for range wishedDishNumber - 1 {
		heap.Pop(ratingList)
	}

	wishedDish, ok := heap.Pop(ratingList).(int)

	if !ok {
		fmt.Println("cannot convert to int (heap is empty)")

		return
	}

	fmt.Println(wishedDish)
}
