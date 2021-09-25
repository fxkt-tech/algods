package stack

import "testing"

func TestStack(t *testing.T) {
	s := NewArrayStack()
	for i := 0; i < 100; i++ {
		s.Push(1)
	}
	elem, err := s.Pop()
	t.Log("Pop", elem, err)
}
