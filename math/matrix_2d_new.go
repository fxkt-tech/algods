package math

import (
	"fmt"
	"math"
)

type Mat2DNew []float64

// alloc matrix with zero.
func (m *Mat2DNew) Zeros(l int) {
	*m = make(Mat2DNew, l)
}

// alloc matrix with array.
func (m *Mat2DNew) Array(n Mat2D) {
	a, b := len(n), len(n[0])
	*m = make(Mat2DNew, a*b)
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			(*m)[i*b+j] = n[i][j]
		}
	}
}

// print matrix.
func (m Mat2DNew) Show() {
	sidelen := int(math.Sqrt(float64(len(m))))
	for i := 0; i < sidelen; i++ {
		fmt.Println(m[i*sidelen : (i+1)*sidelen])
	}
}

// mat a*p dot mat p*b.
func (m Mat2DNew) Dot(n Mat2DNew) Mat2DNew {
	sidelen := int(math.Sqrt(float64(len(m))))
	ma, mb := sidelen, sidelen // mat(a, p)
	na, nb := sidelen, sidelen // mat(p, b)
	if mb != na {
		panic(fmt.Sprintf("mat(%d,%d) and mat(%d,%d) can not dot.", ma, mb, na, nb))
	}

	var r Mat2DNew
	r.Zeros(ma * nb)
	for i := 0; i < ma; i++ {
		for j := 0; j < nb; j++ {
			for k := 0; k < mb; k++ {
				r[i*sidelen+j] = r[i*sidelen+j] + m[i*sidelen+k]*n[k*sidelen+j]
			}
		}
	}
	return r
}

func (m Mat2DNew) T() Mat2DNew {
	sidelen := int(math.Sqrt(float64(len(m))))
	a, b := sidelen, sidelen

	var n Mat2DNew
	n.Zeros(b * a)
	for i := 0; i < b; i++ {
		for j := 0; j < a; j++ {
			n[i*sidelen+j] = m[j*sidelen+i]
		}
	}
	return n
}
