package dict

import (
	base "github.com/rbee3u/golib/stl/collections/associative/avl"
)

type Dict[K any, V any] struct {
	base *base.Tree[Pair[K, V]]
}

type Pair[K any, V any] struct {
	Key   K
	Value V
}

func New[K any, V any](keyLess func(K, K) bool) *Dict[K, V] {
	pairLess := func(x Pair[K, V], y Pair[K, V]) bool {
		return keyLess(x.Key, y.Key)
	}

	return &Dict[K, V]{base: base.New(pairLess)}
}

func (d *Dict[K, V]) Size() int {
	return d.base.Size()
}

func (d *Dict[K, V]) Empty() bool {
	return d.base.Empty()
}

func (d *Dict[K, V]) Begin() Iterator[Pair[K, V]] {
	return Iterator[Pair[K, V]]{base: d.base.Begin()}
}

func (d *Dict[K, V]) End() Iterator[Pair[K, V]] {
	return Iterator[Pair[K, V]]{base: d.base.End()}
}

func (d *Dict[K, V]) ReverseBegin() Iterator[Pair[K, V]] {
	return Iterator[Pair[K, V]]{base: d.base.ReverseBegin()}
}

func (d *Dict[K, V]) ReverseEnd() Iterator[Pair[K, V]] {
	return Iterator[Pair[K, V]]{base: d.base.ReverseEnd()}
}

func (d *Dict[K, V]) Count(key K) int {
	return d.base.CountUnique(Pair[K, V]{Key: key})
}

func (d *Dict[K, V]) Find(key K) Iterator[Pair[K, V]] {
	return Iterator[Pair[K, V]]{base: d.base.Find(Pair[K, V]{Key: key})}
}

func (d *Dict[K, V]) Contains(key K) bool {
	return d.base.Contains(Pair[K, V]{Key: key})
}

func (d *Dict[K, V]) EqualRange(key K) (Iterator[Pair[K, V]], Iterator[Pair[K, V]]) {
	lb, ub := d.base.EqualRangeUnique(Pair[K, V]{Key: key})

	return Iterator[Pair[K, V]]{base: lb}, Iterator[Pair[K, V]]{base: ub}
}

func (d *Dict[K, V]) LowerBound(key K) Iterator[Pair[K, V]] {
	return Iterator[Pair[K, V]]{base: d.base.LowerBound(Pair[K, V]{Key: key})}
}

func (d *Dict[K, V]) UpperBound(key K) Iterator[Pair[K, V]] {
	return Iterator[Pair[K, V]]{base: d.base.UpperBound(Pair[K, V]{Key: key})}
}

func (d *Dict[K, V]) Clear() {
	d.base.Clear()
}

func (d *Dict[K, V]) Insert(key K, value V) (Iterator[Pair[K, V]], bool) {
	it, ok := d.base.InsertUnique(Pair[K, V]{Key: key, Value: value})

	return Iterator[Pair[K, V]]{base: it}, ok
}

func (d *Dict[K, V]) Erase(i Iterator[Pair[K, V]]) Iterator[Pair[K, V]] {
	return Iterator[Pair[K, V]]{base: d.base.Delete(i.base)}
}
