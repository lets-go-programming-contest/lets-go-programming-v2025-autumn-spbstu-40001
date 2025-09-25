package main;

import "container/heap";
import "fmt";

type IntHeap []int;

func (slice IntHeap) Len() int {
	return len(slice);
}
func (slice IntHeap) Less(i, j int) bool {
	return slice[i] < slice[j];
}
func (slice IntHeap) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i];
}
func (slice *IntHeap) Push(value any) {
	*slice = append(*slice, value.(int));
}
func (slice *IntHeap) Pop() any {
	result := (*slice)[len(*slice) - 1];
	*slice = (*slice)[0 : len(*slice) - 1];
	return result;
}
func (slice IntHeap) Top() int {
	return slice[0];
}

func main() {
	slice := &IntHeap{2, 1, 5};
	heap.Init(slice);
	heap.Push(slice, 3);
	fmt.Printf("Minimum = %d\n", slice.Top());
	for slice.Len() > 0 {
		fmt.Println(heap.Pop(slice));
	}
}
