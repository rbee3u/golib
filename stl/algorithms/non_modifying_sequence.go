package algorithms

import (
	"github.com/rbee3u/golib/stl/iterators"
	"github.com/rbee3u/golib/stl/types"
)

func AllOf(first, last iterators.InputIterator, pred types.UnaryPredicate) bool {
	for !first.Equal(last) {
		if !pred(first.Read()) {
			return false
		}

		first = first.Next().(iterators.InputIterator)
	}

	return true
}

func AnyOf(first, last iterators.InputIterator, pred types.UnaryPredicate) bool {
	for !first.Equal(last) {
		if pred(first.Read()) {
			return true
		}

		first = first.Next().(iterators.InputIterator)
	}

	return false
}

func NoneOf(first, last iterators.InputIterator, pred types.UnaryPredicate) bool {
	for !first.Equal(last) {
		if pred(first.Read()) {
			return false
		}

		first = first.Next().(iterators.InputIterator)
	}

	return true
}

func ForEach(first, last iterators.InputIterator, fn types.UnaryPredicate) {
	for !first.Equal(last) {
		fn(first.Read())
		first = first.Next().(iterators.InputIterator)
	}
}

func Find(first, last iterators.InputIterator, val types.Data) iterators.InputIterator {
	for !first.Equal(last) {
		if first.Read() == val {
			return first
		}

		first = first.Next().(iterators.InputIterator)
	}

	return last
}

func FindIf(first, last iterators.InputIterator, pred types.UnaryPredicate) iterators.InputIterator {
	for !first.Equal(last) {
		if pred(first.Read()) {
			return first
		}

		first = first.Next().(iterators.InputIterator)
	}

	return last
}

func FindIfNot(first, last iterators.InputIterator, pred types.UnaryPredicate) iterators.InputIterator {
	for !first.Equal(last) {
		if !pred(first.Read()) {
			return first
		}

		first = first.Next().(iterators.InputIterator)
	}

	return last
}

func Count(first, last iterators.InputIterator, val types.Data) (count types.Size) {
	for !first.Equal(last) {
		if first.Read() == val {
			count++
		}

		first = first.Next().(iterators.InputIterator)
	}

	return count
}

func CountIf(first, last iterators.InputIterator, pred types.UnaryPredicate) (count types.Size) {
	for !first.Equal(last) {
		if pred(first.Read()) {
			count++
		}

		first = first.Next().(iterators.InputIterator)
	}

	return count
}
