package tree

import (
	"testing"
)

func TestBST(t *testing.T) {
	bst := NewBinarySearchTree()
	node := &Node{
		Val: 4,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 1},
			Right: &Node{Val: 3},
		},
		Right: &Node{
			Val:   6,
			Left:  &Node{Val: 5},
			Right: &Node{Val: 7},
		},
	}
	bst.PreOrderByRecursion(node)
	t.Log("PreOder:", bst.LastResult)

	bst.Clear()

	bst.InOrderByRecursion(node)
	t.Log("InOder:", bst.LastResult)

	bst.Clear()

	bst.PostOrderByRecursion(node)
	t.Log("PostOder:", bst.LastResult)
}
