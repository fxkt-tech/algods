package stack

type ArrayStack[T any] struct {
	arr []T // abstract stack.
}

func NewArrayStack[T any]() *ArrayStack[T] {
	return &ArrayStack[T]{}
}

func (s *ArrayStack[T]) Push(val T) {
	s.arr = append(s.arr, val)
}

func (s *ArrayStack[T]) Pop() (t T, b bool) {
	if len(s.arr) == 0 {
		return
	}
	elem := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return elem, true
}

func (s *ArrayStack[T]) IsEmpty() bool {
	return len(s.arr) == 0
}
