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
		fmt.Println("scan dishes number error")

		return
	}

	ratingList := &intheap.IntHeap{}

	for range dishesNumber {
		var dishRating int

		_, err = fmt.Scan(&dishRating)
		if err != nil {
			fmt.Println("scan dish rating error")

			return
		}

		heap.Push(ratingList, dishRating)
	}

	var wishedDish int

	_, err = fmt.Scan(&wishedDish)
	if err != nil {
		fmt.Println("scan wished dish number error")

		return
	}

	if wishedDish <= dishesNumber {
		for range wishedDish - 1 {
			heap.Pop(ratingList)
		}
	} else {
		fmt.Println("invalid wished dish value")
		return
	}

	fmt.Println(heap.Pop(ratingList))
}
