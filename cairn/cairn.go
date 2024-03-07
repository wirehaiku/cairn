package cairn

import (
	"bufio"
	"fmt"
)

// Cairn is a complete program environment.
type Cairn struct {
	Queue  *Queue
	Stack  *Stack
	Table  *Table
	Funcs  map[string]CairnFunc
	Input  *bufio.Reader
	Output *bufio.Writer
}

// CairnFunc is a Cairn program function.
type CairnFunc func(*Cairn) error

// NewCairn returns a pointer to a new Cairn.
func NewCairn(fm map[string]CairnFunc, r *bufio.Reader, w *bufio.Writer) *Cairn {
	return &Cairn{NewQueue(), NewStack(), NewTable(nil), fm, r, w}
}

// Evaluate evaluates an atom against the Cairn.
func (c *Cairn) Evaluate(a any) error {
	switch a := a.(type) {
	case int:
		c.Stack.Push(a)
		return nil

	case string:
		f, err := c.GetFunc(a)
		if err != nil {
			return err
		}

		return f(c)

	case CairnFunc:
		return a(c)

	default:
		return fmt.Errorf(`cannot evaluate atom type "%T"`, a)
	}
}

// EvaluateAll evaluates an atom slice against the Cairn.
func (c *Cairn) EvaluateAll(as []any) error {
	for _, a := range as {
		if err := c.Evaluate(a); err != nil {
			return err
		}
	}

	return nil
}

// EvaluateString parses a program string and evaluates it against the Cairn.
func (c *Cairn) EvaluateString(s string) error {
	ss := Tokenise(s)
	as := AtomiseAll(ss)
	return c.EvaluateAll(as)
}

// GetFunc returns a CairnFunc from the Cairn.
func (c *Cairn) GetFunc(s string) (CairnFunc, error) {
	f, ok := c.Funcs[s]
	if !ok {
		return nil, fmt.Errorf("function %q does not exist", s)
	}

	return f, nil
}

// Read returns a newline-ending string from the Cairn's input Reader.
func (c *Cairn) Read() string {
	s, _ := c.Input.ReadString('\n')
	return s
}

// SetFunc sets a CairnFunc in the Cairn.
func (c *Cairn) SetFunc(s string, f CairnFunc) {
	c.Funcs[s] = f
}

// SetFuncAtoms seta a CairnFunc in the Cairn from an atom slice.
func (c *Cairn) SetFuncAtoms(s string, as []any) {
	c.Funcs[s] = func(c *Cairn) error {
		return c.EvaluateAll(as)
	}
}

// Write writes a formatted string to the Cairn's output Writer.
func (c *Cairn) Write(s string, vs ...any) {
	fmt.Fprintf(c.Output, s, vs...)
	c.Output.Flush()
}
