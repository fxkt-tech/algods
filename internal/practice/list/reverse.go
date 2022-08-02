package list

import (
	"fxkt.tech/ringo/node"
)

func Reverse[T any](n *node.OneWay[T]) *node.OneWay[T] {
	var (
		prev *node.OneWay[T]
		cur  = n
	)
	for cur != nil {
		nxt := cur.Next // 保存当前节点的后继
		cur.Next = prev // 修改当前节点的后继为之前的前驱
		prev = cur      // 设置当前值为前驱
		cur = nxt       // 设置上面保存的后继为当前节点

	}
	return prev
}
