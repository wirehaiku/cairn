package cairn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFlags(t *testing.T) {
	// setup
	ss := []string{"-c", "cmd", "a.txt", "b.txt"}

	// success
	f, err := ParseFlags(ss)
	assert.Equal(t, "cmd", f.Command)
	assert.Equal(t, []string{"a.txt", "b.txt"}, f.Files)
	assert.NoError(t, err)
}
