package list

import "fxkt.tech/ringo/node"

// func ReverseLinked[T any](v []T) {
// 	l := len(v)
// 	for i := 0; i < l/2; i++ {
// 		v[i], v[l-1-i] = v[l-1-i], v[i]
// 	}
// }

func NewLinkedList[T any]() *node.OneWay[T] {
	return &node.OneWay[T]{}
}

func Push[T any](n *node.OneWay[T], value T) {
	if n == nil {
		return
	}

}
