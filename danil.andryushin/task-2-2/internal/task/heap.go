package task

import "cmp"

type Heap struct {
	Functor func(int, int) bool
	data    []int
}

func (obj *Heap) Len() int {
	return len(obj.data)
}

func (obj *Heap) Less(lhs, rhs int) bool {
	if lhs >= obj.Len() || rhs >= obj.Len() {
		panic("invalid index range")
	}
	if obj.Functor == nil {
		return cmp.Less(obj.data[lhs], obj.data[rhs])
	}
	return obj.Functor(obj.data[lhs], obj.data[rhs])
}

func (obj *Heap) Swap(lhs, rhs int) {
	obj.data[lhs], obj.data[rhs] = obj.data[rhs], obj.data[lhs]
}

func (obj *Heap) Push(data any) {
	iData, casted := data.(int)
	if !casted {
		panic("unexpected data type (expected: int)")
	}
	obj.data = append(obj.data, iData)
}

func (obj *Heap) Pop() any {
	size := obj.Len()
	if size == 0 {
		panic("heap underflow")
	}
	data := obj.data[size-1]
	obj.data = obj.data[0 : size-1]
	return data
}
