package algorithms

import (
	"github.com/rbee3u/golib/stl/constraints"
)

func MakeMinHeap[S constraints.MutableRandomAccessIterator[S, T], T constraints.LessThanComparable[T]](
	first S, last S,
) {
	MakeHeap(first, last, func(x T, y T) bool { return x.Less(y) })
}

func PushMinHeap[S constraints.MutableRandomAccessIterator[S, T], T constraints.LessThanComparable[T]](
	first S, last S,
) {
	PushHeap(first, last, func(x T, y T) bool { return x.Less(y) })
}

func PopMinHeap[S constraints.MutableRandomAccessIterator[S, T], T constraints.LessThanComparable[T]](
	first S, last S,
) {
	PopHeap(first, last, func(x T, y T) bool { return x.Less(y) })
}

func SortMinHeap[S constraints.MutableRandomAccessIterator[S, T], T constraints.LessThanComparable[T]](
	first S, last S,
) {
	SortHeap(first, last, func(x T, y T) bool { return x.Less(y) })
}

func IsMinHeap[S constraints.MutableRandomAccessIterator[S, T], T constraints.LessThanComparable[T]](
	first S, last S,
) bool {
	return IsHeap(first, last, func(x T, y T) bool { return x.Less(y) })
}

func IsMinHeapUntil[S constraints.MutableRandomAccessIterator[S, T], T constraints.LessThanComparable[T]](
	first S, last S,
) S {
	return IsHeapUntil(first, last, func(x T, y T) bool { return x.Less(y) })
}

func MakeMaxHeap[S constraints.MutableRandomAccessIterator[S, T], T constraints.LessThanComparable[T]](
	first S, last S,
) {
	MakeHeap(first, last, func(x T, y T) bool { return y.Less(x) })
}

func PushMaxHeap[S constraints.MutableRandomAccessIterator[S, T], T constraints.LessThanComparable[T]](
	first S, last S,
) {
	PushHeap(first, last, func(x T, y T) bool { return y.Less(x) })
}

func PopMaxHeap[S constraints.MutableRandomAccessIterator[S, T], T constraints.LessThanComparable[T]](
	first S, last S,
) {
	PopHeap(first, last, func(x T, y T) bool { return y.Less(x) })
}

func SortMaxHeap[S constraints.MutableRandomAccessIterator[S, T], T constraints.LessThanComparable[T]](
	first S, last S,
) {
	SortHeap(first, last, func(x T, y T) bool { return y.Less(x) })
}

func IsMaxHeap[S constraints.MutableRandomAccessIterator[S, T], T constraints.LessThanComparable[T]](
	first S, last S,
) bool {
	return IsHeap(first, last, func(x T, y T) bool { return y.Less(x) })
}

func IsMaxHeapUntil[S constraints.MutableRandomAccessIterator[S, T], T constraints.LessThanComparable[T]](
	first S, last S,
) S {
	return IsHeapUntil(first, last, func(x T, y T) bool { return y.Less(x) })
}

func MakeHeap[S constraints.MutableRandomAccessIterator[S, T], T any](
	first S, last S, less func(T, T) bool,
) {
	if n := first.Distance(last); n > 1 {
		for start := (n - 2) / 2; start >= 0; start-- {
			siftDown(first, last, less, n, first.Advance(start))
		}
	}
}

func PushHeap[S constraints.MutableRandomAccessIterator[S, T], T any](
	first S, last S, less func(T, T) bool,
) {
	siftUp(first, last, less, first.Distance(last))
}

func PopHeap[S constraints.MutableRandomAccessIterator[S, T], T any](
	first S, last S, less func(T, T) bool,
) {
	popHeap(first, last, less, first.Distance(last))
}

func SortHeap[S constraints.MutableRandomAccessIterator[S, T], T any](
	first S, last S, less func(T, T) bool,
) {
	sortHeap(first, last, less)
}

func IsHeap[S constraints.MutableRandomAccessIterator[S, T], T any](
	first S, last S, less func(T, T) bool,
) bool {
	return IsHeapUntil(first, last, less).Equal(last)
}

func IsHeapUntil[S constraints.MutableRandomAccessIterator[S, T], T any](
	first S, last S, less func(T, T) bool,
) S {
	length := first.Distance(last)
	p, c := 0, 1
	pp := first

	for c < length {
		cp := first.Advance(c)
		if less(pp.Read(), cp.Read()) {
			return cp
		}

		c++

		cp = cp.Next()

		if c == length {
			return last
		}

		if less(pp.Read(), cp.Read()) {
			return cp
		}

		p++

		pp = pp.Next()
		c = 2*p + 1
	}

	return last
}

func popHeap[S constraints.MutableRandomAccessIterator[S, T], T any](
	first S, last S, less func(T, T) bool, length int,
) {
	if length > 1 {
		last = last.Prev()
		firstData, lastData := first.Read(), last.Read()
		first.Write(lastData)
		last.Write(firstData)
		siftDown(first, last, less, length-1, first)
	}
}

func sortHeap[S constraints.MutableRandomAccessIterator[S, T], T any](
	first S, last S, less func(T, T) bool,
) {
	for n := first.Distance(last); n > 1; n-- {
		popHeap(first, last, less, n)
		last = last.Prev()
	}
}

func siftDown[S constraints.MutableRandomAccessIterator[S, T], T any](
	first S, _ S, less func(T, T) bool, length int, start S,
) {
	child := first.Distance(start)
	if length < 2 || (length-2)/2 < child {
		return
	}

	child = child*2 + 1
	childIt := first.Advance(child)

	if (child+1) < length && less(childIt.Read(), childIt.Next().Read()) {
		// Right child exists and is greater than left child.
		childIt = childIt.Next()
		child++
	}

	// Check if we are in heap-order.
	if less(childIt.Read(), start.Read()) {
		return
	}

	for top := start.Read(); ; {
		// We are not in heap-order, swap the parent with its largest child.
		start.Write(childIt.Read())
		start = childIt

		if (length-2)/2 < child {
			break
		}

		// Recompute the child based off of the updated parent.
		child = child*2 + 1
		childIt = first.Advance(child)

		if (child+1) < length && less(childIt.Read(), childIt.Next().Read()) {
			// Right child exists and is greater than left child.
			childIt = childIt.Next()
			child++
		}

		if less(childIt.Read(), top) {
			break
		}
	}
}

func siftUp[S constraints.MutableRandomAccessIterator[S, T], T any](
	first S, last S, less func(T, T) bool, length int,
) {
	if length > 1 {
		length = (length - 2) / 2
		ptr := first.Advance(length)
		last = last.Prev()

		if less(ptr.Read(), last.Read()) {
			t := last.Read()

			for {
				last.Write(ptr.Read())
				last = ptr

				if length == 0 {
					break
				}

				length = (length - 1) / 2
				ptr = first.Advance(length)

				if !less(ptr.Read(), t) {
					break
				}
			}

			last.Write(t)
		}
	}
}
