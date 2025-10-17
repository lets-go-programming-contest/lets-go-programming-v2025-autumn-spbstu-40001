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

	var wishedDish int

	_, err = fmt.Scan(&wishedDish)
	if err != nil {
		fmt.Println("scan wished dish number error", err)

		return
	}

	if wishedDish > dishesNumber {
		fmt.Println("invalid wished dish value")

		return
	}

	for range wishedDish - 1 {
		heap.Pop(ratingList)
	}

	fmt.Println(heap.Pop(ratingList))
}
