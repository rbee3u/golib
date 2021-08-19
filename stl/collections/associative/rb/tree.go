package rb

type Tree[T any] struct {
	sentinel node[T]
	start    *node[T]
	size     int
	less     func(T, T) bool
}

const (
	red   = 0
	black = 1
)

type node[T any] struct {
	parent *node[T]
	left   *node[T]
	right  *node[T]
	extra  int8
	data   T
}

func New[T any](less func(T, T) bool) *Tree[T] {
	t := &Tree[T]{less: less}
	t.start = &t.sentinel

	return t
}

func (t *Tree[T]) Size() int {
	return t.size
}

func (t *Tree[T]) Empty() bool {
	return t.Size() == 0
}

func (t *Tree[T]) Begin() Iterator[T] {
	return Iterator[T]{n: t.start}
}

func (t *Tree[T]) End() Iterator[T] {
	return Iterator[T]{n: &t.sentinel}
}

func (t *Tree[T]) ReverseBegin() Iterator[T] {
	return Iterator[T]{n: &t.sentinel}
}

func (t *Tree[T]) ReverseEnd() Iterator[T] {
	return Iterator[T]{n: t.start}
}

func (t *Tree[T]) CountUnique(data T) int {
	x := t.LowerBound(data)
	if x != t.End() && !t.less(data, x.Read()) {
		return 1
	}

	return 0
}

func (t *Tree[T]) CountMulti(data T) int {
	var c int

	for x, y := t.LowerBound(data), t.UpperBound(data); x != y; x = x.Next() {
		c++
	}

	return c
}

func (t *Tree[T]) Find(data T) Iterator[T] {
	x := t.LowerBound(data)
	if x != t.End() && !t.less(data, x.Read()) {
		return x
	}

	return t.End()
}

func (t *Tree[T]) Contains(data T) bool {
	x := t.LowerBound(data)
	if x != t.End() && !t.less(data, x.Read()) {
		return true
	}

	return false
}

func (t *Tree[T]) EqualRangeUnique(data T) (Iterator[T], Iterator[T]) {
	x := t.LowerBound(data)
	if x != t.End() && !t.less(data, x.Read()) {
		return x, x.Next()
	}

	return x, x
}

func (t *Tree[T]) EqualRangeMulti(data T) (Iterator[T], Iterator[T]) {
	return t.LowerBound(data), t.UpperBound(data)
}

func (t *Tree[T]) LowerBound(data T) Iterator[T] {
	return Iterator[T]{n: t.lowerBound(data)}
}

func (t *Tree[T]) lowerBound(data T) *node[T] {
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

func (t *Tree[T]) UpperBound(data T) Iterator[T] {
	return Iterator[T]{n: t.upperBound(data)}
}

func (t *Tree[T]) upperBound(data T) *node[T] {
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

func (t *Tree[T]) Clear() {
	t.sentinel.left = nil
	t.start = &t.sentinel
	t.size = 0
}

func (t *Tree[T]) InsertUnique(data T) (Iterator[T], bool) {
	lb := t.LowerBound(data)
	if lb != t.End() && !t.less(data, lb.Read()) {
		return t.End(), false
	}

	z := &node[T]{data: data}
	t.insert(z)

	return Iterator[T]{n: z}, true
}

func (t *Tree[T]) InsertMulti(data T) Iterator[T] {
	z := &node[T]{data: data}
	t.insert(z)

	return Iterator[T]{n: z}
}

func (t *Tree[T]) Delete(i Iterator[T]) Iterator[T] {
	r := i.Next()
	t.delete(i.n)

	return r
}

func transplant[T any](u *node[T], v *node[T]) {
	if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}

	if v != nil {
		v.parent = u.parent
	}
}

func minimum[T any](x *node[T]) *node[T] {
	for x.left != nil {
		x = x.left
	}

	return x
}

func maximum[T any](x *node[T]) *node[T] {
	for x.right != nil {
		x = x.right
	}

	return x
}

func successor[T any](x *node[T]) *node[T] {
	if x.right != nil {
		return minimum(x.right)
	}

	for x == x.parent.right {
		x = x.parent
	}

	return x.parent
}

func predecessor[T any](x *node[T]) *node[T] {
	if x.left != nil {
		return maximum(x.left)
	}

	for x == x.parent.left {
		x = x.parent
	}

	return x.parent
}

func (t *Tree[T]) insert(z *node[T]) {
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

func (t *Tree[T]) balanceAfterInsert(x *node[T], z *node[T]) {
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

func (t *Tree[T]) delete(z *node[T]) {
	if t.start == z {
		t.start = successor(z)
	}

	x, deletedColor := z.parent, z.extra

	var n *node[T]

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

func (t *Tree[T]) balanceAfterDelete(x *node[T], n *node[T]) {
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

func rotateLeft[T any](x *node[T]) {
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

func rotateRight[T any](x *node[T]) {
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

func isRed[T any](x *node[T]) bool {
	return x != nil && x.extra == red
}

func isBlack[T any](x *node[T]) bool {
	return x == nil || x.extra == black
}
