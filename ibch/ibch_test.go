package ibch_test

import (
	"testing"

	"github.com/rbee3u/golib/ibch"
)

func BenchmarkIBCh(b *testing.B) {
	ibCh := ibch.New()
	defer ibCh.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ibCh.SendCh() <- 42
			_ = <-ibCh.RecvCh()
		}
	})
}
