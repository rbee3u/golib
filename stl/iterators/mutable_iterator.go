package iterators

import (
	"github.com/rbee3u/golib/stl/constraints"
)

type MutableIterator interface {
	constraints.Writeable
}
