package algorithms

import (
	"github.com/rbee3u/golib/stl/iterators"
	"github.com/rbee3u/golib/stl/types"
)

func PartitionPoint(
	first, last iterators.ForwardIterator, pred types.UnaryPredicate,
) iterators.ForwardIterator {
	for count := iterators.Distance(first, last); count > 0; {
		mid, step := first, count/2
		mid = iterators.Advance(mid, step)

		if !pred(mid.Read()) {
			first = mid.Next().(iterators.InputIterator)
			count -= step + 1
		} else {
			count = step
		}
	}

	return first
}
