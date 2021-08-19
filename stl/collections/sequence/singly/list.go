package singly

import (
	"github.com/rbee3u/golib/stl/types"
)

type List struct {
	sentinel node
	size     types.Size
}

type node struct {
	next *node
	data types.Data
}

func NewList() *List {
	l := &List{}

	return l
}

func (l *List) Size() types.Size {
	return l.size
}

func (l *List) Empty() bool {
	return l.Size() == 0
}

func (l *List) PushFront(data types.Data) {
	s := &l.sentinel
	s.next = &node{next: s.next, data: data}
	l.size++
}

func (l *List) Front() types.Data {
	return l.sentinel.next.data
}

func (l *List) PopFront() {
	s := &l.sentinel
	s.next = s.next.next
	l.size--
}

func (l *List) Clear() {
	l.sentinel.next = nil
	l.size = 0
}

func (l *List) Begin() Iterator {
	return Iterator{n: l.sentinel.next}
}

func (l *List) End() Iterator {
	return Iterator{}
}

func (l *List) InsertAfter(i Iterator, data types.Data) Iterator {
	i.n.next = &node{next: i.n.next, data: data}
	l.size++

	return Iterator{n: i.n.next}
}

func (l *List) EraseAfter(i Iterator) Iterator {
	i.n.next = i.n.next.next
	l.size--

	return Iterator{n: i.n.next}
}
