package cairn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	// success
	s := NewStack(1, 2, 3)
	assert.Equal(t, []int{1, 2, 3}, s.Integers)
}

func TestStackClear(t *testing.T) {
	// setup
	s := NewStack(1, 2, 3)

	// success
	s.Clear()
	assert.Empty(t, s.Integers)
}

func TestStackEmpty(t *testing.T) {
	// success - true
	b := NewStack().Empty()
	assert.True(t, b)

	// success - false
	b = NewStack(1, 2, 3).Empty()
	assert.False(t, b)
}

func TestStackLen(t *testing.T) {
	// success
	n := NewStack(1, 2, 3).Len()
	assert.Equal(t, 3, n)
}

func TestStackPop(t *testing.T) {
	// setup
	s := NewStack(1)

	// success
	i, err := s.Pop()
	assert.Equal(t, 1, i)
	assert.Empty(t, s.Integers)
	assert.NoError(t, err)

	// failure - stack is empty
	i, err = s.Pop()
	assert.Zero(t, i)
	assert.EqualError(t, err, "stack is empty")
}

func TestStackPopN(t *testing.T) {
	// setup
	s := NewStack(1, 2, 3)

	// success
	is, err := s.PopN(3)
	assert.Equal(t, []int{3, 2, 1}, is)
	assert.Empty(t, s.Integers)
	assert.NoError(t, err)

	// failure - stack is empty
	is, err = s.PopN(1)
	assert.Empty(t, is)
	assert.EqualError(t, err, "stack is empty")
}

func TestStackPush(t *testing.T) {
	// setup
	s := NewStack(1, 2, 3)

	// success
	s.Push(4)
	assert.Equal(t, []int{1, 2, 3, 4}, s.Integers)
}

func TestStackPushAll(t *testing.T) {
	// setup
	s := NewStack(1, 2, 3)

	// success
	s.PushAll([]int{4, 5, 6})
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, s.Integers)
}
