package algorithms

import (
	"github.com/rbee3u/golib/stl/constraints"
)

func LowerBound[S constraints.RandomAccessIterator[S, T], T constraints.LessThanComparable[T]](
	first S, last S, target T,
) S {
	return PartitionPoint(first, last, func(x T) bool { return !x.Less(target) })
}

func UpperBound[S constraints.RandomAccessIterator[S, T], T constraints.LessThanComparable[T]](
	first S, last S, target T,
) S {
	return PartitionPoint(first, last, func(x T) bool { return target.Less(x) })
}

func PartitionPoint[S constraints.RandomAccessIterator[S, T], T any](
	first S, last S, pred func(T) bool,
) S {
	for count := first.Distance(last); count > 0; {
		mid, step := first, count/2
		mid = mid.Advance(step)

		if !pred(mid.Read()) {
			first = mid.Next()
			count -= step + 1
		} else {
			count = step
		}
	}

	return first
}
