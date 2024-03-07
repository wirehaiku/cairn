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

func TestIn(t *testing.T) {
	// setup
	ss := []string{"a", "b", "c", "d"}

	// success - true
	b := In("a", ss)
	assert.True(t, b)

	// success - false
	b = In("nope", ss)
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
