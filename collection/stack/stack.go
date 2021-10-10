package stack

type ArrayStack struct {
	arr []interface{} // abstract stack.
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{}
}

func (s *ArrayStack) Push(val interface{}) {
	s.arr = append(s.arr, val)
}

func (s *ArrayStack) Pop() (interface{}, bool) {
	if len(s.arr) == 0 {
		return 0, false
	}
	elem := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return elem, true
}

func (s *ArrayStack) IsEmpty() bool {
	return len(s.arr) == 0
}
