package algorithms

import (
	"github.com/rbee3u/golib/stl/constraints"
)

func AllOf[S constraints.InputIterator[S, T], T any](
	first S, last S, pred func(T) bool,
) bool {
	for ; !first.Equal(last); first = first.Next() {
		if !pred(first.Read()) {
			return false
		}
	}

	return true
}

func AnyOf[S constraints.InputIterator[S, T], T any](
	first S, last S, pred func(T) bool,
) bool {
	for ; !first.Equal(last); first = first.Next() {
		if pred(first.Read()) {
			return true
		}
	}

	return false
}

func ForEach[S constraints.InputIterator[S, T], T any](
	first S, last S, fn func(T),
) {
	for ; !first.Equal(last); first = first.Next() {
		fn(first.Read())
	}
}

func Find[S constraints.InputIterator[S, T], T constraints.EqualityComparable[T]](
	first S, last S, target T,
) S {
	return FindIf(first, last, func(x T) bool { return x.Equal(target) })
}

func FindIf[S constraints.InputIterator[S, T], T any](
	first S, last S, pred func(T) bool,
) S {
	for ; !first.Equal(last); first = first.Next() {
		if pred(first.Read()) {
			return first
		}
	}

	return last
}

func Count[S constraints.InputIterator[S, T], T constraints.EqualityComparable[T]](
	first S, last S, target T,
) int {
	return CountIf(first, last, func(x T) bool { return x.Equal(target) })
}

func CountIf[S constraints.InputIterator[S, T], T any](
	first S, last S, pred func(T) bool,
) int {
	var count int

	for ; !first.Equal(last); first = first.Next() {
		if pred(first.Read()) {
			count++
		}
	}

	return count
}
