package cairn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	// success - true
	i := Bool(true)
	assert.Equal(t, 1, i)

	// success - false
	i = Bool(false)
	assert.Equal(t, 0, i)
}

func TestDequeueEnd(t *testing.T) {
	// setup
	q := NewQueue("ift", 123, "end", "end", "nop")

	// success
	as, err := DequeueEnd(q)
	assert.Equal(t, []any{"ift", 123, "end"}, as)
	assert.Equal(t, []any{"nop"}, q.Atoms)
	assert.NoError(t, err)
}

func TestIn(t *testing.T) {
	// setup
	as := []any{"a", "b", "c", "d"}

	// success - true
	b := In("a", as)
	assert.True(t, b)

	// success - false
	b = In("nope", as)
	assert.False(t, b)
}

func TestPure(t *testing.T) {
	// setup
	c, _ := xCairn("")
	c.Stack.Push(123)
	var i int

	// success
	err := Pure(c, 1, func(is []int) { i = is[0] })
	assert.Equal(t, 123, i)
	assert.NoError(t, err)

	// failure - stack is empty
	err = Pure(c, 2, nil)
	assert.EqualError(t, err, "stack is empty")
}

func TestPurePush(t *testing.T) {
	// setup
	c, _ := xCairn("")
	c.Stack.PushAll([]int{1, 2})

	// success
	err := PurePush(c, 2, func(is []int) int { return is[0] + is[1] })
	assert.Equal(t, []int{3}, c.Stack.Integers)
	assert.NoError(t, err)

	// failure - stack is empty
	err = PurePush(c, 2, nil)
	assert.EqualError(t, err, "stack is empty")
}

func TestToSymbol(t *testing.T) {
	// success
	s, err := ToSymbol("foo")
	assert.IsType(t, "foo", s)
	assert.NoError(t, err)

	// failure - non-symbol
	s, err = ToSymbol(123)
	assert.Empty(t, s)
	assert.EqualError(t, err, `non-symbol "123" provided`)
}
