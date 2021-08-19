package constraints

import (
	"github.com/rbee3u/golib/stl/types"
)

type Writeable interface {
	Write(data types.Data)
}
