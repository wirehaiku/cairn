package cairn

import (
	"os"
)

// ExitFunc is the default system exit function.
var ExitFunc = os.Exit

// Funcs is the default map of Cairn program functions.
var Funcs = map[string]CairnFunc{
	"+":   MathAddFunc,
	"-":   MathSubFunc,
	"==":  LogicEqualFunc,
	"<":   MathLesserThanFunc,
	">":   MathGreaterThanFunc,
	"die": IOExitFunc,
	"clr": StackClearFunc,
	"def": SystemDefineFunc,
	"eva": SystemEvalFunc,
	"get": TableGetFunc,
	"ift": LogicIfTrueFunc,
	"iff": LogicIfFalseFunc,
	"for": LogicLoopFunc,
	"inn": IOReadFunc,
	"out": IOWriteFunc,
	"nop": LogicNoOpFunc,
	"set": TableSetFunc,
}

// IOExitFunc (a --) exits the program with an integer exit code.
func IOExitFunc(c *Cairn) error {
	return Pure(c, 1, func(is []int) {
		ExitFunc(is[0])
	})
}

// IOReadFunc (-- a) pushes an input character as an integer.
func IOReadFunc(c *Cairn) error {
	r := c.Read()
	c.Stack.Push(int(r))
	return nil
}

// IOWriteFunc (a --) writes an integer as an output character.
func IOWriteFunc(c *Cairn) error {
	return Pure(c, 1, func(is []int) {
		r := rune(is[0])
		c.Write(r)
	})
}

// LogicEqualFunc (a b -- c) pushes true if a == b.
func LogicEqualFunc(c *Cairn) error {
	return PurePush(c, 2, func(is []int) int {
		return Bool(is[1] == is[0])
	})
}

// LogicIfFalseFunc (a --) evaluates code if a is false.
func LogicIfFalseFunc(c *Cairn) error {
	i, err := c.Stack.Pop()
	if err != nil {
		return err
	}

	as, err := DequeueEnd(c.Queue)
	if err != nil {
		return err
	}

	if i == 0 {
		return c.EvaluateAll(as)
	}

	return nil
}

// LogicIfTrueFunc (a --) evaluates code if a is true.
func LogicIfTrueFunc(c *Cairn) error {
	i, err := c.Stack.Pop()
	if err != nil {
		return err
	}

	as, err := DequeueEnd(c.Queue)
	if err != nil {
		return err
	}

	if i != 0 {
		return c.EvaluateAll(as)
	}

	return nil
}

// LogicLoopFunc (--) repeats code until a register is zero.
func LogicLoopFunc(c *Cairn) error {
	a, err := c.Queue.Dequeue()
	if err != nil {
		return err
	}

	i, err := ToInteger(a)
	if err != nil {
		return err
	}

	as, err := DequeueEnd(c.Queue)
	if err != nil {
		return err
	}

	for {
		if err := c.EvaluateAll(as); err != nil {
			return err
		}

		if c.Table.Get(i) == 0 {
			break
		}
	}

	return nil
}

// LogicNoOpFunc does nothing.
func LogicNoOpFunc(c *Cairn) error {
	return nil
}

// MathAddFunc (a b -- c) pushes the sum of the top two integers on the Stack.
func MathAddFunc(c *Cairn) error {
	return PurePush(c, 2, func(is []int) int {
		return is[1] + is[0]
	})
}

// MathGreaterThanFunc (a b -- c) pushes true if a > b.
func MathGreaterThanFunc(c *Cairn) error {
	return PurePush(c, 2, func(is []int) int {
		return Bool(is[1] > is[0])
	})
}

// MathLesserThanFunc (a b -- c) pushes true if a < b.
func MathLesserThanFunc(c *Cairn) error {
	return PurePush(c, 2, func(is []int) int {
		return Bool(is[1] < is[0])
	})
}

// MathSubFunc (a b -- c) pushes the difference of the top two integers on the Stack.
func MathSubFunc(c *Cairn) error {
	return PurePush(c, 2, func(is []int) int {
		return is[1] - is[0]
	})
}

// StackClearFunc (--) clears the Stack.
func StackClearFunc(c *Cairn) error {
	c.Stack.Clear()
	return nil
}

// SystemDefineFunc (--) sets a function in the Cairn.
func SystemDefineFunc(c *Cairn) error {
	a, err := c.Queue.Dequeue()
	if err != nil {
		return err
	}

	s, err := ToSymbol(a)
	if err != nil {
		return err
	}

	as, err := DequeueEnd(c.Queue)
	if err != nil {
		return err
	}

	c.SetFuncAtoms(s, as)
	return nil
}

// SystemEvalFunc (... --) evaluates all integers in the Stack up to a newline as a string.
func SystemEvalFunc(c *Cairn) error {
	is, err := c.Stack.PopTo(10)
	if err != nil {
		return err
	}

	var rs []rune
	for _, i := range is {
		rs = append(rs, rune(i))
	}

	ss := Tokenise(string(rs))
	as := AtomiseAll(ss)
	return c.EvaluateAll(as)
}

// TableGetFunc (a -- b) pushes a value from the Table.
func TableGetFunc(c *Cairn) error {
	return PurePush(c, 1, func(is []int) int {
		return c.Table.Get(is[0])
	})
}

// TableSetFunc (a b --) sets a value in the Table.
func TableSetFunc(c *Cairn) error {
	return Pure(c, 2, func(is []int) {
		c.Table.Set(is[0], is[1])
	})
}
