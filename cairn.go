///////////////////////////////////////////////////////////////////////////////////////
//                   Cairn: A personal programming language in Go.                   //
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - //
//                            Version 0.0.0 (2024-03-05).                            //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"strconv"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////////////
//                              Part 1: Global Variables                             //
///////////////////////////////////////////////////////////////////////////////////////

// 1.1: Memory Variables
/////////////////////////

// Commands is a map of symbols to built-in command functions.
var Commands = make(map[string]func() error)

// Functions is a map of symbols to program-defined functions.
var Functions = make(map[string][]any)

// Queue is a first-in-first-out queue of parsed atoms.
var Queue = make([]any, 0, 63356)

// Registers is a fixed array of stored register values.
var Registers [8]uint8

// Stack is a last-in-first-out stack of stored memory values.
var Stack = make([]uint8, 0, 65536)

// 1.2: Error Definitions
//////////////////////////

// ErrQueueEmpty is the error for dequeuing an empty Queue.
var ErrQueueEmpty = errors.New("queue is empty")

// ErrStackEmpty is the error for popping an empty Stack.
var ErrStackEmpty = errors.New("stack is empty")

// ErrRegisterNone is the error for accessing a non-existent register.
var ErrRegisterNone = errors.New("register does not exist")

// ErrSymbolNone is the error for accessing a non-existent symbol
var ErrSymbolNone = errors.New("symbol does not exist")

// ErrAtomUndefined is the error for evaluating undefined atoms.
var ErrAtomUndefined = errors.New("atom type is not defined")

// 1.3: System Variables
/////////////////////////

// MainRun is a boolean indicating if the main loop should continue.
var MainRun = true

// VersionDate is the date of the current Cairn version.
var VersionDate = "2024-03-05"

// VersionNums is the SemVer number of the current Cairn version.
var VersionNums = "0.0.0"

///////////////////////////////////////////////////////////////////////////////////////
//                            Part 2: Collection Functions                           //
///////////////////////////////////////////////////////////////////////////////////////

// 2.1: Queue Functions
////////////////////////

// Dequeue removes and returns the first atom in the Queue.
func Dequeue() (any, error) {
	if len(Queue) == 0 {
		return nil, ErrQueueEmpty
	}

	a := Queue[0]
	Queue = Queue[1:]
	return a, nil
}

// DequeueTo removes and returns all atoms before an Atom in the Queue.
func DequeueTo(a any) ([]any, error) {
	var as []any
	for {
		a2, err := Dequeue()
		if err != nil {
			return nil, err
		}

		as = append(as, a2)
		if as[len(as)-1] == a {
			break
		}
	}

	return as[:len(as)-1], nil
}

// Enqueue appends an atom to the end of the Queue.
func Enqueue(a any) {
	Queue = append(Queue, a)
}

// EnqueueAll appends an atom slice to the end of the Queue.
func EnqueueAll(as []any) {
	Queue = append(Queue, as...)
}

// 2.2: Register Functions
///////////////////////////

// GetRegister returns the value of a given register.
func GetRegister(i uint8) (uint8, error) {
	if i > 7 {
		return 0, ErrRegisterNone
	}

	return Registers[int(i)], nil
}

// SetRegister sets the value of a given register.
func SetRegister(i, u uint8) error {
	if i > 7 {
		return ErrRegisterNone
	}

	Registers[int(i)] = u
	return nil
}

// 2.3: Stack Functions
////////////////////////

// Pop removes and returns the top item on the Stack.
func Pop() (uint8, error) {
	if len(Stack) == 0 {
		return 0, ErrStackEmpty
	}

	u := Stack[len(Stack)-1]
	Stack = Stack[:len(Stack)-1]
	return u, nil
}

// PopN removes and returns the top N items on the Stack.
func PopN(i int) ([]uint8, error) {
	var us []uint8
	for len(us) < i {
		u, err := Pop()
		if err != nil {
			return nil, err
		}

		us = append(us, u)
	}

	return us, nil
}

// Push appends an integer to the top of the Stack.
func Push(u uint8) {
	Stack = append(Stack, u)
}

// PushAll appends an integer slice to the top of the Stack.
func PushAll(us []uint8) {
	Stack = append(Stack, us...)
}

///////////////////////////////////////////////////////////////////////////////////////
//                       Part 3: Parsing & Evaluation Functions                      //
///////////////////////////////////////////////////////////////////////////////////////

// 3.1: Parsing Functions
//////////////////////////

// Clean returns an uppercase program string without comments.
func Clean(s string) string {
	var ss []string
	for _, s := range strings.Split(s, "\n") {
		s = strings.SplitN(s, "//", 2)[0]
		s = strings.TrimSpace(s)
		ss = append(ss, strings.ToUpper(s))
	}

	return strings.Join(ss, "\n")
}

// Tokenise returns a token slice from a clean program string.
func Tokenise(s string) []string {
	return strings.Fields(s)
}

// 3.2: Evaluation Functions
/////////////////////////////

// Atomise returns an atom from a token string.
func Atomise(s string) (any, error) {
	if u, err := strconv.ParseUint(s, 10, 8); err == nil {
		return uint8(u), nil
	}

	if c, ok := Commands[s]; ok {
		return c, nil
	}

	if as, ok := Functions[s]; ok {
		return as, nil
	}

	return nil, ErrSymbolNone
}

// Evaluate evaluates an atom.
func Evaluate(a any) error {
	switch a := a.(type) {
	case uint8:
		Push(a)
		return nil

	case func() error:
		return a()

	case []any:
		return EvaluateAll(a)

	default:
		return ErrAtomUndefined
	}
}

// EvaluateAll evaluates an atom slice.
func EvaluateAll(as []any) error {
	for _, a := range as {
		if err := Evaluate(a); err != nil {
			return err
		}
	}

	return nil
}

///////////////////////////////////////////////////////////////////////////////////////
//                             Part 4: Command Functions                             //
///////////////////////////////////////////////////////////////////////////////////////

// 4.1: Command Helper Functions
/////////////////////////////////

// Bool returns a boolean as an integer.
func Bool(b bool) uint8 {
	if b {
		return 1
	}

	return 0
}

// 4.2: Integer Commands
/////////////////////////

// ADD (a b > c) returns a + b.
func ADD() error {
	us, err := PopN(2)
	if err != nil {
		return err
	}

	Push(us[1] + us[0])
	return nil
}

// SUB (a b > c) returns a - b.
func SUB() error {
	us, err := PopN(2)
	if err != nil {
		return err
	}

	Push(us[1] - us[0])
	return nil
}

// MOD (a b > c) returns a % b.
func MOD() error {
	us, err := PopN(2)
	if err != nil {
		return err
	}

	Push(us[1] % us[0])
	return nil
}

// GTE (a b > c) returns a >= b.
func GTE() error {
	us, err := PopN(2)
	if err != nil {
		return err
	}

	Push(Bool(us[1] >= us[0]))
	return nil
}

// LTE (a b > c) returns a <= b.
func LTE() error {
	us, err := PopN(2)
	if err != nil {
		return err
	}

	Push(Bool(us[1] <= us[0]))
	return nil
}

// 4.3: Memory Commands
////////////////////////

// CLR (... > _) clears the stack.
func CLR() error {
	Stack = make([]uint8, 0, 65536)
	return nil
}

// DUP (a > a a) duplicates the top stack item.
func DUP() error {
	u, err := Pop()
	if err != nil {
		return err
	}

	PushAll([]uint8{u, u})
	return nil
}

// DRP (a b > a) deletes the top stack item.
func DRP() error {
	_, err := Pop()
	if err != nil {
		return err
	}

	return nil
}

// SWP (a b > b a) swaps the top two stack items.
func SWP() error {
	us, err := PopN(2)
	if err != nil {
		return err
	}

	PushAll([]uint8{us[0], us[1]})
	return nil
}

// GET (a > b) returns the value of register a.
func GET() error {
	u, err := Pop()
	if err != nil {
		return nil
	}

	u, err = GetRegister(u)
	if err != nil {
		return nil
	}

	Push(u)
	return nil
}

// SET (a b > _) sets a to register b.
func SET() error {
	us, err := PopN(2)
	if err != nil {
		return nil
	}

	return SetRegister(us[0], us[1])
}
