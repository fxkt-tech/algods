package math

import (
	"fmt"
)

type Mat2D [][]float64

// alloc matrix with zero.
func (m *Mat2D) Zeros(w, h int) {
	*m = make(Mat2D, w)
	for i := 0; i < w; i++ {
		(*m)[i] = make([]float64, h)
	}
}

// alloc matrix with array.
func (m *Mat2D) Array(n Mat2D) {
	*m = n
}

// print matrix.
func (m Mat2D) Show() {
	for i := 0; i < len(m); i++ {
		fmt.Println(m[i])
	}
}

// mat a*p dot mat p*b.
func (m Mat2D) Dot(n Mat2D) Mat2D {
	ma, mb := len(m), len(m[0]) // mat(a, p)
	na, nb := len(n), len(n[0]) // mat(p, b)
	if mb != na {
		panic(fmt.Sprintf("mat(%d,%d) and mat(%d,%d) can not dot.", ma, mb, na, nb))
	}
	// fmt.Printf("mat(%d,%d) and mat(%d,%d) can dot.\n", ma, mb, na, nb)

	var r Mat2D
	r.Zeros(ma, nb)
	for i := 0; i < ma; i++ {
		for j := 0; j < nb; j++ {
			for k := 0; k < mb; k++ {
				r[i][j] = r[i][j] + m[i][k]*n[k][j]
			}
		}
	}
	return r
}

func (m Mat2D) T() Mat2D {
	a, b := len(m), len(m[0])

	var n Mat2D
	n.Zeros(b, a)
	for i := 0; i < b; i++ {
		for j := 0; j < a; j++ {
			n[i][j] = m[j][i]
		}
	}
	return n
}
