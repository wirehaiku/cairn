package cairn

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func xAdd(c *Cairn) error {
	is, err := c.Stack.PopN(2)
	if err != nil {
		return err
	}

	c.Stack.Push(is[0] + is[1])
	return nil
}

func xCairn(s string) (*Cairn, *bytes.Buffer) {
	ib := bytes.NewBufferString(s)
	ob := bytes.NewBuffer(nil)
	return NewCairn(
		map[string]CairnFunc{"ADD": xAdd},
		bufio.NewReader(ib),
		bufio.NewWriter(ob),
	), ob
}

func TestNewCairn(t *testing.T) {
	// success
	c, _ := xCairn("")
	assert.NotNil(t, c.Queue)
	assert.NotNil(t, c.Stack)
	assert.NotNil(t, c.Table)
	assert.NotNil(t, c.Funcs)
	assert.NotNil(t, c.Input)
	assert.NotNil(t, c.Output)
}

func TestCairnEvaluate(t *testing.T) {
	// setup
	c, _ := xCairn("")
	c.Stack.Push(1)

	// success - integer
	err := c.Evaluate(2)
	assert.Equal(t, []int{1, 2}, c.Stack.Integers)
	assert.NoError(t, err)

	// success - symbol string
	err = c.Evaluate("ADD")
	assert.Equal(t, []int{3}, c.Stack.Integers)
	assert.NoError(t, err)

	// setup
	c.Stack.Integers = []int{1, 2}

	// success - cairn function
	err = c.Evaluate(CairnFunc(xAdd))
	assert.Equal(t, []int{3}, c.Stack.Integers)
	assert.NoError(t, err)

	// failure - invalid type
	err = c.Evaluate(false)
	assert.EqualError(t, err, `cannot evaluate atom type "bool"`)
}

func TestCairnEvaluateAll(t *testing.T) {
	// setup
	c, _ := xCairn("")

	// success
	err := c.EvaluateAll([]any{1, 2, "ADD"})
	assert.Equal(t, []int{3}, c.Stack.Integers)
	assert.NoError(t, err)
}

func TestCairnEvaluateString(t *testing.T) {
	// setup
	c, _ := xCairn("")

	// success
	err := c.EvaluateString("1 2 ADD")
	assert.Equal(t, []int{3}, c.Stack.Integers)
	assert.NoError(t, err)
}

func TestCairnGetFunc(t *testing.T) {
	// setup
	c, _ := xCairn("")

	// success
	f, err := c.GetFunc("ADD")
	assert.NotNil(t, f)
	assert.NoError(t, err)

	// failure - function does not exist
	f, err = c.GetFunc("NOPE")
	assert.Nil(t, f)
	assert.EqualError(t, err, `function "NOPE" does not exist`)
}

func TestCairnRead(t *testing.T) {
	// setup
	c, _ := xCairn("test\n")

	// success
	s := c.Read()
	assert.Equal(t, "test\n", s)
}

func TestCairnSetFunc(t *testing.T) {
	// setup
	c, _ := xCairn("")

	// success
	c.SetFunc("TEST", xAdd)
	assert.NotNil(t, c.Funcs["TEST"])
}

func TestCairnSetFuncAtoms(t *testing.T) {
	// setup
	c, _ := xCairn("")

	// success
	c.SetFuncAtoms("TEST", []any{1, 2, "ADD"})
	err := c.Evaluate("TEST")
	assert.NotNil(t, c.Funcs["TEST"])
	assert.Equal(t, []int{3}, c.Stack.Integers)
	assert.NoError(t, err)
}

func TestCairnWrite(t *testing.T) {
	// setup
	c, b := xCairn("")

	// success
	c.Write("%s\n", "test")
	assert.Equal(t, "test\n", b.String())
}