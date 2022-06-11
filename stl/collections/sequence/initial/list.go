package initial

type List[T any] struct {
	items []T
}

func NewList[T any](items ...T) *List[T] {
	return &List[T]{items: items}
}

func (l *List[T]) Size() int {
	return len(l.items)
}

func (l *List[T]) Empty() bool {
	return l.Size() == 0
}

func (l *List[T]) Begin() Iterator[T] {
	return Iterator[T]{l: l, n: 0}
}

func (l *List[T]) End() Iterator[T] {
	return Iterator[T]{l: l, n: len(l.items)}
}
