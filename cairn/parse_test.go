package cairn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomise(t *testing.T) {
	// success - integer
	a := Atomise("123")
	assert.Equal(t, 123, a)

	// success - symbol
	a = Atomise("foo")
	assert.Equal(t, "foo", a)
}

func TestAtomiseAll(t *testing.T) {
	// success
	as := AtomiseAll([]string{"123", "foo"})
	assert.Equal(t, []any{123, "foo"}, as)
}

func TestTokenise(t *testing.T) {
	// setup
	s := `
		// comment
		123 foo // comment
		// comment
	`

	ss := Tokenise(s)
	assert.Equal(t, []string{"123", "foo"}, ss)
}
