package node

import "fmt"

// 单向节点
type OneWay[T any] struct {
	Value T
	Next  *OneWay[T]
}

func Show[T any](n *OneWay[T]) {
	for {
		if n == nil {
			break
		}
		fmt.Printf("%v->", n.Value)
		n = n.Next
	}
	fmt.Println()
}

// 双向节点
type TwoWay[T any] struct {
	Value T
	Prev  *TwoWay[T]
	Next  *TwoWay[T]
}

// 二叉节点
type Binary[T any] struct {
	Value T
	Left  *Binary[T]
	Right *Binary[T]
}
