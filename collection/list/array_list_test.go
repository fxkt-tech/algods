package list

import "testing"

func TestReverseArray(t *testing.T) {
	v := []int{
		1, 2, 3,
	}
	ReverseArray(v)
	t.Log(v)
}
