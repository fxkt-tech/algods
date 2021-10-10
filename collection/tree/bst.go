package tree

import (
	"fmt"
	"strconv"
	"strings"

	"fxkt.tech/algods/collection/queue"
	"fxkt.tech/algods/collection/stack"
	"fxkt.tech/algods/collection/vector"
	"fxkt.tech/algods/pkg/json"
)

func PreOrder(root *Node) (l []int) {
	s := stack.NewArrayStack()
	p := root
	for p != nil || !s.IsEmpty() {
		if p != nil {
			l = append(l, p.Val)
			s.Push(p)
			p = p.Left
		} else {
			n, _ := s.Pop()
			p = n.(*Node).Right
		}
	}
	return
}

func InOrder(root *Node) (l []int) {
	s := stack.NewArrayStack()
	p := root
	for p != nil || !s.IsEmpty() {
		if p != nil {
			s.Push(p)
			p = p.Left
		} else {
			n, _ := s.Pop()
			l = append(l, n.(*Node).Val)
			p = n.(*Node).Right
		}
	}
	return
}

func PostOrder(root *Node) (l []int) {
	s := stack.NewArrayStack()
	p := root
	for p != nil || !s.IsEmpty() {
		if p != nil {
			l = append(l, p.Val)
			s.Push(p)
			p = p.Right
		} else {
			n, _ := s.Pop()
			p = n.(*Node).Left
		}
	}
	vector.Reverse(l)
	return
}

func LevelOrder(root *Node) (l [][]int) {
	q := queue.NewArrayQueue()
	q.Push(root)
	for sz := q.Size(); sz > 0; sz = q.Size() {
		var x []int
		for i := 0; i < sz; i++ {
			elem, _ := q.Pop()
			qn := elem.(*Node)
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

func Marshal(root *Node) string {
	if root == nil {
		return "[]"
	}
	var x []string
	var q []*Node
	q = append(q, root)
	for sz := len(q); sz > 0; sz = len(q) {
		for i := 0; i < sz; i++ {
			fmt.Println("q:", json.ToString(q))
			qn := q[0]
			q = q[1:]
			fmt.Println("q2:", json.ToString(q))
			if qn != nil {
				x = append(x, strconv.Itoa(qn.Val))
				q = append(q, qn.Left)
				q = append(q, qn.Right)
			} else {
				x = append(x, "null")
			}
			fmt.Println("q3:", json.ToString(q))
		}
	}
	return fmt.Sprintf("[%s]", strings.Join(x, ","))
}
