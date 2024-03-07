package cairn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFlags(t *testing.T) {
	// setup
	ss := []string{"-c", "cmd"}

	// success
	fs, err := ParseFlags(ss)
	assert.Equal(t, "cmd", fs.Command)
	assert.NoError(t, err)
}
