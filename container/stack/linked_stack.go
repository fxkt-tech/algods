package stack

import "fxkt.tech/ringo/node"

type LinkedStack[T any] struct {
	elems *node.OneWay[T] // abstract stack.
	nelem int             // number of element.
}

func NewLinkedStack[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{}
}

func (s *LinkedStack[T]) Push(val T) {
	s.elems = &node.OneWay[T]{Value: val, Next: s.elems}
	s.nelem = s.nelem + 1
}

func (s *LinkedStack[T]) Pop() (t T, b bool) {
	if s.nelem == 0 {
		return
	}
	elem := s.elems
	s.elems = s.elems.Next
	s.nelem = s.nelem - 1
	return elem.Value, true
}

func (s *LinkedStack[T]) IsEmpty() bool {
	return s.nelem == 0
}
