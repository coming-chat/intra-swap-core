package util

type Item[T any] struct {
	Value T
	Index int
}

func NewReverseHeap[T any](fixedLen int, less func(a, b Item[T]) bool) *ReverseHeap[T] {
	return &ReverseHeap[T]{
		Items: make([]Item[T], fixedLen),
		less:  less,
	}
}

type ReverseHeap[T any] struct {
	Items []Item[T]
	less  func(a, b Item[T]) bool
}

func (r *ReverseHeap[T]) Push(item any) {
	n := len(r.Items)
	newItem := Item[T]{
		Value: item.(T),
		Index: n,
	}
	r.Items = append(r.Items, newItem)
}

func (r *ReverseHeap[T]) Pop() any {
	oldItems := r.Items
	n := len(oldItems)
	item := oldItems[n-1]
	item.Index = -1
	r.Items = oldItems[0 : n-1]
	return item
}

func (r *ReverseHeap[T]) Len() int {
	return len(r.Items)
}

func (r *ReverseHeap[T]) Swap(i, j int) {
	r.Items[i], r.Items[j] = r.Items[j], r.Items[i]
}

func (r *ReverseHeap[T]) Less(i, j int) bool {
	return r.less(r.Items[i], r.Items[j])
}
