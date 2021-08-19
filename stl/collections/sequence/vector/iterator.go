package vector

import (
	"github.com/rbee3u/golib/stl/constraints"
	"github.com/rbee3u/golib/stl/iterators"
	"github.com/rbee3u/golib/stl/types"
)

var _ iterators.RandomAccessIterator = Iterator{}

type Iterator struct {
	l *List
	n types.Size
}

func (i Iterator) Write(data types.Data) {
	i.l.slice[i.n] = data
}

func (i Iterator) Clone() constraints.Cloneable {
	return i.ImplClone()
}

func (i Iterator) ImplClone() Iterator {
	return i
}

func (i Iterator) Next() constraints.Incrementable {
	return i.ImplNext()
}

func (i Iterator) ImplNext() Iterator {
	i.n++

	return i
}

func (i Iterator) Equal(other constraints.EqualityComparable) bool {
	return i.ImplEqual(other.(Iterator))
}

func (i Iterator) ImplEqual(other Iterator) bool {
	return i == other
}

func (i Iterator) Read() types.Data {
	return i.l.slice[i.n]
}

func (i Iterator) Prev() constraints.Decrementable {
	return i.ImplPrev()
}

func (i Iterator) ImplPrev() Iterator {
	i.n--

	return i
}

func (i Iterator) Less(other constraints.LessThanComparable) bool {
	return i.ImplLess(other.(Iterator))
}

func (i Iterator) ImplLess(other Iterator) bool {
	return i.n < other.n
}

func (i Iterator) At(diff types.Size) types.Data {
	return i.l.slice[i.n+diff]
}

func (i Iterator) Advance(diff types.Size) iterators.RandomAccessIterator {
	return i.ImplAdvance(diff)
}

func (i Iterator) ImplAdvance(diff types.Size) Iterator {
	i.n += diff

	return i
}

func (i Iterator) Distance(other iterators.RandomAccessIterator) types.Size {
	return i.ImplDistance(other.(Iterator))
}

func (i Iterator) ImplDistance(other Iterator) types.Size {
	return other.n - i.n
}
