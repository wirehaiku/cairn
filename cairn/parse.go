package cairn

import (
	"strconv"
	"strings"
)

// Atomise returns an atom from a token string.
func Atomise(s string) any {
	if i, err := strconv.ParseInt(s, 10, 0); err == nil {
		return int(i)
	}

	return s
}

// AtomiseAll returns an atom slice from a token slice.
func AtomiseAll(ss []string) []any {
	var as []any
	for _, s := range ss {
		as = append(as, Atomise(s))
	}

	return as
}

// Tokenise returns a token slice from a program string.
func Tokenise(s string) []string {
	var ss []string
	s = strings.ToUpper(s)

	for _, s := range strings.Split(s, "\n") {
		s = strings.SplitN(s, "//", 2)[0]
		ss = append(ss, strings.Fields(s)...)
	}

	return ss
}
