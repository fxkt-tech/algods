package tree

import (
	"testing"

	"fxkt.tech/algods/node"
)

func TestBST(t *testing.T) {
	bst := NewBinarySearchTree()
	node := &node.Tree{
		Val: 4,
		Left: &node.Tree{
			Val:   2,
			Left:  &node.Tree{Val: 1},
			Right: &node.Tree{Val: 3},
		},
		Right: &node.Tree{
			Val:   6,
			Left:  &node.Tree{Val: 5},
			Right: &node.Tree{Val: 7},
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
