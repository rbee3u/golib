package avl

import (
	"github.com/rbee3u/golib/stl/types"
)

type Tree struct {
	sentinel node
	start    *node
	size     types.Size
	less     types.BinaryPredicate
}

const (
	leftHeavy  = -1
	balanced   = 0
	rightHeavy = +1
)

type node struct {
	parent *node
	left   *node
	right  *node
	extra  int8
	data   types.Data
}

func New(less types.BinaryPredicate) *Tree {
	t := &Tree{less: less}
	t.start = &t.sentinel

	return t
}

func (t *Tree) Size() types.Size {
	return t.size
}

func (t *Tree) Empty() bool {
	return t.Size() == 0
}

func (t *Tree) Begin() Iterator {
	return Iterator{n: t.start}
}

func (t *Tree) End() Iterator {
	return Iterator{n: &t.sentinel}
}

func (t *Tree) ReverseBegin() Iterator {
	return Iterator{n: &t.sentinel}
}

func (t *Tree) ReverseEnd() Iterator {
	return Iterator{n: t.start}
}

func (t *Tree) CountUnique(data types.Data) types.Size {
	x := t.LowerBound(data)
	if x != t.End() && !t.less(data, x.Read()) {
		return 1
	}

	return 0
}

func (t *Tree) CountMulti(data types.Data) (c types.Size) {
	x, y := t.LowerBound(data), t.UpperBound(data)
	for x != y {
		c++

		x = x.ImplNext()
	}

	return c
}

func (t *Tree) Find(data types.Data) Iterator {
	x := t.LowerBound(data)
	if x != t.End() && !t.less(data, x.Read()) {
		return x
	}

	return t.End()
}

func (t *Tree) Contains(data types.Data) bool {
	x := t.LowerBound(data)
	if x != t.End() && !t.less(data, x.Read()) {
		return true
	}

	return false
}

func (t *Tree) EqualRangeUnique(data types.Data) (Iterator, Iterator) {
	x := t.LowerBound(data)
	if x != t.End() && !t.less(data, x.Read()) {
		return x, x.ImplNext()
	}

	return x, x
}

func (t *Tree) EqualRangeMulti(data types.Data) (Iterator, Iterator) {
	return t.LowerBound(data), t.UpperBound(data)
}

func (t *Tree) LowerBound(data types.Data) Iterator {
	return Iterator{n: t.lowerBound(data)}
}

func (t *Tree) lowerBound(data types.Data) *node {
	x := &t.sentinel

	for y := x.left; y != nil; {
		if !t.less(y.data, data) {
			x = y
			y = y.left
		} else {
			y = y.right
		}
	}

	return x
}

func (t *Tree) UpperBound(data types.Data) Iterator {
	return Iterator{n: t.upperBound(data)}
}

func (t *Tree) upperBound(data types.Data) *node {
	x := &t.sentinel

	for y := x.left; y != nil; {
		if t.less(data, y.data) {
			x = y
			y = y.left
		} else {
			y = y.right
		}
	}

	return x
}

func (t *Tree) Clear() {
	t.sentinel.left = nil
	t.start = &t.sentinel
	t.size = 0
}

func (t *Tree) InsertUnique(v types.Data) (Iterator, bool) {
	lb := t.LowerBound(v)
	if lb != t.End() && !t.less(v, lb.Read()) {
		return t.End(), false
	}

	z := &node{data: v}
	t.insert(z)

	return Iterator{n: z}, true
}

func (t *Tree) InsertMulti(v types.Data) Iterator {
	z := &node{data: v}
	t.insert(z)

	return Iterator{n: z}
}

func (t *Tree) insert(z *node) {
	z.extra = balanced
	z.parent, z.left, z.right = nil, nil, nil
	x, childIsLeft := &t.sentinel, true

	for y := x.left; y != nil; {
		x, childIsLeft = y, t.less(z.data, y.data)
		if childIsLeft {
			y = y.left
		} else {
			y = y.right
		}
	}

	z.parent = x

	if childIsLeft {
		x.left = z
	} else {
		x.right = z
	}

	if t.start.left != nil {
		t.start = t.start.left
	}

	t.balanceAfterInsert(x, childIsLeft)
	t.size++
}

func (t *Tree) balanceAfterInsert(x *node, childIsLeft bool) {
	for ; x != &t.sentinel; x = x.parent {
		if !childIsLeft {
			switch x.extra {
			case leftHeavy:
				x.extra = balanced

				return
			case rightHeavy:
				if x.right.extra == leftHeavy {
					rotateRightLeft(x)
				} else {
					rotateLeft(x)
				}

				return
			default:
				x.extra = rightHeavy
			}
		} else {
			switch x.extra {
			case rightHeavy:
				x.extra = balanced

				return
			case leftHeavy:
				if x.left.extra == rightHeavy {
					rotateLeftRight(x)
				} else {
					rotateRight(x)
				}

				return
			default:
				x.extra = leftHeavy
			}
		}

		childIsLeft = x == x.parent.left
	}
}

func (t *Tree) Delete(i Iterator) Iterator {
	r := i.ImplNext()
	t.delete(i.n)

	return r
}

func (t *Tree) delete(z *node) {
	if t.start == z {
		t.start = successor(z)
	}

	x, childIsLeft := z.parent, z == z.parent.left

	switch {
	case z.left == nil:
		transplant(z, z.right)
	case z.right == nil:
		transplant(z, z.left)
	default:
		if z.extra == rightHeavy {
			y := minimum(z.right)
			x, childIsLeft = y, y == y.parent.left

			if y.parent != z {
				x = y.parent
				transplant(y, y.right)
				y.right = z.right
				y.right.parent = y
			}

			transplant(z, y)
			y.left = z.left
			y.left.parent = y
			y.extra = z.extra
		} else {
			y := maximum(z.left)
			x, childIsLeft = y, y == y.parent.left

			if y.parent != z {
				x = y.parent
				transplant(y, y.left)
				y.left = z.left
				y.left.parent = y
			}

			transplant(z, y)
			y.right = z.right
			y.right.parent = y
			y.extra = z.extra
		}
	}

	t.balanceAfterDelete(x, childIsLeft)
	t.size--
}

func (t *Tree) balanceAfterDelete(x *node, childIsLeft bool) {
	for ; x != &t.sentinel; x = x.parent {
		if childIsLeft {
			switch x.extra {
			case balanced:
				x.extra = rightHeavy

				return
			case rightHeavy:
				b := x.right.extra
				if b == leftHeavy {
					rotateRightLeft(x)
				} else {
					rotateLeft(x)
				}

				if b == balanced {
					return
				}

				x = x.parent
			default:
				x.extra = balanced
			}
		} else {
			switch x.extra {
			case balanced:
				x.extra = leftHeavy

				return
			case leftHeavy:
				b := x.left.extra
				if b == rightHeavy {
					rotateLeftRight(x)
				} else {
					rotateRight(x)
				}
				if b == balanced {
					return
				}
				x = x.parent
			default:
				x.extra = balanced
			}
		}

		childIsLeft = x == x.parent.left
	}
}

func rotateLeft(x *node) {
	z := x.right
	x.right = z.left

	if z.left != nil {
		z.left.parent = x
	}

	z.parent = x.parent

	if x == x.parent.left {
		x.parent.left = z
	} else {
		x.parent.right = z
	}

	z.left = x
	x.parent = z

	if z.extra == balanced {
		x.extra, z.extra = rightHeavy, leftHeavy
	} else {
		x.extra, z.extra = balanced, balanced
	}
}

func rotateRight(x *node) {
	z := x.left
	x.left = z.right

	if z.right != nil {
		z.right.parent = x
	}

	z.parent = x.parent

	if x == x.parent.right {
		x.parent.right = z
	} else {
		x.parent.left = z
	}

	z.right = x
	x.parent = z

	if z.extra == balanced {
		x.extra, z.extra = leftHeavy, rightHeavy
	} else {
		x.extra, z.extra = balanced, balanced
	}
}

func rotateRightLeft(x *node) {
	z := x.right
	y := z.left
	z.left = y.right

	if y.right != nil {
		y.right.parent = z
	}

	y.right = z
	z.parent = y
	x.right = y.left

	if y.left != nil {
		y.left.parent = x
	}

	y.parent = x.parent

	if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y

	switch y.extra {
	case rightHeavy:
		x.extra, y.extra, z.extra = leftHeavy, balanced, balanced
	case leftHeavy:
		x.extra, y.extra, z.extra = balanced, balanced, rightHeavy
	default:
		x.extra, z.extra = balanced, balanced
	}
}

func rotateLeftRight(x *node) {
	z := x.left
	y := z.right
	z.right = y.left

	if y.left != nil {
		y.left.parent = z
	}

	y.left = z
	z.parent = y
	x.left = y.right

	if y.right != nil {
		y.right.parent = x
	}

	y.parent = x.parent

	if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}

	y.right = x
	x.parent = y

	switch y.extra {
	case leftHeavy:
		x.extra, y.extra, z.extra = rightHeavy, balanced, balanced
	case rightHeavy:
		x.extra, y.extra, z.extra = balanced, balanced, leftHeavy
	default:
		x.extra, z.extra = balanced, balanced
	}
}

func transplant(u *node, v *node) {
	if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}

	if v != nil {
		v.parent = u.parent
	}
}

func minimum(x *node) *node {
	for x.left != nil {
		x = x.left
	}

	return x
}

func maximum(x *node) *node {
	for x.right != nil {
		x = x.right
	}

	return x
}

func successor(x *node) *node {
	if x.right != nil {
		return minimum(x.right)
	}

	for x == x.parent.right {
		x = x.parent
	}

	return x.parent
}

func predecessor(x *node) *node {
	if x.left != nil {
		return maximum(x.left)
	}

	for x == x.parent.left {
		x = x.parent
	}

	return x.parent
}
