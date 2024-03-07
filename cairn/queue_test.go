package cairn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	// success
	q := NewQueue("a", "b", "c")
	assert.Equal(t, []any{"a", "b", "c"}, q.Atoms)
}

func TestClear(t *testing.T) {
	// setup
	q := NewQueue("a", "b", "c")

	// success
	q.Clear()
	assert.Empty(t, q.Atoms)
}

func TestEmpty(t *testing.T) {
	// success - true
	b := NewQueue().Empty()
	assert.True(t, b)

	// success - false
	b = NewQueue("a", "b", "c").Empty()
	assert.False(t, b)
}

func TestLen(t *testing.T) {
	// success
	n := NewQueue("a", "b", "c").Len()
	assert.Equal(t, 3, n)
}

func TestDequeue(t *testing.T) {
	// setup
	q := NewQueue("a")

	// success
	a, err := q.Dequeue()
	assert.Equal(t, "a", a)
	assert.Empty(t, q.Atoms)
	assert.NoError(t, err)

	// failure - queue is empty
	a, err = q.Dequeue()
	assert.Nil(t, a)
	assert.EqualError(t, err, "queue is empty")
}

func TestDequeueTo(t *testing.T) {
	// setup
	q := NewQueue("a", "b", "c")

	// success
	as, err := q.DequeueTo("b")
	assert.Equal(t, []any{"a"}, as)
	assert.Equal(t, []any{"c"}, q.Atoms)
	assert.NoError(t, err)

	// failure - queue is empty
	as, err = q.DequeueTo("d")
	assert.Nil(t, as)
	assert.EqualError(t, err, "queue is empty")
}

func TestEnqueue(t *testing.T) {
	// setup
	q := NewQueue("a", "b", "c")

	// success
	q.Enqueue("d")
	assert.Equal(t, []any{"a", "b", "c", "d"}, q.Atoms)
}

func TestEnqueueAll(t *testing.T) {
	// setup
	q := NewQueue("a", "b", "c")

	// success
	q.EnqueueAll([]any{"d", "e", "f"})
	assert.Equal(t, []any{"a", "b", "c", "d", "e", "f"}, q.Atoms)
}
