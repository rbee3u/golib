package initial

import "github.com/rbee3u/golib/stl/constraints"

var _ = constraints.IsForwardIterator[Iterator[int], int]()

type Iterator[T any] struct {
	l *List[T]
	n int
}

func (i Iterator[T]) Clone() Iterator[T] {
	return i
}

func (i Iterator[T]) Next() Iterator[T] {
	i.n++

	return i
}

func (i Iterator[T]) Equal(other Iterator[T]) bool {
	return i == other
}

func (i Iterator[T]) Read() T {
	return i.l.items[i.n]
}
