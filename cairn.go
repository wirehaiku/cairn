///////////////////////////////////////////////////////////////////////////////////////
//                   Cairn: A personal programming language in Go.                   //
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - //
//                            Version 0.0.0 (2024-03-05).                            //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
)

///////////////////////////////////////////////////////////////////////////////////////
//                              Part 1: Global Variables                             //
///////////////////////////////////////////////////////////////////////////////////////

// 1.1: Memory Variables
/////////////////////////

// Commands is a map of symbols to built-in command functions.
var Commands map[string]func() error

// Functions is a map of symbols to program-defined functions.
var Functions map[string][]any

// Queue is a first-in-first-out queue of parsed program atoms.
var Queue []any

// Registers is a fixed array of stored register values.
var Registers [8]uint8

// Stack is a last-in-first-out stack of stored memory values.
var Stack []uint8

// 1.2: Error Definitions
//////////////////////////

// ErrQueueEmpty is the error for dequeuing the empty Queue.
var ErrQueueEmpty = errors.New("queue empty")

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
