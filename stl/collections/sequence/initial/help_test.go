package initial_test

import (
	"github.com/rbee3u/golib/stl/collections/sequence/initial"
	"github.com/rbee3u/golib/stl/types"
)

func newList(datas ...types.Data) *initial.List {
	return initial.NewList(datas...)
}

func newIterator() initial.Iterator {
	l := newList(0)
	return l.Begin()
}
