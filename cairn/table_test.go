package cairn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTable(t *testing.T) {
	// success
	ta := NewTable(map[int]int{0: 123})
	assert.Equal(t, map[int]int{0: 123}, ta.Integers)
}

func TestTableClear(t *testing.T) {
	// setup
	ta := NewTable(map[int]int{0: 123})

	// success
	ta.Clear()
	assert.Empty(t, ta.Integers)
}

func TestTableDelete(t *testing.T) {
	// setup
	ta := NewTable(map[int]int{0: 123})

	// success
	ta.Delete(0)
	assert.Empty(t, ta.Integers)
}

func TestTableEmpty(t *testing.T) {
	// success - true
	b := NewTable(nil).Empty()
	assert.True(t, b)

	// success - false
	b = NewTable(map[int]int{0: 123}).Empty()
	assert.False(t, b)
}

func TestTableGet(t *testing.T) {
	// setup
	ta := NewTable(map[int]int{0: 123})

	// success
	i := ta.Get(0)
	assert.Equal(t, 123, i)
}

func TestTableHas(t *testing.T) {
	// setup
	ta := NewTable(map[int]int{0: 123})

	// success - true
	b := ta.Has(0)
	assert.True(t, b)

	// success - false
	b = ta.Has(1)
	assert.False(t, b)
}

func TestTableLen(t *testing.T) {
	// setup
	ta := NewTable(map[int]int{0: 123})

	// success
	n := ta.Len()
	assert.Equal(t, 1, n)
}

func TestTableSet(t *testing.T) {
	// setup
	ta := NewTable(map[int]int{0: 123})

	// success
	ta.Set(0, 456)
	assert.Equal(t, map[int]int{0: 456}, ta.Integers)
}
