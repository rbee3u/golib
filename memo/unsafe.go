package memo

import (
	"unsafe"
)

// nanotime is a link to runtime.nanotime.
//go:linkname nanotime runtime.nanotime
func nanotime() int64

// noescape is copied from runtime.noescape.
//go:nosplit
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)

	return unsafe.Pointer(x ^ 0) //nolint:staticcheck
}
