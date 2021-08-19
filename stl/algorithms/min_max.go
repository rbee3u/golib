package algorithms

import (
	"github.com/rbee3u/golib/stl/constraints"
)

func Smallest[S constraints.ForwardIterator[S, T], T constraints.LessThanComparable[T]](
	first S, last S,
) S {
	return Best(first, last, func(x T, y T) bool { return x.Less(y) })
}

func Largest[S constraints.ForwardIterator[S, T], T constraints.LessThanComparable[T]](
	first S, last S,
) S {
	return Best(first, last, func(x T, y T) bool { return y.Less(x) })
}

func Best[S constraints.ForwardIterator[S, T], T any](
	first S, last S, better func(T, T) bool,
) S {
	best := first

	if !first.Equal(last) {
		for first = first.Next(); !first.Equal(last); first = first.Next() {
			if better(first.Read(), best.Read()) {
				best = first
			}
		}
	}

	return best
}
