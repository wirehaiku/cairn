package cairn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func xTest(t *testing.T, c *Cairn, s string, is ...int) {
	c.Queue.Clear()
	c.Stack.Clear()
	c.Table.Clear()
	err := c.Execute(s)
	msg := fmt.Sprintf("%q should equal %v\n", s, is)
	assert.Equal(t, is, c.Stack.Integers, msg)
	assert.NoError(t, err, msg)
}

func TestLibrary(t *testing.T) {
	// setup
	c, _ := xCairn("")

	// success - evaluate library
	err := c.Execute(Library)
	assert.NoError(t, err)

	// success - stack functions
	xTest(t, c, "1 dup", 1, 1)
	xTest(t, c, "1 2 drop", 1)
	xTest(t, c, "1 2 swap", 2, 1)

	// success - operator functions
	xTest(t, c, "0 1 !=", 1)
	xTest(t, c, "1 1 !=", 0)
	xTest(t, c, "2 3 <=", 1)
	xTest(t, c, "3 3 <=", 1)
	xTest(t, c, "4 3 <=", 0)
	xTest(t, c, "4 3 >=", 1)
	xTest(t, c, "3 3 >=", 1)
	xTest(t, c, "2 3 >=", 0)

	// success - logic functions
	xTest(t, c, "0 f?", 1)
	xTest(t, c, "1 f?", 0)
	xTest(t, c, "1 t?", 1)
	xTest(t, c, "0 t?", 0)
	xTest(t, c, "0 0 and", 0)
	xTest(t, c, "0 1 and", 0)
	xTest(t, c, "1 0 and", 0)
	xTest(t, c, "1 1 and", 1)
	xTest(t, c, "0 0 or", 0)
	xTest(t, c, "0 1 or", 1)
	xTest(t, c, "1 0 or", 1)
	xTest(t, c, "1 1 or", 1)
	xTest(t, c, "0 0 xor", 0)
	xTest(t, c, "0 1 xor", 1)
	xTest(t, c, "1 0 xor", 1)
	xTest(t, c, "1 1 xor", 0)
}
