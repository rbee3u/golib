package multiset

import (
	base "github.com/rbee3u/golib/stl/collections/associative/avl"
	"github.com/rbee3u/golib/stl/types"
)

type Set struct {
	base *base.Tree
	less types.BinaryPredicate
}

func New(keyLess types.BinaryPredicate) *Set {
	valueLess := keyLess

	return &Set{
		base: base.New(valueLess),
		less: keyLess,
	}
}

func (s *Set) Size() types.Size {
	return s.base.Size()
}

func (s *Set) Empty() bool {
	return s.base.Empty()
}

func (s *Set) Begin() Iterator {
	return Iterator{base: s.base.Begin()}
}

func (s *Set) End() Iterator {
	return Iterator{base: s.base.End()}
}

func (s *Set) ReverseBegin() Iterator {
	return Iterator{base: s.base.ReverseBegin()}
}

func (s *Set) ReverseEnd() Iterator {
	return Iterator{base: s.base.ReverseEnd()}
}

func (s *Set) Count(k types.Data) types.Size {
	return s.base.CountMulti(k)
}

func (s *Set) Find(k types.Data) Iterator {
	return Iterator{base: s.base.Find(k)}
}

func (s *Set) Contains(k types.Data) bool {
	return s.base.Contains(k)
}

func (s *Set) EqualRange(k types.Data) (Iterator, Iterator) {
	lb, ub := s.base.EqualRangeMulti(k)

	return Iterator{base: lb}, Iterator{base: ub}
}

func (s *Set) LowerBound(k types.Data) Iterator {
	return Iterator{base: s.base.LowerBound(k)}
}

func (s *Set) UpperBound(k types.Data) Iterator {
	return Iterator{base: s.base.UpperBound(k)}
}

func (s *Set) Clear() {
	s.base.Clear()
}

func (s *Set) Insert(k types.Data, m types.Data) Iterator {
	it := s.base.InsertMulti(k)

	return Iterator{base: it}
}

func (s *Set) Erase(i Iterator) Iterator {
	return Iterator{base: s.base.Delete(i.base)}
}
