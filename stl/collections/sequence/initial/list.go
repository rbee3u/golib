package initial

import (
	"github.com/rbee3u/golib/stl/types"
)

type List struct {
	slice []types.Data
}

func NewList(datas ...types.Data) *List {
	return &List{slice: datas}
}

func (l *List) Size() types.Size {
	return len(l.slice)
}

func (l *List) Empty() bool {
	return l.Size() == 0
}

func (l *List) Begin() Iterator {
	return Iterator{l: l, n: 0}
}

func (l *List) End() Iterator {
	return Iterator{l: l, n: len(l.slice)}
}
