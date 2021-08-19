package vector

import (
	"github.com/rbee3u/golib/stl/types"
)

type List struct {
	slice []types.Data
}

func NewList() *List {
	return &List{}
}

func (l *List) Size() types.Size {
	return len(l.slice)
}

func (l *List) Empty() bool {
	return l.Size() == 0
}

func (l *List) Get(i types.Size) types.Data {
	return l.slice[i]
}

func (l *List) Set(n types.Size, data types.Data) {
	l.slice[n] = data
}

func (l *List) PushBack(data types.Data) {
	l.slice = append(l.slice, data)
}

func (l *List) Back() types.Data {
	return l.slice[len(l.slice)-1]
}

func (l *List) PopBack() {
	l.slice = l.slice[:len(l.slice)-1]
}

func (l *List) Clear() {
	l.slice = nil
}

func (l *List) Begin() Iterator {
	return Iterator{l: l, n: 0}
}

func (l *List) End() Iterator {
	return Iterator{l: l, n: len(l.slice)}
}

func (l *List) ReverseBegin() Iterator {
	return Iterator{l: l, n: len(l.slice) - 1}
}

func (l *List) ReverseEnd() Iterator {
	return Iterator{l: l, n: -1}
}
