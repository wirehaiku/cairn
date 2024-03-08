package cairn

import "fmt"

// Bool returns a boolean as an integer.
func Bool(b bool) int {
	if b {
		return 1
	}

	return 0
}

// DequeueEnd removes returns all atoms in the Queue up to an "end" atom.
func DequeueEnd(q *Queue) ([]any, error) {
	var as []any
	var ac int = 1

loop:
	for !q.Empty() {
		a, err := q.Dequeue()
		if err != nil {
			return nil, err
		}

		as = append(as, a)

		if In(a, []any{"def", "iff", "ift", "for"}) {
			ac++
		} else if a == "end" {
			ac--
		}

		if a == "end" && ac == 0 {
			break loop
		}
	}

	return as[:len(as)-1], nil
}

// In returns true if an atom is in a slice.
func In(a any, as []any) bool {
	for _, a2 := range as {
		if a == a2 {
			return true
		}
	}

	return false
}

// Pure applies a pure integer function to a Cairn.
func Pure(c *Cairn, n int, f func([]int)) error {
	is, err := c.Stack.PopN(n)
	if err != nil {
		return err
	}

	f(is)
	return nil
}

// PurePush applies a pure integer function to a Cairn and pushes the result.
func PurePush(c *Cairn, n int, f func([]int) int) error {
	is, err := c.Stack.PopN(n)
	if err != nil {
		return err
	}

	c.Stack.Push(f(is))
	return nil
}

// ToInteger returns an atom as an integer.
func ToInteger(a any) (int, error) {
	switch a := a.(type) {
	case int:
		return a, nil
	default:
		return 0, fmt.Errorf(`non-integer "%v" provided`, a)
	}
}

// ToSymbol returns an atom as a symbol string.
func ToSymbol(a any) (string, error) {
	switch a := a.(type) {
	case string:
		return a, nil
	default:
		return "", fmt.Errorf(`non-symbol "%v" provided`, a)
	}
}
