package stack

type ArrayStack[T any] struct {
	elems []T // abstract stack.
	nelem int // number of element.
}

func NewArrayStack[T any]() Stack[T] {
	return &ArrayStack[T]{}
}

func (s *ArrayStack[T]) Push(val T) {
	s.elems = append(s.elems, val)
	s.nelem = s.nelem + 1
}

func (s *ArrayStack[T]) Pop() (t T, b bool) {
	if s.nelem == 0 {
		return
	}
	elem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return elem, true
}

func (s *ArrayStack[T]) IsEmpty() bool {
	return s.nelem == 0
}
