package list

import "constraints"

type Node[T constraints.Ordered] struct {
	Val  T
	Next *Node[T]
}
