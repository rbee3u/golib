package constraints

import (
	"github.com/rbee3u/golib/stl/types"
)

type Readable interface {
	Read() types.Data
}
