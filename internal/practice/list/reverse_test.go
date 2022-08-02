package list

import (
	"testing"

	"fxkt.tech/ringo/node"
)

func TestReverse(t *testing.T) {
	list := &node.OneWay[int]{
		Value: 1,
		Next: &node.OneWay[int]{
			Value: 2,
			Next: &node.OneWay[int]{
				Value: 3,
			},
		},
	}
	node.Show(list)
	rlist := Reverse(list)
	node.Show(rlist)
}
