package dict

import (
	base "github.com/rbee3u/golib/stl/collections/associative/avl"
	"github.com/rbee3u/golib/stl/constraints"
)

var _ = constraints.IsMutableBidirectionalIterator[Iterator[Pair[string, int]], Pair[string, int]]()

type Iterator[T any] struct {
	base base.Iterator[T]
}

func (i Iterator[T]) Write(data T) {
	i.base.Write(data)
}

func (i Iterator[T]) Clone() Iterator[T] {
	return i
}

func (i Iterator[T]) Next() Iterator[T] {
	i.base = i.base.Next()

	return i
}

func (i Iterator[T]) Equal(other Iterator[T]) bool {
	return i == other
}

func (i Iterator[T]) Read() T {
	return i.base.Read()
}

func (i Iterator[T]) Prev() Iterator[T] {
	i.base = i.base.Prev()

	return i
}
