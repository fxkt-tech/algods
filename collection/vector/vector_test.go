package vector

import "testing"

func TestReverse(t *testing.T) {
	v := []int{
		1, 2, 3,
	}
	Reverse(v)
	t.Log(v)
}
