///////////////////////////////////////////////////////////////////////////////////////
//                                  Cairn Unit Tests                                 //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

///////////////////////////////////////////////////////////////////////////////////////
//                       Part 1: Unit Testing Helper Functions                       //
///////////////////////////////////////////////////////////////////////////////////////

// AS returns any value slice as an atom slice.
func AS(as ...any) []any { return as }

///////////////////////////////////////////////////////////////////////////////////////
//                        Part 2: Testing Collection Functions                       //
///////////////////////////////////////////////////////////////////////////////////////

func TestDequeue(t *testing.T) {
	// setup
	Queue = AS("a")

	// success
	a, err := Dequeue()
	assert.Equal(t, "a", a)
	assert.NoError(t, err)

	// failure - queue empty
	a, err = Dequeue()
	assert.Nil(t, a)
	assert.Equal(t, ErrQueueEmpty, err)
}

func TestDequeueTo(t *testing.T) {
	// setup
	Queue = AS("a", "b", "end", "c")

	// success
	as, err := DequeueTo("end")
	assert.Equal(t, AS("a", "b"), as)
	assert.Equal(t, AS("c"), Queue)
	assert.NoError(t, err)

	// failure - queue empty
	as, err = DequeueTo("end")
	assert.Empty(t, as)
	assert.Empty(t, Queue)
	assert.Equal(t, ErrQueueEmpty, err)
}

func TestEnqueue(t *testing.T) {
	// setup
	Queue = AS()

	// success
	Enqueue("a")
	assert.Equal(t, AS("a"), Queue)
}

func TestEnqueueAll(t *testing.T) {
	// setup
	Queue = AS()

	// success
	EnqueueAll([]any{"a", "b"})
	assert.Equal(t, AS("a", "b"), Queue)
}
