package vector_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator_Write(t *testing.T) {
	i := newIterator()
	assert.Equal(t, 0, i.Read())
	i.Write(1)
	assert.Equal(t, 1, i.Read())
}

func BenchmarkIterator_Write(b *testing.B) {
	i := newIterator()
	for n := 0; n < b.N; n++ {
		i.Write(1)
	}
}

func TestIterator_Clone(t *testing.T) {
	i := newIterator()
	j := i.Clone()
	assert.True(t, &i != &j)
	assert.Equal(t, 0, i.Read())
	assert.Equal(t, 0, j.Read())
	j.Write(1)
	assert.Equal(t, 1, i.Read())
	assert.Equal(t, 1, j.Read())
}

func BenchmarkIterator_Clone(b *testing.B) {
	i := newIterator()
	for n := 0; n < b.N; n++ {
		_ = i.Clone()
	}
}

func TestIterator_ImplClone(t *testing.T) {
	i := newIterator()
	j := i.Clone()
	assert.True(t, &i != &j)
	assert.Equal(t, 0, i.Read())
	assert.Equal(t, 0, j.Read())
	j.Write(1)
	assert.Equal(t, 1, i.Read())
	assert.Equal(t, 1, j.Read())
}

func BenchmarkIterator_ImplClone(b *testing.B) {
	i := newIterator()
	for n := 0; n < b.N; n++ {
		_ = i.Clone()
	}
}

func TestIterator_Next(t *testing.T) {
	l := newList()
	l.PushBack(1)
	l.PushBack(2)
	i := l.Begin()
	assert.Equal(t, 1, i.Read())
	i = i.Next()
	assert.Equal(t, 2, i.Read())
	i = i.Next()
	assert.Equal(t, l.End(), i)
}

func BenchmarkIterator_Next(b *testing.B) {
	i := newIterator()
	for n := 0; n < b.N; n++ {
		_ = i.Next()
	}
}

func TestIterator_ImplNext(t *testing.T) {
	l := newList()
	l.PushBack(1)
	l.PushBack(2)
	i := l.Begin()
	assert.Equal(t, 1, i.Read())
	i = i.Next()
	assert.Equal(t, 2, i.Read())
	i = i.Next()
	assert.Equal(t, l.End(), i)
}

func BenchmarkIterator_ImplNext(b *testing.B) {
	i := newIterator()
	for n := 0; n < b.N; n++ {
		_ = i.Next()
	}
}

func TestIterator_Equal(t *testing.T) {
	i := newIterator()
	assert.True(t, i.Equal(i))
	j := newIterator()
	assert.False(t, i.Equal(j))
}

func BenchmarkIterator_Equal(b *testing.B) {
	i := newIterator()
	for n := 0; n < b.N; n++ {
		_ = i.Equal(i)
	}
}

func TestIterator_ImplEqual(t *testing.T) {
	i := newIterator()
	assert.True(t, i.Equal(i))
	j := newIterator()
	assert.False(t, i.Equal(j))
}

func BenchmarkIterator_ImplEqual(b *testing.B) {
	i := newIterator()
	for n := 0; n < b.N; n++ {
		_ = i.Equal(i)
	}
}

func TestIterator_Read(t *testing.T) {
	i := newIterator()
	i.Write(1)
	assert.Equal(t, 1, i.Read())
}

func BenchmarkIterator_Read(b *testing.B) {
	i := newIterator()
	for n := 0; n < b.N; n++ {
		_ = i.Read()
	}
}

func TestIterator_Prev(t *testing.T) {
	l := newList()
	l.PushBack(1)
	l.PushBack(2)
	i := l.ReverseBegin()
	assert.Equal(t, 2, i.Read())
	i = i.Prev()
	assert.Equal(t, 1, i.Read())
	i = i.Prev()
	assert.Equal(t, l.ReverseEnd(), i)
}

func BenchmarkIterator_Prev(b *testing.B) {
	i := newIterator()
	for n := 0; n < b.N; n++ {
		_ = i.Prev()
	}
}

func TestIterator_ImplPrev(t *testing.T) {
	l := newList()
	l.PushBack(1)
	l.PushBack(2)
	i := l.ReverseBegin()
	assert.Equal(t, 2, i.Read())
	i = i.Prev()
	assert.Equal(t, 1, i.Read())
	i = i.Prev()
	assert.Equal(t, l.ReverseEnd(), i)
}

func BenchmarkIterator_ImplPrev(b *testing.B) {
	i := newIterator()
	for n := 0; n < b.N; n++ {
		_ = i.Prev()
	}
}

func TestIterator_Less(t *testing.T) {
	l := newList()
	l.PushBack(1)
	i := l.Begin()
	j := l.End()
	assert.True(t, i.Less(j))
}

func BenchmarkIterator_Less(b *testing.B) {
	l := newList()
	l.PushBack(1)
	i := l.Begin()
	j := l.End()
	for n := 0; n < b.N; n++ {
		_ = i.Less(j)
	}
}

func TestIterator_ImplLess(t *testing.T) {
	l := newList()
	l.PushBack(1)
	i := l.Begin()
	j := l.End()
	assert.True(t, i.Less(j))
}

func BenchmarkIterator_ImplLess(b *testing.B) {
	l := newList()
	l.PushBack(1)
	i := l.Begin()
	j := l.End()
	for n := 0; n < b.N; n++ {
		_ = i.Less(j)
	}
}

func TestIterator_At(t *testing.T) {
	l := newList()
	l.PushBack(1)
	i := l.Begin()
	assert.Equal(t, 1, i.At(0))
}

func BenchmarkIterator_At(b *testing.B) {
	l := newList()
	l.PushBack(1)
	i := l.Begin()
	for n := 0; n < b.N; n++ {
		_ = i.At(0)
	}
}

func TestIterator_Advance(t *testing.T) {
	l := newList()
	l.PushBack(1)
	i := l.Begin()
	j := i.Advance(1)
	assert.True(t, j.Equal(l.End()))
}

func BenchmarkIterator_Advance(b *testing.B) {
	l := newList()
	l.PushBack(1)
	i := l.Begin()
	for n := 0; n < b.N; n++ {
		_ = i.Advance(1)
	}
}

func TestIterator_ImplAdvance(t *testing.T) {
	l := newList()
	l.PushBack(1)
	i := l.Begin()
	j := i.Advance(1)
	assert.True(t, j.Equal(l.End()))
}

func BenchmarkIterator_ImplAdvance(b *testing.B) {
	l := newList()
	l.PushBack(1)
	i := l.Begin()
	for n := 0; n < b.N; n++ {
		_ = i.Advance(1)
	}
}

func TestIterator_Distance(t *testing.T) {
	l := newList()
	l.PushBack(1)
	i := l.Begin()
	j := l.End()
	assert.Equal(t, 1, i.Distance(j))
}

func BenchmarkIterator_Distance(b *testing.B) {
	l := newList()
	l.PushBack(1)
	i := l.Begin()
	j := l.End()
	for n := 0; n < b.N; n++ {
		_ = i.Distance(j)
	}
}

func TestIterator_ImplDistance(t *testing.T) {
	l := newList()
	l.PushBack(1)
	i := l.Begin()
	j := l.End()
	assert.Equal(t, 1, i.Distance(j))
}

func BenchmarkIterator_ImplDistance(b *testing.B) {
	l := newList()
	l.PushBack(1)
	i := l.Begin()
	j := l.End()
	for n := 0; n < b.N; n++ {
		_ = i.Distance(j)
	}
}
