package primheap

import "cmp"

type extendedStack[T cmp.Ordered] struct {
	data       []T
	comparator func(T, T) bool
}

func (objHeap *extendedStack[T]) Less(i, j int) bool {
	if objHeap.comparator != nil {
		return objHeap.comparator(objHeap.data[i], objHeap.data[j])
	}

	return cmp.Less(objHeap.data[i], objHeap.data[j])
}

func (objHeap *extendedStack[T]) Len() int {
	return len(objHeap.data)
}

func (objHeap *extendedStack[T]) Swap(i, j int) {
	objHeap.data[i], objHeap.data[j] = objHeap.data[j], objHeap.data[i]
}

func (objHeap *extendedStack[T]) Push(value any) {
	intValue, ok := value.(T)
	if !ok {
		panic("failed to cast heap.push value to stored type")
	}

	objHeap.data = append(objHeap.data, intValue)
}

func (objHeap *extendedStack[T]) Pop() any {
	if objHeap.Len() == 0 {
		return nil
	}

	var empty T

	objHeap.data[objHeap.Len()-1] = empty
	result := objHeap.data[objHeap.Len()-1]
	objHeap.data = objHeap.data[0 : objHeap.Len()-1]

	return result
}
