package initial_test

import (
	"github.com/rbee3u/golib/stl/collections/sequence/initial"
)

func newList(items ...int) *initial.List[int] {
	return initial.NewList(items...)
}

func newIterator() initial.Iterator[int] {
	l := newList(0)
	return l.Begin()
}
