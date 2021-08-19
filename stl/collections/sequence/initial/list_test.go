package initial_test

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
	l = newList(1)
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
	l = newList(1)
	assert.False(t, l.Empty())
}

func BenchmarkList_Empty(b *testing.B) {
	l := newList()
	for n := 0; n < b.N; n++ {
		_ = l.Empty()
	}
}

func TestList_Begin(t *testing.T) {
	l := newList(1)
	assert.Equal(t, 1, l.Begin().Read())
}

func BenchmarkList_Begin(b *testing.B) {
	l := newList()
	for n := 0; n < b.N; n++ {
		_ = l.Begin()
	}
}

func TestList_End(t *testing.T) {
	l := newList(1)
	end := l.Begin().ImplNext()
	assert.True(t, l.End().Equal(end))
}

func BenchmarkList_End(b *testing.B) {
	l := newList()
	for n := 0; n < b.N; n++ {
		_ = l.End()
	}
}
