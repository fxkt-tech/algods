package stack

import (
	"errors"
	"fmt"
)

type ArrayStack struct {
	arr []int // abstract stack.
	len int   // current element length in stack.
	cap int   // current arr cap.
}

func NewArrayStack() *ArrayStack {
	var cap int = 8
	return &ArrayStack{
		arr: make([]int, cap),
		cap: cap,
	}
}

func (s *ArrayStack) Push(val int) {
	if s.cap <= s.len {
		s.expansion()
	}
	s.arr[s.len] = val
	s.len = s.len + 1
}

func (s *ArrayStack) Pop() (int, error) {
	if s.len <= 0 {
		return 0, errors.New("stack is empty.")
	}
	elem := s.arr[s.len-1]
	s.len = s.len - 1
	return elem, nil
}

func (s *ArrayStack) expansion() {
	newarr := make([]int, s.cap*2)
	newarr = append(newarr, s.arr...)
	s.arr = newarr
	s.cap = cap(newarr)
	fmt.Println(s.cap)
}
