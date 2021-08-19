package rb

import (
	"github.com/rbee3u/golib/stl/constraints"
)

var _ = constraints.IsMutableBidirectionalIterator[Iterator[int], int]()

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
	i.n = successor(i.n)

	return i
}

func (i Iterator[T]) Equal(other Iterator[T]) bool {
	return i == other
}

func (i Iterator[T]) Read() T {
	return i.n.data
}

func (i Iterator[T]) Prev() Iterator[T] {
	i.n = predecessor(i.n)

	return i
}
