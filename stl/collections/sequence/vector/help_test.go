package vector_test

import (
	"github.com/rbee3u/golib/stl/collections/sequence/vector"
)

func newList() *vector.List {
	return vector.NewList()
}

func newIterator() vector.Iterator {
	l := newList()
	l.PushBack(0)
	return l.Begin()
}
