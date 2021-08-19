package doubly

import (
	"github.com/rbee3u/golib/stl/types"
)

type List struct {
	sentinel node
	size     types.Size
}

type node struct {
	prev *node
	next *node
	data types.Data
}

func NewList() *List {
	l := &List{}
	l.sentinel.prev = &l.sentinel
	l.sentinel.next = &l.sentinel

	return l
}

func (l *List) Size() types.Size {
	return l.size
}

func (l *List) Empty() bool {
	return l.Size() == 0
}

func (l *List) PushFront(data types.Data) {
	p := &node{data: data}
	l.link(l.sentinel.next, p, p)
	l.size++
}

func (l *List) Front() types.Data {
	return l.sentinel.next.data
}

func (l *List) PopFront() {
	p := l.sentinel.next
	l.unlink(p, p)
	l.size--
}

func (l *List) PushBack(data types.Data) {
	p := &node{data: data}
	l.link(&l.sentinel, p, p)
	l.size++
}

func (l *List) Back() types.Data {
	return l.sentinel.prev.data
}

func (l *List) PopBack() {
	p := l.sentinel.prev
	l.unlink(p, p)
	l.size--
}

func (l *List) Clear() {
	s := &l.sentinel
	s.prev, s.next = s, s
	l.size = 0
}

func (l *List) Begin() Iterator {
	return Iterator{n: l.sentinel.next}
}

func (l *List) End() Iterator {
	return Iterator{n: &l.sentinel}
}

func (l *List) ReverseBegin() Iterator {
	return Iterator{n: l.sentinel.prev}
}

func (l *List) ReverseEnd() Iterator {
	return Iterator{n: &l.sentinel}
}

func (l *List) Insert(i Iterator, data types.Data) Iterator {
	p := &node{data: data}
	l.link(i.n, p, p)
	l.size++

	return Iterator{n: p}
}

func (l *List) Erase(i Iterator) Iterator {
	p := i.n
	r := p.next
	l.unlink(p, p)
	l.size--

	return Iterator{n: r}
}

func (l *List) link(p, head, tail *node) {
	p.prev.next = head
	head.prev = p.prev
	p.prev = tail
	tail.next = p
}

func (l *List) unlink(head, tail *node) {
	head.prev.next = tail.next
	tail.next.prev = head.prev
}
