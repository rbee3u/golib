package singly

type List[T any] struct {
	sentinel node[T]
	size     int
}

type node[T any] struct {
	next *node[T]
	data T
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) Empty() bool {
	return l.Size() == 0
}

func (l *List[T]) PushFront(data T) {
	s := &l.sentinel
	s.next = &node[T]{next: s.next, data: data}
	l.size++
}

func (l *List[T]) Front() T {
	return l.sentinel.next.data
}

func (l *List[T]) PopFront() {
	s := &l.sentinel
	s.next = s.next.next
	l.size--
}

func (l *List[T]) Clear() {
	l.sentinel.next = nil
	l.size = 0
}

func (l *List[T]) Begin() Iterator[T] {
	return Iterator[T]{n: l.sentinel.next}
}

func (l *List[T]) End() Iterator[T] {
	return Iterator[T]{}
}

func (l *List[T]) InsertAfter(i Iterator[T], data T) Iterator[T] {
	i.n.next = &node[T]{next: i.n.next, data: data}
	l.size++

	return Iterator[T]{n: i.n.next}
}

func (l *List[T]) EraseAfter(i Iterator[T]) Iterator[T] {
	i.n.next = i.n.next.next
	l.size--

	return Iterator[T]{n: i.n.next}
}
