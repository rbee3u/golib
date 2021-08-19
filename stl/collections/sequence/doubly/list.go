package doubly

type List[T any] struct {
	sentinel node[T]
	size     int
}

type node[T any] struct {
	prev *node[T]
	next *node[T]
	data T
}

func NewList[T any]() *List[T] {
	l := &List[T]{}
	l.sentinel.prev = &l.sentinel
	l.sentinel.next = &l.sentinel

	return l
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) Empty() bool {
	return l.Size() == 0
}

func (l *List[T]) PushFront(data T) {
	p := &node[T]{data: data}
	l.link(l.sentinel.next, p, p)
	l.size++
}

func (l *List[T]) Front() T {
	return l.sentinel.next.data
}

func (l *List[T]) PopFront() {
	p := l.sentinel.next
	l.unlink(p, p)
	l.size--
}

func (l *List[T]) PushBack(data T) {
	p := &node[T]{data: data}
	l.link(&l.sentinel, p, p)
	l.size++
}

func (l *List[T]) Back() T {
	return l.sentinel.prev.data
}

func (l *List[T]) PopBack() {
	p := l.sentinel.prev
	l.unlink(p, p)
	l.size--
}

func (l *List[T]) Clear() {
	s := &l.sentinel
	s.prev, s.next = s, s
	l.size = 0
}

func (l *List[T]) Begin() Iterator[T] {
	return Iterator[T]{n: l.sentinel.next}
}

func (l *List[T]) End() Iterator[T] {
	return Iterator[T]{n: &l.sentinel}
}

func (l *List[T]) ReverseBegin() Iterator[T] {
	return Iterator[T]{n: l.sentinel.prev}
}

func (l *List[T]) ReverseEnd() Iterator[T] {
	return Iterator[T]{n: &l.sentinel}
}

func (l *List[T]) Insert(i Iterator[T], data T) Iterator[T] {
	p := &node[T]{data: data}
	l.link(i.n, p, p)
	l.size++

	return Iterator[T]{n: p}
}

func (l *List[T]) Erase(i Iterator[T]) Iterator[T] {
	p := i.n
	r := p.next
	l.unlink(p, p)
	l.size--

	return Iterator[T]{n: r}
}

func (l *List[T]) link(p, head, tail *node[T]) {
	p.prev.next = head
	head.prev = p.prev
	p.prev = tail
	tail.next = p
}

func (l *List[T]) unlink(head, tail *node[T]) {
	head.prev.next = tail.next
	tail.next.prev = head.prev
}
