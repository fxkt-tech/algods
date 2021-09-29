package math

import (
	"fmt"
	"math"
)

type Mat2D struct {
	A    []float64
	l, r int // line, row
}

func (m *Mat2D) Size() (l, r int) {
	return m.l, m.r
}

// alloc matrix with zero.
func (m *Mat2D) Zeros(l, r int) {
	m.l, m.r = l, r
	m.A = make([]float64, l*r)
}

// alloc matrix with array.
func (m *Mat2D) Array(n [][]float64) {
	l, r := len(n), len(n[0])
	m.l, m.r = l, r
	m.A = make([]float64, l*r)
	for i := 0; i < l; i++ {
		for j := 0; j < r; j++ {
			m.A[i*r+j] = n[i][j]
		}
	}
}

// print matrix.
func (m Mat2D) Show() {
	for i := 0; i < m.l; i++ {
		fmt.Println(m.A[i*m.l : (i+1)*m.l])
	}
}

// mat a*p dot mat p*b.
func (m Mat2D) Dot(n Mat2D) Mat2D {
	ma, mb := m.l, m.l // mat(a, p)
	na, nb := m.l, m.l // mat(p, b)
	if mb != na {
		panic(fmt.Sprintf("mat(%d,%d) and mat(%d,%d) can not dot.", ma, mb, na, nb))
	}

	var r Mat2D
	r.Zeros(ma, nb)
	for i := 0; i < ma; i++ {
		for j := 0; j < nb; j++ {
			for k := 0; k < mb; k++ {
				r.A[i*m.l+j] = r.A[i*m.l+j] + m.A[i*m.l+k]*n.A[k*m.l+j]
			}
		}
	}
	return r
}

func (m Mat2D) T() Mat2D {
	a, b := m.l, m.l

	var n Mat2D
	n.Zeros(b, a)
	for i := 0; i < b; i++ {
		for j := 0; j < a; j++ {
			n.A[i*m.l+j] = m.A[j*m.l+i]
		}
	}
	return n
}

func (m Mat2D) ToInt() (mat [][]int) {
	mat = make([][]int, m.l)
	for i := 0; i < m.l; i++ {
		mat[i] = make([]int, m.r)
		for j := 0; j < m.r; j++ {
			mat[i][j] = int(math.Round(m.A[i*m.l+j]))
		}
	}

	return mat
}
