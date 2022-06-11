package vector

import "github.com/rbee3u/golib/stl/constraints"

var _ = constraints.IsMutableRandomAccessIterator[Iterator[int], int]()

type Iterator[T any] struct {
	l *List[T]
	n int
}

func (i Iterator[T]) Write(data T) {
	i.l.slice[i.n] = data
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
	return i.l.slice[i.n]
}

func (i Iterator[T]) Prev() Iterator[T] {
	i.n--

	return i
}

func (i Iterator[T]) Less(other Iterator[T]) bool {
	return i.n < other.n
}

func (i Iterator[T]) Advance(offset int) Iterator[T] {
	i.n += offset

	return i
}

func (i Iterator[T]) At(offset int) T {
	return i.l.slice[i.n+offset]
}

func (i Iterator[T]) Distance(other Iterator[T]) int {
	return other.n - i.n
}
