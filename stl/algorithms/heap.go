package algorithms

import (
	"github.com/rbee3u/golib/stl/iterators"
	"github.com/rbee3u/golib/stl/types"
)

func MakeHeap(first, last iterators.MutableRandomAccessIterator, less types.BinaryPredicate) {
	if n := first.Distance(last); n > 1 {
		for start := (n - 2) / 2; start >= 0; start-- {
			siftDown(first, last, less, n, first.Advance(start).(iterators.MutableRandomAccessIterator))
		}
	}
}

func PushHeap(first, last iterators.MutableRandomAccessIterator, less types.BinaryPredicate) {
	siftUp(first, last, less, first.Distance(last))
}

func PopHeap(first, last iterators.MutableRandomAccessIterator, less types.BinaryPredicate) {
	popHeap(first, last, less, first.Distance(last))
}

func popHeap(first, last iterators.MutableRandomAccessIterator, less types.BinaryPredicate,
	length types.Size) {
	if length > 1 {
		last = last.Prev().(iterators.MutableRandomAccessIterator)
		firstData, lastData := first.Read(), last.Read()
		first.Write(lastData)
		last.Write(firstData)
		siftDown(first, last, less, length-1, first)
	}
}

func SortHeap(first, last iterators.MutableRandomAccessIterator, less types.BinaryPredicate) {
	sortHeap(first, last, less)
}

func sortHeap(first, last iterators.MutableRandomAccessIterator, less types.BinaryPredicate) {
	for n := first.Distance(last); n > 1; n-- {
		popHeap(first, last, less, n)
		last = last.Prev().(iterators.MutableRandomAccessIterator)
	}
}

//nolint:cyclop
func siftDown(first, _ iterators.MutableRandomAccessIterator, less types.BinaryPredicate,
	length types.Size, start iterators.MutableRandomAccessIterator) {
	child := first.Distance(start)
	if length < 2 || (length-2)/2 < child {
		return
	}

	child = child*2 + 1
	childIt := first.Advance(child).(iterators.MutableRandomAccessIterator)

	if (child+1) < length && less(childIt.Read(), childIt.Next().(iterators.MutableRandomAccessIterator).Read()) {
		// Right child exists and is greater than left child.
		childIt = childIt.Next().(iterators.MutableRandomAccessIterator)
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
		childIt = first.Advance(child).(iterators.MutableRandomAccessIterator)

		if (child+1) < length && less(childIt.Read(), childIt.Next().(iterators.MutableRandomAccessIterator).Read()) {
			// Right child exists and is greater than left child.
			childIt = childIt.Next().(iterators.MutableRandomAccessIterator)
			child++
		}

		if less(childIt.Read(), top) {
			break
		}
	}
}

//nolint:nestif
func siftUp(first, last iterators.MutableRandomAccessIterator, less types.BinaryPredicate,
	length types.Size) {
	if length > 1 {
		length = (length - 2) / 2
		ptr := first.Advance(length).(iterators.MutableRandomAccessIterator)
		last = last.Prev().(iterators.MutableRandomAccessIterator)

		if less(ptr.Read(), last.Read()) {
			t := last.Read()

			for {
				last.Write(ptr.Read())
				last = ptr

				if length == 0 {
					break
				}

				length = (length - 1) / 2
				ptr = first.Advance(length).(iterators.MutableRandomAccessIterator)

				if !less(ptr.Read(), t) {
					break
				}
			}

			last.Write(t)
		}
	}
}

func IsHeap(first, last iterators.RandomAccessIterator, less types.BinaryPredicate) bool {
	return IsHeapUntil(first, last, less) == last
}

func IsHeapUntil(first, last iterators.RandomAccessIterator, less types.BinaryPredicate,
) iterators.RandomAccessIterator {
	length := first.Distance(last)
	p, c := 0, 1
	pp := first

	for c < length {
		cp := first.Advance(c)
		if less(pp.Read(), cp.Read()) {
			return cp
		}

		c++

		cp = cp.Next().(iterators.RandomAccessIterator)

		if c == length {
			return last
		}

		if less(pp.Read(), cp.Read()) {
			return cp
		}

		p++

		pp = pp.Next().(iterators.RandomAccessIterator)
		c = 2*p + 1
	}

	return last
}
