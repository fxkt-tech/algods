package math

import (
	"fmt"
	"testing"
)

func TestDot(t *testing.T) {
	var m Mat2D
	m.Array([][]float64{
		{1, 2, 3},
		{3, 4, 5},
	})
	m.Show()
	fmt.Println()

	var n Mat2D
	n.Array([][]float64{
		{1, 2},
		{3, 4},
		{3, 4},
	})
	n.Show()
	fmt.Println()

	r := m.Dot(n)
	r.Show()
}

func TestMatT(t *testing.T) {
	var m Mat2D
	m.Array([][]float64{
		{1, 2, 3},
		{3, 4, 5},
	})
	m.Show()
	fmt.Println()

	n := m.T()
	n.Show()
}
