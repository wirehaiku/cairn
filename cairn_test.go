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

// SS is a shorthand function for string slices.
func SS(ss ...string) []string { return ss }

// US is a shorthand function for integer slices.
func US(us ...uint8) []uint8 { return us }

///////////////////////////////////////////////////////////////////////////////////////
//                        Part 2: Testing Collection Functions                       //
///////////////////////////////////////////////////////////////////////////////////////

// 2.1: Testing Queue Functions
////////////////////////////////

func TestDequeue(t *testing.T) {
	// setup
	Queue = AS("A")

	// success
	a, err := Dequeue()
	assert.Equal(t, "A", a)
	assert.NoError(t, err)

	// failure - ErrQueueEmpty
	a, err = Dequeue()
	assert.Nil(t, a)
	assert.Equal(t, ErrQueueEmpty, err)
}

func TestDequeueTo(t *testing.T) {
	// setup
	Queue = AS("A", "B", "END", "C")

	// success
	as, err := DequeueTo("END")
	assert.Equal(t, AS("A", "B"), as)
	assert.Equal(t, AS("C"), Queue)
	assert.NoError(t, err)

	// failure - ErrQueueEmpty
	as, err = DequeueTo("END")
	assert.Empty(t, as)
	assert.Empty(t, Queue)
	assert.Equal(t, ErrQueueEmpty, err)
}

func TestEnqueue(t *testing.T) {
	// setup
	Queue = AS()

	// success
	Enqueue("A")
	assert.Equal(t, AS("A"), Queue)
}

func TestEnqueueAll(t *testing.T) {
	// setup
	Queue = AS()

	// success
	EnqueueAll(AS("A", "B"))
	assert.Equal(t, AS("A", "B"), Queue)
}

// 2.2: Testing Register Functions
///////////////////////////////////

func TestGetRegister(t *testing.T) {
	// setup
	Registers[0] = 123

	// success
	u, err := GetRegister(0)
	assert.Equal(t, uint8(123), u)
	assert.NoError(t, err)

	// failure - ErrRegisterNone
	u, err = GetRegister(8)
	assert.Zero(t, u)
	assert.Equal(t, ErrRegisterNone, err)
}

func TestSetRegister(t *testing.T) {
	// setup
	Registers[0] = 0

	// success
	err := SetRegister(0, 123)
	assert.Equal(t, uint8(123), Registers[0])
	assert.NoError(t, err)

	// failure - ErrRegisterNone
	err = SetRegister(8, 123)
	assert.Equal(t, ErrRegisterNone, err)
}

// 2.3: Testing Stack Functions
////////////////////////////////

func TestPop(t *testing.T) {
	// setup
	Stack = US(1)

	// success
	u, err := Pop()
	assert.Equal(t, uint8(1), u)
	assert.NoError(t, err)

	// failure - ErrStackEmpty
	u, err = Pop()
	assert.Zero(t, u)
	assert.Equal(t, ErrStackEmpty, err)
}

func TestPopN(t *testing.T) {
	// setup
	Stack = US(1, 2)

	// success
	us, err := PopN(2)
	assert.Equal(t, US(2, 1), us)
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
	Push(1)
	assert.Equal(t, US(1), Stack)
}

func TestPushAll(t *testing.T) {
	// setup
	Stack = US()

	// success
	PushAll(US(1, 2))
	assert.Equal(t, US(1, 2), Stack)
}

///////////////////////////////////////////////////////////////////////////////////////
//                   Part 3: Testing Parsing & Evaluation Functions                  //
///////////////////////////////////////////////////////////////////////////////////////

// 3.1: Testing Parsing Functions
//////////////////////////////////

func TestClean(t *testing.T) {
	// success
	s := Clean("a // comment\nb c // comment\n")
	assert.Equal(t, "A\nB C\n", s)
}

func TestTokenise(t *testing.T) {
	// success
	ss := Tokenise("\t A  B  C \n")
	assert.Equal(t, SS("A", "B", "C"), ss)
}

// 3.2: Testing Evaluation Functions
/////////////////////////////////////

func TestAtomise(t *testing.T) {
	// setup
	Commands["CMD"] = func() error { return nil }
	Functions["FUN"] = []any{"A"}

	// success - uint8
	a, err := Atomise("1")
	assert.Equal(t, uint8(1), a)
	assert.NoError(t, err)

	// success - command
	a, err = Atomise("CMD")
	assert.NotNil(t, a)
	assert.NoError(t, err)

	// success - function
	a, err = Atomise("FUN")
	assert.Equal(t, AS("A"), a)
	assert.NoError(t, err)

	// failure - ErrSymbolNone
	a, err = Atomise("nope")
	assert.Nil(t, a)
	assert.Equal(t, ErrSymbolNone, err)
}

func TestEvaluate(t *testing.T) {
	// setup
	Stack = US()
	f := func() error {
		Push(1)
		return nil
	}

	// success - uint8
	err := Evaluate(uint8(1))
	assert.Equal(t, US(1), Stack)
	assert.NoError(t, err)

	// setup
	Stack = US()

	// success - command
	err = Evaluate(f)
	assert.Equal(t, US(1), Stack)
	assert.NoError(t, err)

	// setup
	Stack = US()

	// success - function
	err = Evaluate(AS(uint8(1)))
	assert.Equal(t, US(1), Stack)
	assert.NoError(t, err)

	// failure - ErrAtomUndefined
	err = Evaluate(false)
	assert.Equal(t, ErrAtomUndefined, err)
}

func TestEvaluateAll(t *testing.T) {
	// setup
	Stack = US()

	// success
	err := EvaluateAll(AS(uint8(1)))
	assert.Equal(t, US(1), Stack)
	assert.NoError(t, err)

	// failure - ErrAtomUndefined
	err = EvaluateAll(AS(false))
	assert.Equal(t, ErrAtomUndefined, err)
}

///////////////////////////////////////////////////////////////////////////////////////
//                         Part 4: Testing Command Functions                         //
///////////////////////////////////////////////////////////////////////////////////////

// 4.1: Testing Command Helper Functions
/////////////////////////////////////////

func TestBool(t *testing.T) {
	// success - true
	u := Bool(true)
	assert.Equal(t, uint8(1), u)

	// success - false
	u = Bool(false)
	assert.Equal(t, uint8(0), u)
}

// 4.2: Testing Integer Commands
/////////////////////////////////

func TestADD(t *testing.T) {
	// setup
	Stack = US(1, 2)

	// success
	err := ADD()
	assert.Equal(t, US(3), Stack)
	assert.NoError(t, err)
}

func TestSUB(t *testing.T) {
	// setup
	Stack = US(3, 2)

	// success
	err := SUB()
	assert.Equal(t, US(1), Stack)
	assert.NoError(t, err)
}

func TestMOD(t *testing.T) {
	// setup
	Stack = US(4, 3)

	// success
	err := MOD()
	assert.Equal(t, US(1), Stack)
	assert.NoError(t, err)
}

func TestGTE(t *testing.T) {
	// setup
	Stack = US(3, 2)

	// success - true
	err := GTE()
	assert.Equal(t, US(1), Stack)
	assert.NoError(t, err)

	// setup
	Stack = US(2, 3)

	// success - false
	err = GTE()
	assert.Equal(t, US(0), Stack)
	assert.NoError(t, err)
}

func TestLTE(t *testing.T) {
	// setup
	Stack = US(2, 3)

	// success - true
	err := LTE()
	assert.Equal(t, US(1), Stack)
	assert.NoError(t, err)

	// setup
	Stack = US(3, 2)

	// success - false
	err = LTE()
	assert.Equal(t, US(0), Stack)
	assert.NoError(t, err)
}

// 4.3: Memory Commands
////////////////////////

func TestCLR(t *testing.T) {
	// setup
	Stack = US(1, 2, 3)

	// success
	err := CLR()
	assert.Empty(t, Stack)
	assert.NoError(t, err)
}

func TestDUP(t *testing.T) {
	// setup
	Stack = US(1)

	// success
	err := DUP()
	assert.Equal(t, US(1, 1), Stack)
	assert.NoError(t, err)
}

func TestDRP(t *testing.T) {
	// setup
	Stack = US(1, 2)

	// success
	err := DRP()
	assert.Equal(t, US(1), Stack)
	assert.NoError(t, err)
}

func TestSWP(t *testing.T) {
	// setup
	Stack = US(1, 2)

	// success
	err := SWP()
	assert.Equal(t, US(2, 1), Stack)
	assert.NoError(t, err)
}

func TestGET(t *testing.T) {
	// setup
	Registers[0] = 123
	Stack = US(0)

	// success
	err := GET()
	assert.Equal(t, US(123), Stack)
	assert.NoError(t, err)
}

func TestSET(t *testing.T) {
	// setup
	Registers[0] = 0
	Stack = US(123, 0)

	// success
	err := SET()
	assert.Empty(t, Stack)
	assert.Equal(t, uint8(123), Registers[0])
	assert.NoError(t, err)
}
