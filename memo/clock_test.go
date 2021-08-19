package memo_test

import (
	"testing"
	"time"

	"github.com/rbee3u/golib/memo"
)

func TestClock(t *testing.T) {
	tests := []struct {
		name  string
		delay time.Duration
	}{
		{name: "Monotonic", delay: 0 * time.Millisecond},
		{name: "Precise", delay: 100 * time.Millisecond},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lowerBound := int64(tt.delay)
			realClock := memo.NewRealClock()
			start := realClock.Now()
			time.Sleep(tt.delay)
			diff := realClock.Now() - start
			if diff < lowerBound {
				t.Errorf("want: diff(%v) >= lowerBound(%v)", diff, lowerBound)
			}
		})
	}
}

func BenchmarkClock(b *testing.B) {
	b.Run("TimeClock", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = time.Now().UnixNano()
			}
		})
	})

	b.Run("RealClock", func(b *testing.B) {
		realClock := memo.NewRealClock()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = realClock.Now()
			}
		})
	})
}
