package cairn

import "fmt"

// Stack is a last-in-first-out stack of integers.
type Stack struct {
	Integers []int
}

// NewStack returns a pointer to a new Stack.
func NewStack(is ...int) *Stack {
	return &Stack{is}
}

// Clear removes all integers from the Stack.
func (s *Stack) Clear() {
	s.Integers = make([]int, 0)
}

// Empty returns true if the Stack has no integers.
func (s *Stack) Empty() bool {
	return len(s.Integers) == 0
}

// Len returns the number of integers on the Stack.
func (s *Stack) Len() int {
	return len(s.Integers)
}

// Pop removes and returns the top integer on the Stack.
func (s *Stack) Pop() (int, error) {
	if len(s.Integers) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	i := s.Integers[len(s.Integers)-1]
	s.Integers = s.Integers[:len(s.Integers)-1]
	return i, nil
}

// PopN removes and returns the top N integers on the Stack.
func (s *Stack) PopN(n int) ([]int, error) {
	var is []int
	for len(is) < n {
		i, err := s.Pop()
		if err != nil {
			return nil, err
		}

		is = append(is, i)
	}

	return is, nil
}

// Push appends an integer to the top of the Stack.
func (s *Stack) Push(i int) {
	s.Integers = append(s.Integers, i)
}

// PushAll appends an integer slice to the top of the Stack.
func (s *Stack) PushAll(is []int) {
	s.Integers = append(s.Integers, is...)
}
