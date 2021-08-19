package algorithms

import (
	"github.com/rbee3u/golib/stl/iterators"
	"github.com/rbee3u/golib/stl/types"
)

func LowerBound(
	first, last iterators.ForwardIterator, value types.Data, less types.BinaryPredicate,
) iterators.ForwardIterator {
	for count := iterators.Distance(first, last); count > 0; {
		mid, step := first, count/2
		mid = iterators.Advance(mid, step)

		if less(mid.Read(), value) {
			first = mid.Next().(iterators.InputIterator)
			count -= step + 1
		} else {
			count = step
		}
	}

	return first
}

func UpperBound(
	first, last iterators.ForwardIterator, value types.Data, less types.BinaryPredicate,
) iterators.ForwardIterator {
	for count := iterators.Distance(first, last); count > 0; {
		mid, step := first, count/2
		mid = iterators.Advance(mid, step)

		if !less(value, mid.Read()) {
			first = mid.Next().(iterators.InputIterator)
			count -= step + 1
		} else {
			count = step
		}
	}

	return first
}
