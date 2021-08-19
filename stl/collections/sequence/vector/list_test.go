package vector_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	l := newList()
	assert.NotNil(t, l)
}

func BenchmarkNewList(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = newList()
	}
}

func TestList_Size(t *testing.T) {
	l := newList()
	assert.Equal(t, 0, l.Size())
	l.PushBack(1)
	assert.Equal(t, 1, l.Size())
}

func BenchmarkList_Size(b *testing.B) {
	l := newList()
	for n := 0; n < b.N; n++ {
		_ = l.Size()
	}
}

func TestList_Empty(t *testing.T) {
	l := newList()
	assert.True(t, l.Empty())
	l.PushBack(1)
	assert.False(t, l.Empty())
}

func BenchmarkList_Empty(b *testing.B) {
	l := newList()
	for n := 0; n < b.N; n++ {
		_ = l.Empty()
	}
}

func TestList_Get(t *testing.T) {
	l := newList()
	l.PushBack(1)
	assert.Equal(t, 1, l.Get(0))
}

func BenchmarkList_Get(b *testing.B) {
	l := newList()
	l.PushBack(1)
	for n := 0; n < b.N; n++ {
		_ = l.Get(0).(int)
	}
}

func TestList_Set(t *testing.T) {
	l := newList()
	l.PushBack(1)
	assert.Equal(t, 1, l.Get(0))
	l.Set(0, 2)
	assert.Equal(t, 2, l.Get(0))
}

func BenchmarkList_Set(b *testing.B) {
	l := newList()
	l.PushBack(1)
	for n := 0; n < b.N; n++ {
		l.Set(0, 1)
	}
}

func TestList_PushBack(t *testing.T) {
	l := newList()
	l.PushBack(1)
	assert.Equal(t, 1, l.Back())
}

func BenchmarkList_PushBack(b *testing.B) {
	l := newList()
	for n := 0; n < b.N; n++ {
		if l.Size() == 1000000 {
			l.Clear()
		}
		l.PushBack(1)
	}
}

func TestList_Back(t *testing.T) {
	l := newList()
	l.PushBack(1)
	assert.Equal(t, 1, l.Back())
}

func BenchmarkList_Back(b *testing.B) {
	l := newList()
	l.PushBack(1)
	for n := 0; n < b.N; n++ {
		_ = l.Back()
	}
}

func TestList_PopBack(t *testing.T) {
	l := newList()
	l.PushBack(1)
	assert.False(t, l.Empty())
	l.PopBack()
	assert.True(t, l.Empty())
}

func BenchmarkList_PopBack(b *testing.B) {
	l := newList()
	for n := 0; n < b.N; n++ {
		if l.Size() == 0 {
			b.StopTimer()
			for k := 1; k < 1000000; k++ {
				l.PushBack(1)
			}
			b.StartTimer()
		}
		l.PopBack()
	}
}

func TestList_Clear(t *testing.T) {
	l := newList()
	l.PushBack(1)
	assert.False(t, l.Empty())
	l.Clear()
	assert.True(t, l.Empty())
}

func BenchmarkList_Clear(b *testing.B) {
	l := newList()
	for n := 0; n < b.N; n++ {
		l.Clear()
	}
}

func TestList_Begin(t *testing.T) {
	l := newList()
	l.PushBack(1)
	assert.Equal(t, 1, l.Begin().Read())
}

func BenchmarkList_Begin(b *testing.B) {
	l := newList()
	for n := 0; n < b.N; n++ {
		_ = l.Begin()
	}
}

func TestList_End(t *testing.T) {
	l := newList()
	l.PushBack(1)
	end := l.Begin().ImplNext()
	assert.True(t, l.End().Equal(end))
}

func BenchmarkList_End(b *testing.B) {
	l := newList()
	for n := 0; n < b.N; n++ {
		_ = l.End()
	}
}

func TestList_ReverseBegin(t *testing.T) {
	l := newList()
	l.PushBack(1)
	assert.Equal(t, 1, l.ReverseBegin().Read())
}

func BenchmarkList_ReverseBegin(b *testing.B) {
	l := newList()
	for n := 0; n < b.N; n++ {
		_ = l.ReverseBegin()
	}
}

func TestList_ReverseEnd(t *testing.T) {
	l := newList()
	l.PushBack(1)
	end := l.ReverseBegin().ImplPrev()
	assert.True(t, l.ReverseEnd().Equal(end))
}

func BenchmarkList_ReverseEnd(b *testing.B) {
	l := newList()
	for n := 0; n < b.N; n++ {
		_ = l.ReverseEnd()
	}
}
