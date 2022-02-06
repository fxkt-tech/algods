package tree

import (
	"constraints"
	"fmt"
	"strings"

	"fxkt.tech/ringo/collection/list"
	"fxkt.tech/ringo/collection/queue"
	"fxkt.tech/ringo/collection/stack"
)

// 先序遍历
func PreOrder[T constraints.Ordered](root *Node[T]) (l []T) {
	s := stack.NewArrayStack[*Node[T]]()
	p := root
	for p != nil || !s.IsEmpty() {
		if p != nil {
			l = append(l, p.Val)
			s.Push(p)
			p = p.Left
		} else {
			n, _ := s.Pop()
			p = n.Right
		}
	}
	return
}

// 中序遍历
func InOrder[T constraints.Ordered](root *Node[T]) (l []T) {
	s := stack.NewArrayStack[*Node[T]]()
	p := root
	for p != nil || !s.IsEmpty() {
		if p != nil {
			s.Push(p)
			p = p.Left
		} else {
			n, _ := s.Pop()
			l = append(l, n.Val)
			p = n.Right
		}
	}
	return
}

// 后序遍历
func PostOrder[T constraints.Ordered](root *Node[T]) (l []T) {
	s := stack.NewArrayStack[*Node[T]]()
	p := root
	for p != nil || !s.IsEmpty() {
		if p != nil {
			l = append(l, p.Val)
			s.Push(p)
			p = p.Right
		} else {
			n, _ := s.Pop()
			p = n.Left
		}
	}
	list.ReverseArray(l)
	return
}

// 层序遍历
func LevelOrder[T constraints.Ordered](root *Node[T]) (l [][]T) {
	q := queue.NewArrayQueue[*Node[T]]()
	q.Push(root)
	for sz := q.Size(); sz > 0; sz = q.Size() {
		var x []T
		for i := 0; i < sz; i++ {
			elem, _ := q.Pop()
			qn := elem
			x = append(x, qn.Val)
			if qn.Left != nil {
				q.Push(qn.Left)
			}
			if qn.Right != nil {
				q.Push(qn.Right)
			}
		}
		l = append(l, x)
	}
	return
}

func Marshal[T constraints.Ordered](root *Node[T]) string {
	if root == nil {
		return "[]"
	}
	var x []string
	q := queue.NewArrayQueue[*Node[T]]()
	q.Push(root)
	for sz := q.Size(); sz > 0; sz = q.Size() {
		for i := 0; i < sz; i++ {
			qn, _ := q.Pop()
			if qn != nil {
				x = append(x, fmt.Sprintf("%v", qn.Val))
				q.Push(qn.Left)
				q.Push(qn.Right)
			} else {
				x = append(x, "null")
			}
		}
	}
	return fmt.Sprintf("[%s]", strings.Join(x, ","))
}
