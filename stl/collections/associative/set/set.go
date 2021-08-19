package set

import (
	base "github.com/rbee3u/golib/stl/collections/associative/avl"
)

type Set[T any] struct {
	base *base.Tree[T]
}

func New[T any](less func(T, T) bool) *Set[T] {
	return &Set[T]{base: base.New(less)}
}

func (s *Set[T]) Size() int {
	return s.base.Size()
}

func (s *Set[T]) Empty() bool {
	return s.base.Empty()
}

func (s *Set[T]) Begin() Iterator[T] {
	return Iterator[T]{base: s.base.Begin()}
}

func (s *Set[T]) End() Iterator[T] {
	return Iterator[T]{base: s.base.End()}
}

func (s *Set[T]) ReverseBegin() Iterator[T] {
	return Iterator[T]{base: s.base.ReverseBegin()}
}

func (s *Set[T]) ReverseEnd() Iterator[T] {
	return Iterator[T]{base: s.base.ReverseEnd()}
}

func (s *Set[T]) Count(data T) int {
	return s.base.CountUnique(data)
}

func (s *Set[T]) Find(data T) Iterator[T] {
	return Iterator[T]{base: s.base.Find(data)}
}

func (s *Set[T]) Contains(data T) bool {
	return s.base.Contains(data)
}

func (s *Set[T]) EqualRange(data T) (Iterator[T], Iterator[T]) {
	lb, ub := s.base.EqualRangeUnique(data)

	return Iterator[T]{base: lb}, Iterator[T]{base: ub}
}

func (s *Set[T]) LowerBound(data T) Iterator[T] {
	return Iterator[T]{base: s.base.LowerBound(data)}
}

func (s *Set[T]) UpperBound(data T) Iterator[T] {
	return Iterator[T]{base: s.base.UpperBound(data)}
}

func (s *Set[T]) Clear() {
	s.base.Clear()
}

func (s *Set[T]) Insert(data T) (Iterator[T], bool) {
	it, ok := s.base.InsertUnique(data)

	return Iterator[T]{base: it}, ok
}

func (s *Set[T]) Erase(i Iterator[T]) Iterator[T] {
	return Iterator[T]{base: s.base.Delete(i.base)}
}
