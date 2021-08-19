package memo

import (
	_ "unsafe"
)

// nanotime is a link to runtime.nanotime.
//
//go:linkname nanotime runtime.nanotime
func nanotime() int64
