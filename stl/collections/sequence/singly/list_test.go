package singly_test

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
	l.PushFront(1)
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
	l.PushFront(1)
	assert.False(t, l.Empty())
}

func BenchmarkList_Empty(b *testing.B) {
	l := newList()
	for n := 0; n < b.N; n++ {
		_ = l.Empty()
	}
}

func TestList_PushFront(t *testing.T) {
	l := newList()
	l.PushFront(1)
	assert.Equal(t, 1, l.Front())
}

func BenchmarkList_PushFront(b *testing.B) {
	l := newList()
	for n := 0; n < b.N; n++ {
		if l.Size() == 1000000 {
			l.Clear()
		}
		l.PushFront(1)
	}
}

func TestList_Front(t *testing.T) {
	l := newList()
	l.PushFront(1)
	assert.Equal(t, 1, l.Front())
}

func BenchmarkList_Front(b *testing.B) {
	l := newList()
	l.PushFront(1)
	for n := 0; n < b.N; n++ {
		_ = l.Front()
	}
}

func TestList_PopFront(t *testing.T) {
	l := newList()
	l.PushFront(1)
	assert.False(t, l.Empty())
	l.PopFront()
	assert.True(t, l.Empty())
}

func BenchmarkList_PopFront(b *testing.B) {
	l := newList()
	for n := 0; n < b.N; n++ {
		if l.Size() == 0 {
			b.StopTimer()
			for k := 1; k < 1000000; k++ {
				l.PushFront(1)
			}
			b.StartTimer()
		}
		l.PopFront()
	}
}

func TestList_Clear(t *testing.T) {
	l := newList()
	l.PushFront(1)
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
	l.PushFront(1)
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
	l.PushFront(1)
	end := l.Begin().ImplNext()
	assert.True(t, l.End().Equal(end))
}

func BenchmarkList_End(b *testing.B) {
	l := newList()
	for n := 0; n < b.N; n++ {
		_ = l.End()
	}
}

func TestList_InsertAfter(t *testing.T) {
	l := newList()
	l.PushFront(1)
	i := l.Begin()
	j := l.InsertAfter(i, 2)
	k := l.InsertAfter(j, 3)
	assert.Equal(t, 1, i.Read())
	assert.Equal(t, 2, j.Read())
	assert.Equal(t, 3, k.Read())
}

func BenchmarkList_InsertAfter(b *testing.B) {
	l := newList()
	l.PushFront(1)
	i := l.Begin()
	for n := 0; n < b.N; n++ {
		if l.Size() == 1000000 {
			l.Clear()
			l.PushFront(1)
			i = l.Begin()
		}
		_ = l.InsertAfter(i, 1)
	}
}

func TestList_EraseAfter(t *testing.T) {
	l := newList()
	l.PushFront(3)
	l.PushFront(2)
	l.PushFront(1)
	i := l.Begin()
	_ = l.EraseAfter(i)
	assert.Equal(t, 2, l.Size())
	_ = l.EraseAfter(i)
	assert.Equal(t, 1, l.Size())
}

func BenchmarkList_EraseAfter(b *testing.B) {
	l := newList()
	i := l.Begin()
	for n := 0; n < b.N; n++ {
		if l.Size() <= 1 {
			b.StopTimer()
			for k := 1; k < 1000000; k++ {
				l.PushFront(1)
			}
			i = l.Begin()
			b.StartTimer()
		}
		_ = l.EraseAfter(i)
	}
}
