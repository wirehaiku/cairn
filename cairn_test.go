///////////////////////////////////////////////////////////////////////////////////////
//                                  Cairn Unit Tests                                 //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

///////////////////////////////////////////////////////////////////////////////////////
//                         Part 1: Helper Globals & Functions                        //
///////////////////////////////////////////////////////////////////////////////////////

// AS is a shorthand function for atom slices.
func AS(as ...any) []any { return as }

// US is a shorthand function for integer slices.
func US(us ...uint8) []uint8 { return us }

///////////////////////////////////////////////////////////////////////////////////////
//                        Part 2: Testing Collection Functions                       //
///////////////////////////////////////////////////////////////////////////////////////

// 2.1: Testing Queue Functions
////////////////////////////////

func TestDequeue(t *testing.T) {
	// setup
	Queue = AS("a")

	// success
	a, err := Dequeue()
	assert.Equal(t, "a", a)
	assert.NoError(t, err)

	// failure - ErrQueueEmpty
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

	// failure - ErrQueueEmpty
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
	EnqueueAll(AS("a", "b"))
	assert.Equal(t, AS("a", "b"), Queue)
}

// 2.2: Testing Register Functions
///////////////////////////////////

func TestGetRegister(t *testing.T) {
	// setup
	Registers[0] = 255

	// success
	u, err := GetRegister(0)
	assert.Equal(t, uint8(255), u)
	assert.NoError(t, err)

	// failure - ErrRegisterNone
	u, err = GetRegister(99)
	assert.Zero(t, u)
	assert.Equal(t, ErrRegisterNone, err)
}

func TestSetRegister(t *testing.T) {
	// setup
	Registers[0] = 0

	// success
	err := SetRegister(0, 255)
	assert.Equal(t, uint8(255), Registers[0])
	assert.NoError(t, err)

	// failure - ErrRegisterNone
	err = SetRegister(99, 255)
	assert.Equal(t, ErrRegisterNone, err)
}

// 2.3: Testing Stack Functions
////////////////////////////////

func TestPop(t *testing.T) {
	// setup
	Stack = US(255)

	// success
	u, err := Pop()
	assert.Equal(t, uint8(255), u)
	assert.NoError(t, err)

	// failure - ErrStackEmpty
	u, err = Pop()
	assert.Zero(t, u)
	assert.Equal(t, ErrStackEmpty, err)
}

func TestPopN(t *testing.T) {
	// setup
	Stack = US(255, 255)

	// success
	us, err := PopN(2)
	assert.Equal(t, US(255, 255), us)
	assert.NoError(t, err)

	// failure - ErrStackEmpty
	us, err = PopN(1)
	assert.Zero(t, us)
	assert.Equal(t, ErrStackEmpty, err)
}

func TestPush(t *testing.T) {
	// setup
	Stack = US()

	// success
	Push(255)
	assert.Equal(t, US(255), Stack)
}

func TestPushAll(t *testing.T) {
	// setup
	Stack = US()

	// success
	PushAll(US(255, 255))
	assert.Equal(t, US(255, 255), Stack)
}
