package singly_test

import (
	"github.com/rbee3u/golib/stl/collections/sequence/singly"
)

func newList() *singly.List {
	return singly.NewList()
}

func newIterator() singly.Iterator {
	l := newList()
	l.PushFront(0)
	return l.Begin()
}
