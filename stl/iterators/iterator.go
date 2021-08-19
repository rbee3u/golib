package iterators

import (
	"github.com/rbee3u/golib/stl/constraints"
	"github.com/rbee3u/golib/stl/types"
)

type Iterator interface {
	constraints.Cloneable
	constraints.Incrementable
}

func Distance(first, last InputIterator) (result types.Size) {
	firstRA, firstOK := first.(RandomAccessIterator)
	lastRA, lastOK := last.(RandomAccessIterator)

	if firstOK && lastOK {
		result = firstRA.Distance(lastRA)
	} else {
		for ; !first.Equal(last); first = (first.Next()).(InputIterator) {
			result++
		}
	}

	return result
}

func Advance(it InputIterator, n types.Size) InputIterator {
	if it, ok := it.(RandomAccessIterator); ok {
		return it.Advance(n)
	}

	if it, ok := it.(BidirectionalIterator); ok {
		for ; n > 0; n-- {
			it = it.Next().(BidirectionalIterator)
		}

		for ; n < 0; n++ {
			it = it.Prev().(BidirectionalIterator)
		}

		return it
	}

	for ; n > 0; n-- {
		it = it.Next().(BidirectionalIterator)
	}

	return it
}
