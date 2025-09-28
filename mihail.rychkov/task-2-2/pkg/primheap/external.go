package primheap;

import "cmp";
import "errors"
import "container/heap";

type PrimHeap[T cmp.Ordered] struct {
	heap unwrappedPrimHeap[T];
}
func (obj *PrimHeap[T]) Init() {
	heap.Init(&obj.heap);
}
func (obj *PrimHeap[T]) Len() int {
	return obj.heap.Len();
}

var ErrEmptyHeapTop = errors.New("cannot get top element from empty heap");
func (obj *PrimHeap[T]) Top() (T, error) {
	if (obj.heap.Len() == 0) {
		var result T;
		return result, ErrEmptyHeapTop;
	}
	return obj.heap.data[0], nil;
}
func (obj *PrimHeap[T]) Push(value T) {
	heap.Push(&obj.heap, value);
}

var ErrHeapUnderflow = errors.New("cannot pop from empty heap");
func (obj *PrimHeap[T]) Pop() (T, error) {
	result := heap.Pop(&obj.heap);
	if (result == nil) {
		var result T;
		return result, ErrHeapUnderflow;
	}
	castedResult, ok := result.(T);
	if (!ok) {
		panic("heap.Pop returned any with unexpected type");
	}
	return castedResult, nil;
}

func New[T cmp.Ordered](less func(T, T) bool, values ...T) PrimHeap[T] {
	result := PrimHeap[T]{unwrappedPrimHeap[T]{values, less}};
	result.Init();
	return result;
}
