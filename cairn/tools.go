package cairn

// Bool returns a boolean as an integer.
func Bool(b bool) int {
	if b {
		return 1
	}

	return 0
}

// In returns true if a string is in a slice.
func In(s string, ss []string) bool {
	for _, s2 := range ss {
		if s == s2 {
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
