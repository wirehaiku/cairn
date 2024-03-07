package cairn

import "os"

// ExitFunc is the default system exit function.
var ExitFunc = os.Exit

// Funcs is the default map of Cairn program functions.
var Funcs = map[string]CairnFunc{
	"+":   MathAddFunc,
	"-":   MathSubFunc,
	"==":  LogicEqualFunc,
	"<":   MathLesserThanFunc,
	">":   MathGreaterThanFunc,
	"BYE": IOExitFunc,
	"CLR": StackClearFunc,
	// "DEF": SystemDefineFunc,
	"GET": TableGetFunc,
	// "IF": LogicIfTrueFunc,
	// "IFF": LogicIfFalseFunc,
	// "FOR": LogicLoopFunc,
	"INN": IOReadFunc,
	"OUT": IOWriteFunc,
	"NOP": LogicNoOpFunc,
	"SET": TableSetFunc,
}

// IOExitFunc (a --) exits the program with an integer exit code.
func IOExitFunc(c *Cairn) error {
	return Pure(c, 1, func(is []int) {
		ExitFunc(is[0])
	})
}

// IOReadFunc (-- a) pushes an input character as an integer.
func IOReadFunc(c *Cairn) error {
	r, _ := c.Input.ReadByte()
	c.Stack.Push(int(r))
	return nil
}

// IOWriteFunc (a --) writes an integer as an output character.
func IOWriteFunc(c *Cairn) error {
	return Pure(c, 1, func(is []int) {
		b := byte(is[0])
		c.Output.WriteByte(b)
		c.Output.Flush()
	})
}

// LogicEqualFunc (a b -- c) pushes true if a == b.
func LogicEqualFunc(c *Cairn) error {
	return PurePush(c, 2, func(is []int) int {
		return Bool(is[0] == is[1])
	})
}

// LogicNoOpFunc does nothing.
func LogicNoOpFunc(c *Cairn) error {
	return nil
}

// MathAddFunc (a b -- c) pushes the sum of the top two integers on the Stack.
func MathAddFunc(c *Cairn) error {
	return PurePush(c, 2, func(is []int) int {
		return is[0] + is[1]
	})
}

// MathGreaterThanFunc (a b -- c) pushes true if a > b.
func MathGreaterThanFunc(c *Cairn) error {
	return PurePush(c, 2, func(is []int) int {
		return Bool(is[0] > is[1])
	})
}

// MathLesserThanFunc (a b -- c) pushes true if a < b.
func MathLesserThanFunc(c *Cairn) error {
	return PurePush(c, 2, func(is []int) int {
		return Bool(is[0] < is[1])
	})
}

// MathSubFunc (a b -- c) pushes the difference of the top two integers on the Stack.
func MathSubFunc(c *Cairn) error {
	return PurePush(c, 2, func(is []int) int {
		return is[0] - is[1]
	})
}

// StackClearFunc (--) clears the Stack.
func StackClearFunc(c *Cairn) error {
	c.Stack.Clear()
	return nil
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
