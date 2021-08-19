package memo_test

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/rbee3u/golib/memo"
)

func TestMemo(t *testing.T) {
	fc := newFakeClock()
	g := &generator{r: rand.New(rand.NewSource(142857677367)), mk: 100, mv: 1000000000}
	m := memo.New(memo.WithClock[int, int](fc), memo.WithLoader[int, int](nil), memo.WithExpiration[int, int](0))
	c := &competitor{clock: fc, dict: make(map[int]*entry)}

	for i := 0; i < 100000; i++ {
		fc.advance(time.Second)
		switch op := g.next().(type) {
		case opGet:
			var loader func(int) (int, error)
			if op.v != 0 || op.err != nil {
				loader = func(_ int) (int, error) {
					return op.v, op.err
				}
			}
			v1, err1 := m.Get(op.k, memo.GetWithLoader(loader), memo.GetWithExpiration[int, int](op.expiration))
			v2, err2 := c.get(op.k, loader, op.expiration)
			if v1 != v2 {
				t.Errorf("got: %v, want: %v", v1, v2)
			}
			if err1 != err2 {
				t.Errorf("got: %v, want: %v", err1, err2)
			}
		case opSet:
			m.Set(op.k, op.v, memo.SetWithExpiration[int, int](op.expiration))
			c.set(op.k, op.v, op.expiration)
		case opDel:
			m.Del(op.k)
			c.del(op.k)
		}
	}
}

func TestInvalidExpiration(t *testing.T) {
	tests := []struct {
		name string
		exec func()
	}{
		{name: "New", exec: func() {
			_ = memo.New(memo.WithExpiration[int, int](-1))
		}},
		{name: "Get", exec: func() {
			_, _ = memo.New[int, int]().Get(0, memo.GetWithExpiration[int, int](-1))
		}},
		{name: "Set", exec: func() {
			memo.New[int, int]().Set(0, 0, memo.SetWithExpiration[int, int](-1))
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err, _ := recover().(error)
				if !errors.Is(err, memo.ErrInvalidExpiration) {
					t.Errorf("got: %v, want: %v", err, memo.ErrInvalidExpiration)
				}
			}()
			tt.exec()
		})
	}
}

func BenchmarkMemo_Get(b *testing.B) {
	b.Run("FastPath", func(b *testing.B) {
		m := memo.New[string, string]()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_, _ = m.Get("k")
			}
		})
	})

	b.Run("SlowPath", func(b *testing.B) {
		loader := func(_ string) (string, error) {
			return "v", nil
		}
		m := memo.New(
			memo.WithLoader(loader),
			memo.WithExpiration[string, string](time.Nanosecond),
		)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_, _ = m.Get("k")
			}
		})
	})
}

func BenchmarkMemo_Set(b *testing.B) {
	m := memo.New[string, string]()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Set("k", "v")
		}
	})
}

func BenchmarkMemo_Del(b *testing.B) {
	m := memo.New[string, string]()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Del("k")
		}
	})
}

type fakeClock struct {
	mu       sync.Mutex
	nanotime int64
}

func newFakeClock() *fakeClock {
	return &fakeClock{}
}

func (fc *fakeClock) Now() int64 {
	fc.mu.Lock()
	nt := fc.nanotime
	fc.mu.Unlock()
	return nt
}

func (fc *fakeClock) advance(d time.Duration) {
	fc.mu.Lock()
	fc.nanotime += int64(d)
	fc.mu.Unlock()
}

type generator struct {
	r  *rand.Rand
	mk int
	mv int
}

func (g *generator) next() interface{} {
	switch g.r.Intn(9) {
	case 0, 1, 2:
		return g.opGet()
	case 3, 4, 5:
		return g.opSet()
	default:
		return g.opDel()
	}
}

type opGet struct {
	k          int
	v          int
	err        error
	expiration time.Duration
}

func (g *generator) opGet() (op opGet) {
	op.k = 1 + g.r.Intn(g.mk)
	switch g.r.Intn(3) {
	case 0:
		op.v = 1 + g.r.Intn(g.mv)
	case 1:
		op.err = fmt.Errorf("e%v", 1+g.r.Intn(g.mv))
	default:
		return
	}
	op.v = 1 + g.r.Intn(g.mv)
	if g.r.Intn(3) == 0 {
		return
	}
	op.expiration = time.Minute
	return
}

type opSet struct {
	k          int
	v          int
	expiration time.Duration
}

func (g *generator) opSet() (op opSet) {
	op.k = 1 + g.r.Intn(g.mk)
	op.v = 1 + g.r.Intn(g.mk)
	if g.r.Intn(2) == 0 {
		return
	}
	op.expiration = time.Minute
	return
}

type opDel struct {
	k int
}

func (g *generator) opDel() (op opDel) {
	op.k = 1 + g.r.Intn(g.mk)
	return
}

type competitor struct {
	clock memo.Clock
	dict  map[int]*entry
}

type entry struct {
	value    int
	err      error
	expireAt int64
}

func (c *competitor) get(k int, loader func(int) (int, error), expiration time.Duration) (int, error) {
	now := c.clock.Now()

	e := c.dict[k]
	if e != nil && (e.expireAt == 0 || e.expireAt > now) {
		return e.value, e.err
	}

	if loader == nil {
		return 0, memo.ErrNotFound
	}

	e = &entry{}
	e.value, e.err = loader(k)
	if expiration != 0 {
		e.expireAt = now + int64(expiration)
	}
	c.dict[k] = e
	return e.value, e.err
}

func (c *competitor) set(k int, v int, expiration time.Duration) {
	now := c.clock.Now()

	e := c.dict[k]
	if e == nil {
		e = &entry{}
		c.dict[k] = e
	}

	e.value, e.err = v, nil
	e.expireAt = 0
	if expiration != 0 {
		e.expireAt = now + int64(expiration)
	}
}

func (c *competitor) del(k int) {
	delete(c.dict, k)
}
