package doubly_test

import (
	"github.com/rbee3u/golib/stl/collections/sequence/doubly"
)

func newList() *doubly.List[int] {
	return doubly.NewList[int]()
}

func newIterator() doubly.Iterator[int] {
	l := newList()
	l.PushBack(0)
	return l.Begin()
}
