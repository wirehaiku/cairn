///////////////////////////////////////////////////////////////////////////////////////
//                                  Cairn Unit Tests                                 //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"bytes"
	"os"
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

// Bufs sets Stdin and Stdout to mock buffers, populates Stdin with a string and
// returns the Stdout buffer.
func Bufs(s string) *bytes.Buffer {
	inn := bytes.NewBufferString(s)
	out := bytes.NewBuffer(nil)
	Stdin = bufio.NewReader(inn)
	Stdout = bufio.NewWriter(out)
	return out
}

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

func TestDump(t *testing.T) {
	// setup
	Stack = US(1, 2, 3)

	// success
	s := Dump()
	assert.Equal(t, "1 2 3", s)
}

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

func TestAtomiseAll(t *testing.T) {
	// success
	as, err := AtomiseAll(SS("1", "2", "3"))
	assert.Equal(t, AS(uint8(1), uint8(2), uint8(3)), as)
	assert.NoError(t, err)
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
//                       Part 4: Testing Input/Output Functions                      //
///////////////////////////////////////////////////////////////////////////////////////

// 4.1: Testing Standard IO Functions
//////////////////////////////////////

func TestInput(t *testing.T) {
	// setup
	Bufs("A")

	// success
	u, err := Input()
	assert.Equal(t, uint8(65), u)
	assert.NoError(t, err)
}

func TestOutput(t *testing.T) {
	// setup
	b := Bufs("")

	// success
	err := Output(65)
	assert.Equal(t, "A", b.String())
	assert.NoError(t, err)
}

// 4.2: Testing Command-Line Functions
///////////////////////////////////////

func TestParseFlags(t *testing.T) {
	// setup
	ss := []string{"-c", "cmd", "-d"}

	// success
	fs, err := ParseFlags(ss)
	assert.Equal(t, "cmd", fs.Command)
	assert.Equal(t, true, fs.Debug)
	assert.NoError(t, err)
}

///////////////////////////////////////////////////////////////////////////////////////
//                         Part 5: Testing Command Functions                         //
///////////////////////////////////////////////////////////////////////////////////////

// 5.1: Testing Command Helper Functions
/////////////////////////////////////////

func TestBool(t *testing.T) {
	// success - true
	u := Bool(true)
	assert.Equal(t, uint8(1), u)

	// success - false
	u = Bool(false)
	assert.Equal(t, uint8(0), u)
}

// 5.2: Testing Integer Commands
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

// 5.3: Memory Commands
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

// 5.4: Testing Logic Commands
///////////////////////////////

func TestEQU(t *testing.T) {
	// setup
	Stack = US(2, 2)

	// success - true
	err := EQU()
	assert.Equal(t, US(1), Stack)
	assert.NoError(t, err)

	// setup
	Stack = US(3, 2)

	// success - false
	err = EQU()
	assert.Equal(t, US(0), Stack)
	assert.NoError(t, err)
}

func TestNEQ(t *testing.T) {
	// setup
	Stack = US(3, 2)

	// success - true
	err := NEQ()
	assert.Equal(t, US(1), Stack)
	assert.NoError(t, err)

	// setup
	Stack = US(2, 2)

	// success - false
	err = NEQ()
	assert.Equal(t, US(0), Stack)
	assert.NoError(t, err)
}

func TestAND(t *testing.T) {
	// setup
	Stack = US(1, 1)

	// success
	err := AND()
	assert.Equal(t, US(1), Stack)
	assert.NoError(t, err)

	// setup
	Stack = US(1, 0)

	// failure - a false
	err = AND()
	assert.Equal(t, US(0), Stack)
	assert.NoError(t, err)

	// setup
	Stack = US(0, 1)

	// failure - b false
	err = AND()
	assert.Equal(t, US(0), Stack)
	assert.NoError(t, err)
}

func TestORR(t *testing.T) {
	// setup
	Stack = US(0, 1)

	// success - a true
	err := ORR()
	assert.Equal(t, US(1), Stack)
	assert.NoError(t, err)

	// setup
	Stack = US(1, 0)

	// success - b true
	err = ORR()
	assert.Equal(t, US(1), Stack)
	assert.NoError(t, err)

	// setup
	Stack = US(1, 1)

	// success - both true
	err = ORR()
	assert.Equal(t, US(1), Stack)
	assert.NoError(t, err)

	// setup
	Stack = US(0, 0)

	// success - false
	err = ORR()
	assert.Equal(t, US(0), Stack)
	assert.NoError(t, err)
}

func TestXOR(t *testing.T) {
	// setup
	Stack = US(0, 1)

	// success - a true
	err := XOR()
	assert.Equal(t, US(1), Stack)
	assert.NoError(t, err)

	// setup
	Stack = US(1, 0)

	// success - b true
	err = XOR()
	assert.Equal(t, US(1), Stack)
	assert.NoError(t, err)

	// setup
	Stack = US(1, 1)

	// success - exclusive false
	err = XOR()
	assert.Equal(t, US(0), Stack)
	assert.NoError(t, err)

	// setup
	Stack = US(0, 0)

	// success - false
	err = XOR()
	assert.Equal(t, US(0), Stack)
	assert.NoError(t, err)
}

func TestNOT(t *testing.T) {
	// setup
	Stack = US(0)

	// success - true
	err := NOT()
	assert.Equal(t, US(1), Stack)
	assert.NoError(t, err)

	// setup
	Stack = US(1)

	// success - false
	err = NOT()
	assert.Equal(t, US(0), Stack)
	assert.NoError(t, err)
}

// 5.5: Testing System Commands
////////////////////////////////

func TestINN(t *testing.T) {
	// setup
	Stack = US()
	Bufs("A")

	// success
	err := INN()
	assert.Equal(t, US(65), Stack)
	assert.NoError(t, err)
}

func TestOUT(t *testing.T) {
	// setup
	Stack = US(65)
	b := Bufs("")

	// success
	err := OUT()
	assert.Empty(t, Stack)
	assert.Equal(t, "A", b.String())
	assert.NoError(t, err)
}

func TestBYE(t *testing.T) {
	// setup
	var x int
	ExitFunc = func(i int) { x = i }

	// success
	err := BYE()
	assert.Equal(t, 0, x)
	assert.NoError(t, err)
}

func TestDIE(t *testing.T) {
	// setup
	var x int
	Stack = US(123)
	ExitFunc = func(i int) { x = i }

	// success
	err := DIE()
	assert.Equal(t, 123, x)
	assert.NoError(t, err)
}

// 5.6: Testing Flow Control Commands
//////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////
//                          Part 6: Testing Main Functions                           //
///////////////////////////////////////////////////////////////////////////////////////

// 6.1: Testing Main Helper Functions
//////////////////////////////////////

func TestPrint(t *testing.T) {
	// setup
	b := Bufs("")

	// success
	print("%s", "test")
	assert.Equal(t, "test", b.String())
}

func TestPrompt(t *testing.T) {
	// setup
	b := Bufs("test\n")

	// success
	s := prompt(">>> ")
	assert.Equal(t, ">>> ", b.String())
	assert.Equal(t, "test\n", s)
}

func TestOnce(t *testing.T) {
	// setup
	Commands["ADD"] = ADD
	Queue = AS()
	Stack = US()

	// success
	err := once("1 2 ADD")
	assert.Empty(t, Queue)
	assert.Equal(t, US(3), Stack)
	assert.NoError(t, err)
}

// 6.2: Testing Main Boot Functions
////////////////////////////////////

func TestInit(t *testing.T) {
	// success
	assert.NotZero(t, len(Commands))
}

func TestMain(t *testing.T) {
	// setup
	b := Bufs("")
	os.Args = SS("cairn", "-c", "72 OUT 73 OUT")

	// success - command string
	main()
	assert.Equal(t, "HI", b.String())

	// setup
	b = Bufs("1 2 ADD\nBYE\n")
	os.Args = SS("cairn")

	// success - REPL
	main()
	assert.Equal(t, ">>> [ 3 ]\n>>> ", b.String())
}
