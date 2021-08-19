package doubly_test

import (
	"github.com/rbee3u/golib/stl/collections/sequence/doubly"
)

func newList() *doubly.List {
	return doubly.NewList()
}

func newIterator() doubly.Iterator {
	l := newList()
	l.PushBack(0)
	return l.Begin()
}
