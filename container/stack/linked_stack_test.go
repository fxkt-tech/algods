package stack

import "testing"

func TestLinkedStack(t *testing.T) {
	s := NewLinkedStack[int]()
	for i := 0; i < 100; i++ {
		s.Push(i)
	}
	for i := 0; i < 100; i++ {
		elem, err := s.Pop()
		t.Log("Pop", elem, err)
	}
}
