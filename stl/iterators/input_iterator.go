package iterators

import (
	"github.com/rbee3u/golib/stl/constraints"
)

type InputIterator interface {
	Iterator
	constraints.EqualityComparable
	constraints.Readable
}
