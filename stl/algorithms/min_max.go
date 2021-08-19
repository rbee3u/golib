package algorithms

import (
	"github.com/rbee3u/golib/stl/iterators"
	"github.com/rbee3u/golib/stl/types"
)

func MinElement(
	first, last iterators.ForwardIterator, less types.BinaryPredicate,
) iterators.ForwardIterator {
	if first.Equal(last) {
		return last
	}

	smallest := first
	first = first.Next().(iterators.ForwardIterator)

	for !first.Equal(last) {
		if less(first.Read(), smallest.Read()) {
			smallest = first
		}

		first = first.Next().(iterators.ForwardIterator)
	}

	return smallest
}

func MaxElement(
	first, last iterators.ForwardIterator, less types.BinaryPredicate,
) iterators.ForwardIterator {
	if first.Equal(last) {
		return last
	}

	largest := first
	first = first.Next().(iterators.ForwardIterator)

	for !first.Equal(last) {
		if less(largest.Read(), first.Read()) {
			largest = first
		}

		first = first.Next().(iterators.ForwardIterator)
	}

	return largest
}
