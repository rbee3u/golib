package rb

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
	red   = 0
	black = 1
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

func (t *Tree) Delete(i Iterator) Iterator {
	r := i.ImplNext()
	t.delete(i.n)

	return r
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

func (t *Tree) insert(z *node) {
	z.parent = nil
	z.left, z.right, z.extra = nil, nil, red
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

	t.balanceAfterInsert(x, z)
	t.size++
}

func (t *Tree) balanceAfterInsert(x *node, z *node) {
	for ; x != &t.sentinel && x.extra == red; x = z.parent {
		if x == x.parent.left {
			y := x.parent.right
			if isRed(y) {
				z = z.parent
				z.extra = black
				z = z.parent
				z.extra = red
				y.extra = black
			} else {
				if z == x.right {
					z = x
					rotateLeft(z)
				}
				z = z.parent
				z.extra = black
				z = z.parent
				z.extra = red
				rotateRight(z)
			}
		} else {
			y := x.parent.left
			if isRed(y) {
				z = z.parent
				z.extra = black
				z = z.parent
				z.extra = red
				y.extra = black
			} else {
				if z == x.left {
					z = x
					rotateRight(z)
				}
				z = z.parent
				z.extra = black
				z = z.parent
				z.extra = red
				rotateLeft(z)
			}
		}
	}

	t.sentinel.left.extra = black
}

func (t *Tree) delete(z *node) {
	if t.start == z {
		t.start = successor(z)
	}

	x, deletedColor := z.parent, z.extra

	var n *node

	switch {
	case z.left == nil:
		n = z.right
		transplant(z, n)
	case z.right == nil:
		n = z.left
		transplant(z, n)
	default:
		y := minimum(z.right)
		x, deletedColor = y, y.extra
		n = y.right

		if y.parent != z {
			x = y.parent
			transplant(y, n)
			y.right = z.right
			y.right.parent = y
		}

		transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.extra = z.extra
	}

	if deletedColor == black {
		t.balanceAfterDelete(x, n)
	}

	t.size--
}

func (t *Tree) balanceAfterDelete(x *node, n *node) {
	for ; x != &t.sentinel && isBlack(n); x = n.parent {
		if n == x.left {
			z := x.right
			if isRed(z) {
				z.extra = black
				x.extra = red
				rotateLeft(x)
				z = x.right
			}

			if isBlack(z.left) && isBlack(z.right) {
				z.extra = red
				n = x
			} else {
				if isBlack(z.right) {
					z.left.extra = black
					z.extra = red
					rotateRight(z)
					z = x.right
				}
				z.extra = x.extra
				x.extra = black
				z.right.extra = black
				rotateLeft(x)
				n = t.sentinel.left
			}
		} else {
			z := x.left
			if isRed(z) {
				z.extra = black
				x.extra = red
				rotateRight(x)
				z = x.left
			}
			if isBlack(z.right) && isBlack(z.left) {
				z.extra = red
				n = x
			} else {
				if isBlack(z.left) {
					z.right.extra = black
					z.extra = red
					rotateLeft(z)
					z = x.left
				}
				z.extra = x.extra
				x.extra = black
				z.left.extra = black
				rotateRight(x)
				n = t.sentinel.left
			}
		}
	}

	if isRed(n) {
		n.extra = black
	}
}

func rotateLeft(x *node) {
	y := x.right
	x.right = y.left

	if x.right != nil {
		x.right.parent = x
	}

	y.parent = x.parent

	if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

func rotateRight(x *node) {
	y := x.left
	x.left = y.right

	if x.left != nil {
		x.left.parent = x
	}

	y.parent = x.parent

	if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}

	y.right = x
	x.parent = y
}

func isRed(x *node) bool {
	return x != nil && x.extra == red
}

func isBlack(x *node) bool {
	return x == nil || x.extra == black
}
