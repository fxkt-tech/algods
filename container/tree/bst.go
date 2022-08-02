package tree

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"

	"fxkt.tech/ringo/container/list"
	"fxkt.tech/ringo/container/queue"
	"fxkt.tech/ringo/container/stack"
	"fxkt.tech/ringo/node"
)

// 先序遍历
func PreOrder[T constraints.Ordered](root *node.Binary[T]) (l []T) {
	s := stack.NewArrayStack[*node.Binary[T]]()
	p := root
	for p != nil || !s.IsEmpty() {
		if p != nil {
			l = append(l, p.Value)
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
func InOrder[T constraints.Ordered](root *node.Binary[T]) (l []T) {
	s := stack.NewArrayStack[*node.Binary[T]]()
	p := root
	for p != nil || !s.IsEmpty() {
		if p != nil {
			s.Push(p)
			p = p.Left
		} else {
			n, _ := s.Pop()
			l = append(l, n.Value)
			p = n.Right
		}
	}
	return
}

// 后序遍历
func PostOrder[T constraints.Ordered](root *node.Binary[T]) (l []T) {
	s := stack.NewArrayStack[*node.Binary[T]]()
	p := root
	for p != nil || !s.IsEmpty() {
		if p != nil {
			l = append(l, p.Value)
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
func LevelOrder[T constraints.Ordered](root *node.Binary[T]) (l [][]T) {
	q := queue.NewArrayQueue[*node.Binary[T]]()
	q.Push(root)
	for sz := q.Size(); sz > 0; sz = q.Size() {
		var x []T
		for i := 0; i < sz; i++ {
			elem, _ := q.Pop()
			qn := elem
			x = append(x, qn.Value)
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

func Marshal[T constraints.Ordered](root *node.Binary[T]) string {
	if root == nil {
		return "[]"
	}
	var x []string
	q := queue.NewArrayQueue[*node.Binary[T]]()
	q.Push(root)
	for sz := q.Size(); sz > 0; sz = q.Size() {
		for i := 0; i < sz; i++ {
			qn, _ := q.Pop()
			if qn != nil {
				x = append(x, fmt.Sprintf("%v", qn.Value))
				q.Push(qn.Left)
				q.Push(qn.Right)
			} else {
				x = append(x, "null")
			}
		}
	}
	return fmt.Sprintf("[%s]", strings.Join(x, ","))
}
