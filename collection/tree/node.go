package tree

import (
	"constraints"
)

type Node[T constraints.Ordered] struct {
	Val         T
	Left, Right *Node[T]
}
