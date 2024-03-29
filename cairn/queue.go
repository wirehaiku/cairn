package cairn

import "fmt"

// Queue is a first-in-first-out queue of atoms.
type Queue struct {
	Atoms []any
}

// NewQueue returns a pointer to a new Queue.
func NewQueue(as ...any) *Queue {
	return &Queue{as}
}

// Clear removes all atoms from the Queue.
func (q *Queue) Clear() {
	q.Atoms = make([]any, 0)
}

// Dequeue removes and returns the first atom in the Queue.
func (q *Queue) Dequeue() (any, error) {
	if len(q.Atoms) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}

	a := q.Atoms[0]
	q.Atoms = q.Atoms[1:]
	return a, nil
}

// Empty returns true if the Queue has no atoms.
func (q *Queue) Empty() bool {
	return len(q.Atoms) == 0
}

// Enqueue appends an atom to the end of the Queue.
func (q *Queue) Enqueue(a any) {
	q.Atoms = append(q.Atoms, a)
}

// EnqueueAll appends an atom slice to the end of the Queue.
func (q *Queue) EnqueueAll(as []any) {
	q.Atoms = append(q.Atoms, as...)
}

// Len returns the number of atoms in the Queue.
func (q *Queue) Len() any {
	return len(q.Atoms)
}
