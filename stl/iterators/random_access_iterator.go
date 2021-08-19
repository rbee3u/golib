package iterators

import (
	"github.com/rbee3u/golib/stl/constraints"
	"github.com/rbee3u/golib/stl/types"
)

type RandomAccessIterator interface {
	BidirectionalIterator
	constraints.LessThanComparable
	At(diff types.Size) types.Data
	Advance(diff types.Size) RandomAccessIterator
	Distance(other RandomAccessIterator) types.Size
}
