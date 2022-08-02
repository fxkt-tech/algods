package tree

import (
	"testing"

	"fxkt.tech/ringo/node"
)

func TestBST(t *testing.T) {
	te := &node.Binary[int]{
		Value: 4,
		Left: &node.Binary[int]{
			Value: 2,
			Left:  &node.Binary[int]{Value: 1},
			Right: &node.Binary[int]{Value: 3},
		},
		Right: &node.Binary[int]{
			Value: 6,
			Left:  &node.Binary[int]{Value: 5},
			Right: &node.Binary[int]{Value: 7},
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
