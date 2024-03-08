package cairn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIOExitFunc(t *testing.T) {
	// setup
	var x int
	c, _ := xCairn("")
	c.Stack.Push(123)
	ExitFunc = func(i int) { x = i }

	// success
	err := IOExitFunc(c)
	assert.Equal(t, 123, x)
	assert.NoError(t, err)
}

func TestIOReadFunc(t *testing.T) {
	// setup
	c, _ := xCairn("test\n")

	// success
	err := IOReadFunc(c)
	assert.Equal(t, []int{116}, c.Stack.Integers)
	assert.NoError(t, err)
}

func TestIOWriteFunc(t *testing.T) {
	// setup
	c, b := xCairn("")
	c.Stack.Push(116)

	// success
	err := IOWriteFunc(c)
	assert.Equal(t, "t", b.String())
	assert.NoError(t, err)
}

func TestLogicEqualFunc(t *testing.T) {
	// setup
	c, _ := xCairn("")
	c.Stack.PushAll([]int{1, 1})

	// success - true
	err := LogicEqualFunc(c)
	assert.Equal(t, []int{1}, c.Stack.Integers)
	assert.NoError(t, err)

	// setup
	c.Stack.Clear()
	c.Stack.PushAll([]int{1, 0})

	// success
	err = LogicEqualFunc(c)
	assert.Equal(t, []int{0}, c.Stack.Integers)
	assert.NoError(t, err)
}

func TestLogicIfFalseFunc(t *testing.T) {
	// setup
	c, _ := xCairn("")
	c.Queue.EnqueueAll([]any{123, "end"})
	c.Stack.Push(0)

	// success - true
	err := LogicIfFalseFunc(c)
	assert.Empty(t, c.Queue.Atoms)
	assert.Equal(t, []int{123}, c.Stack.Integers)
	assert.NoError(t, err)

	// setup
	c.Queue.EnqueueAll([]any{123, "end"})
	c.Stack.Clear()
	c.Stack.Push(1)

	// success - false
	err = LogicIfFalseFunc(c)
	assert.Empty(t, c.Queue.Atoms)
	assert.Empty(t, c.Stack.Integers)
	assert.NoError(t, err)
}

func TestLogicIfTrueFunc(t *testing.T) {
	// setup
	c, _ := xCairn("")
	c.Queue.EnqueueAll([]any{123, "end"})
	c.Stack.Push(1)

	// success - true
	err := LogicIfTrueFunc(c)
	assert.Empty(t, c.Queue.Atoms)
	assert.Equal(t, []int{123}, c.Stack.Integers)
	assert.NoError(t, err)

	// setup
	c.Queue.EnqueueAll([]any{123, "end"})
	c.Stack.Clear()
	c.Stack.Push(0)

	// success - false
	err = LogicIfTrueFunc(c)
	assert.Empty(t, c.Queue.Atoms)
	assert.Empty(t, c.Stack.Integers)
	assert.NoError(t, err)
}

func TestLogicLoopFunc(t *testing.T) {
	// setup
	c, _ := xCairn("")
	c.Queue.EnqueueAll([]any{0, 123, "end"})

	// success
	err := LogicLoopFunc(c)
	assert.Equal(t, []int{123}, c.Stack.Integers)
	assert.NoError(t, err)
}

func TestLogicNoOpFunc(t *testing.T) {
	// setup
	c, _ := xCairn("")

	// success
	err := LogicNoOpFunc(c)
	assert.NoError(t, err)
}

func TestMathAddFunc(t *testing.T) {
	// setup
	c, _ := xCairn("")
	c.Stack.PushAll([]int{1, 2})

	// success
	err := MathAddFunc(c)
	assert.Equal(t, []int{3}, c.Stack.Integers)
	assert.NoError(t, err)
}

func TestMathGreaterThanFunc(t *testing.T) {
	// setup
	c, _ := xCairn("")
	c.Stack.PushAll([]int{2, 1})

	// success - true
	err := MathGreaterThanFunc(c)
	assert.Equal(t, []int{1}, c.Stack.Integers)
	assert.NoError(t, err)

	// setup
	c.Stack.Clear()
	c.Stack.PushAll([]int{1, 2})

	// success - false
	err = MathGreaterThanFunc(c)
	assert.Equal(t, []int{0}, c.Stack.Integers)
	assert.NoError(t, err)
}

func TestMathLesserThanFunc(t *testing.T) {
	// setup
	c, _ := xCairn("")
	c.Stack.PushAll([]int{1, 2})

	// success - true
	err := MathLesserThanFunc(c)
	assert.Equal(t, []int{1}, c.Stack.Integers)
	assert.NoError(t, err)

	// setup
	c.Stack.Clear()
	c.Stack.PushAll([]int{2, 1})

	// success - false
	err = MathLesserThanFunc(c)
	assert.Equal(t, []int{0}, c.Stack.Integers)
	assert.NoError(t, err)
}

func TestMathSubFunc(t *testing.T) {
	// setup
	c, _ := xCairn("")
	c.Stack.PushAll([]int{3, 2})

	// success
	err := MathSubFunc(c)
	assert.Equal(t, []int{1}, c.Stack.Integers)
	assert.NoError(t, err)
}

func TestStackClearFunc(t *testing.T) {
	// setup
	c, _ := xCairn("")
	c.Stack.PushAll([]int{1, 2, 3})

	// success
	err := StackClearFunc(c)
	assert.Empty(t, c.Stack.Integers)
	assert.NoError(t, err)
}

func TestSystemDefineFunc(t *testing.T) {
	// setup
	c, _ := xCairn("")
	c.Queue.EnqueueAll([]any{"foo", 123, "end"})

	// success
	err := SystemDefineFunc(c)
	assert.Empty(t, c.Queue.Atoms)
	assert.NotNil(t, c.Funcs["foo"])
	assert.NoError(t, err)

	// success - function test
	err = c.Funcs["foo"](c)
	assert.Equal(t, []int{123}, c.Stack.Integers)
	assert.NoError(t, err)
}

func TestTableGetFunc(t *testing.T) {
	// setup
	c, _ := xCairn("")
	c.Stack.Push(0)
	c.Table.Set(0, 123)

	// success
	err := TableGetFunc(c)
	assert.Equal(t, []int{123}, c.Stack.Integers)
	assert.NoError(t, err)
}

func TestTableSetFunc(t *testing.T) {
	// setup
	c, _ := xCairn("")
	c.Stack.PushAll([]int{123, 0})

	// success
	err := TableSetFunc(c)
	assert.Equal(t, map[int]int{0: 123}, c.Table.Integers)
	assert.NoError(t, err)
}
