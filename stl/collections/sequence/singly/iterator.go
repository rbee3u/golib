package singly

import "github.com/rbee3u/golib/stl/constraints"

var _ = constraints.IsMutableForwardIterator[Iterator[int], int]()

type Iterator[T any] struct {
	n *node[T]
}

func (i Iterator[T]) Write(data T) {
	i.n.data = data
}

func (i Iterator[T]) Clone() Iterator[T] {
	return i
}

func (i Iterator[T]) Next() Iterator[T] {
	i.n = i.n.next

	return i
}

func (i Iterator[T]) Equal(other Iterator[T]) bool {
	return i == other
}

func (i Iterator[T]) Read() T {
	return i.n.data
}
