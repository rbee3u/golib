package singly_test

import (
	"github.com/rbee3u/golib/stl/collections/sequence/singly"
)

func newList() *singly.List[int] {
	return singly.NewList[int]()
}

func newIterator() singly.Iterator[int] {
	l := newList()
	l.PushFront(0)
	return l.Begin()
}
