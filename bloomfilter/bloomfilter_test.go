package bloomfilter_test

import (
	"encoding/binary"
	"strings"
	"sync"
	"testing"

	"github.com/rbee3u/golib/bloomfilter"
)

func TestNewWith(t *testing.T) {
	t.Run("ArgumentNP", func(t *testing.T) {
		tests := []struct {
			n uint64
			p float64
			e string
		}{
			{n: 0, p: 0.3, e: "invalid argument: n"},
			{n: 1000000, p: -1, e: "invalid argument: p"},
			{n: 1000000, p: 0, e: "invalid argument: p"},
			{n: 1000000, p: 1, e: "invalid argument: p"},
			{n: 1000000, p: 2, e: "invalid argument: p"},
			{n: 10, p: 0.03, e: ""},
			{n: 100, p: 0.03, e: ""},
			{n: 10, p: 0.003, e: ""},
		}

		for _, tt := range tests {
			bf, err := bloomfilter.NewWithEstimate(tt.n, tt.p)
			assertNew(t, bf, err, tt.e)
		}
	})

	t.Run("ArgumentMK", func(t *testing.T) {
		tests := []struct {
			m uint64
			k uint64
			e string
		}{
			{m: 0, k: 5, e: "invalid argument: m"},
			{m: 1000000, k: 0, e: "invalid argument: k"},
			{m: 10, k: 5, e: ""},
			{m: 100, k: 5, e: ""},
			{m: 10, k: 1, e: ""},
		}

		for _, tt := range tests {
			bf, err := bloomfilter.New(tt.m, tt.k)
			assertNew(t, bf, err, tt.e)
		}
	})

	t.Run("NilHasher", func(t *testing.T) {
		bf, err := bloomfilter.New(10, 1, bloomfilter.WithHasher(nil))
		assertNew(t, bf, err, "invalid argument: nil hasher")
	})
}

func TestBloomFilter_Add(t *testing.T) {
	bf, err := bloomfilter.NewWithEstimate(1000, 0.03)
	assertNew(t, bf, err, "")

	foo, bar := []byte("foo"), []byte("bar")
	assertFalse(t, bf.Contains(foo))
	assertFalse(t, bf.Contains(bar))
	bf.Add(foo)
	assertTrue(t, bf.Contains(foo))
	assertFalse(t, bf.Contains(bar))
	bf.Add(bar)
	assertTrue(t, bf.Contains(foo))
	assertTrue(t, bf.Contains(bar))
}

func TestBloomFilter_Contains(t *testing.T) {
	bf, err := bloomfilter.NewWithEstimate(1000, 0.03)
	assertNew(t, bf, err, "")

	foo, bar := []byte("foo"), []byte("bar")
	assertFalse(t, bf.Contains(foo))
	assertFalse(t, bf.Contains(bar))
	bf.Add(foo)
	assertTrue(t, bf.Contains(foo))
	assertFalse(t, bf.Contains(bar))
	bf.Add(bar)
	assertTrue(t, bf.Contains(foo))
	assertTrue(t, bf.Contains(bar))
}

func assertNew(t *testing.T, bf *bloomfilter.BloomFilter, err error, e string) {
	if len(e) != 0 {
		if bf != nil {
			t.Errorf("expect bf(%v) to be nil", bf)
		}
		if err == nil || !strings.Contains(err.Error(), e) {
			t.Errorf(`expect err(%v) contains "%s"`, err, e)
		}
	} else {
		if bf == nil {
			t.Errorf("expect bf(%v) not to be nil", bf)
		}
		if err != nil {
			t.Errorf("expect err(%v) to be nil", err)
		}
	}
}

func assertTrue(t *testing.T, b bool) {
	if !b {
		t.Errorf("expect b(%t) to be true", b)
	}
}

func assertFalse(t *testing.T, b bool) {
	if b {
		t.Errorf("expect b(%t) to be false", b)
	}
}

func BenchmarkBloomFilter_Add(b *testing.B) {
	pool := sync.Pool{
		New: func() interface{} {
			item := make([]byte, 8)
			return &item
		},
	}

	const n = 1000000
	bf, _ := bloomfilter.NewWithEstimate(n, 0.03)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			item := pool.Get().(*[]byte)
			binary.BigEndian.PutUint64(*item, uint64(i%n))
			bf.Add(*item)
			pool.Put(item)
		}
	})
}

func BenchmarkBloomFilter_AddWithoutLock(b *testing.B) {
	item := make([]byte, 8)

	const n = 1000000
	bf, _ := bloomfilter.NewWithEstimate(n, 0.03)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		binary.BigEndian.PutUint64(item, uint64(i%n))
		bf.AddWithoutLock(item)
	}
}

func BenchmarkBloomFilter_Contains(b *testing.B) {
	pool := sync.Pool{
		New: func() interface{} {
			item := make([]byte, 8)
			return &item
		},
	}

	const n = 1000000
	bf, _ := bloomfilter.NewWithEstimate(n, 0.03)
	for i := 0; i < n; i++ {
		item := pool.Get().(*[]byte)
		binary.BigEndian.PutUint64(*item, uint64(i%n))
		bf.Add(*item)
		pool.Put(item)
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			item := pool.Get().(*[]byte)
			binary.BigEndian.PutUint64(*item, uint64(i%n))
			_ = bf.Contains(*item)
			pool.Put(item)
		}
	})
}

func BenchmarkBloomFilter_ContainsWithoutLock(b *testing.B) {
	item := make([]byte, 8)

	const n = 1000000
	bf, _ := bloomfilter.NewWithEstimate(n, 0.03)
	for i := 0; i < n; i++ {
		binary.BigEndian.PutUint64(item, uint64(i%n))
		bf.Add(item)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		binary.BigEndian.PutUint64(item, uint64(i%n))
		_ = bf.ContainsWithoutLock(item)
	}
}
