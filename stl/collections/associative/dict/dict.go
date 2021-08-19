package dict

import (
	base "github.com/rbee3u/golib/stl/collections/associative/avl"
	"github.com/rbee3u/golib/stl/types"
)

type Dict struct {
	base *base.Tree
	less types.BinaryPredicate
}

type Value struct {
	Key    types.Data
	Mapped types.Data
}

func New(keyLess types.BinaryPredicate) *Dict {
	valueLess := func(x types.Data, y types.Data) bool {
		return keyLess(x.(Value).Key, y.(Value).Key)
	}

	return &Dict{
		base: base.New(valueLess),
		less: keyLess,
	}
}

func (d *Dict) Size() types.Size {
	return d.base.Size()
}

func (d *Dict) Empty() bool {
	return d.base.Empty()
}

func (d *Dict) Begin() Iterator {
	return Iterator{base: d.base.Begin()}
}

func (d *Dict) End() Iterator {
	return Iterator{base: d.base.End()}
}

func (d *Dict) ReverseBegin() Iterator {
	return Iterator{base: d.base.ReverseBegin()}
}

func (d *Dict) ReverseEnd() Iterator {
	return Iterator{base: d.base.ReverseEnd()}
}

func (d *Dict) Count(k types.Data) types.Size {
	return d.base.CountUnique(Value{Key: k})
}

func (d *Dict) Find(k types.Data) Iterator {
	return Iterator{base: d.base.Find(Value{Key: k})}
}

func (d *Dict) Contains(k types.Data) bool {
	return d.base.Contains(Value{Key: k})
}

func (d *Dict) EqualRange(k types.Data) (Iterator, Iterator) {
	lb, ub := d.base.EqualRangeUnique(Value{Key: k})

	return Iterator{base: lb}, Iterator{base: ub}
}

func (d *Dict) LowerBound(k types.Data) Iterator {
	return Iterator{base: d.base.LowerBound(Value{Key: k})}
}

func (d *Dict) UpperBound(k types.Data) Iterator {
	return Iterator{base: d.base.UpperBound(Value{Key: k})}
}

func (d *Dict) Clear() {
	d.base.Clear()
}

func (d *Dict) Insert(k types.Data, m types.Data) (Iterator, bool) {
	it, ok := d.base.InsertUnique(Value{Key: k, Mapped: m})

	return Iterator{base: it}, ok
}

func (d *Dict) Erase(i Iterator) Iterator {
	return Iterator{base: d.base.Delete(i.base)}
}
