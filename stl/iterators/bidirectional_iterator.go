package iterators

import (
	"github.com/rbee3u/golib/stl/constraints"
)

type BidirectionalIterator interface {
	ForwardIterator
	constraints.Decrementable
}
