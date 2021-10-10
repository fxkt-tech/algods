package tree

import (
	"testing"
)

func TestBST(t *testing.T) {
	te := &Node{
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

	l := PreOrder(te)
	t.Log("PreOrder:", l)

	l = InOrder(te)
	t.Log("InOrder:", l)

	l = PostOrder(te)
	t.Log("PostOrder:", l)

	l2 := LevelOrder(te)
	t.Log("LevelOrder:", l2)

	s := Marshal(te)
	t.Log("Marshal:", s)
}
