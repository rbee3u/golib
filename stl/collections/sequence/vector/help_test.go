package vector_test

import (
	"github.com/rbee3u/golib/stl/collections/sequence/vector"
)

func newList() *vector.List[int] {
	return vector.NewList[int]()
}

func newIterator() vector.Iterator[int] {
	l := newList()
	l.PushBack(0)
	return l.Begin()
}
