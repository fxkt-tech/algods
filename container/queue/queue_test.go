package queue

import "testing"

func TestArrayQueue(t *testing.T) {
	s := NewArrayQueue[int]()
	for i := 0; i < 100; i++ {
		s.Push(i)
	}
	for i := 0; i < 100; i++ {
		elem, err := s.Pop()
		t.Log("Pop", elem, err)
	}
}
