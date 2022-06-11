package vector

type List[T any] struct {
	slice []T
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Size() int {
	return len(l.slice)
}

func (l *List[T]) Empty() bool {
	return l.Size() == 0
}

func (l *List[T]) Get(i int) T {
	return l.slice[i]
}

func (l *List[T]) Set(n int, data T) {
	l.slice[n] = data
}

func (l *List[T]) PushBack(data T) {
	l.slice = append(l.slice, data)
}

func (l *List[T]) Back() T {
	return l.slice[len(l.slice)-1]
}

func (l *List[T]) PopBack() {
	l.slice = l.slice[:len(l.slice)-1]
}

func (l *List[T]) Clear() {
	l.slice = nil
}

func (l *List[T]) Begin() Iterator[T] {
	return Iterator[T]{l: l, n: 0}
}

func (l *List[T]) End() Iterator[T] {
	return Iterator[T]{l: l, n: len(l.slice)}
}

func (l *List[T]) ReverseBegin() Iterator[T] {
	return Iterator[T]{l: l, n: len(l.slice) - 1}
}

func (l *List[T]) ReverseEnd() Iterator[T] {
	return Iterator[T]{l: l, n: -1}
}
