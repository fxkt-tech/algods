package cg

import (
	sysmath "math"

	"fxkt.tech/algods/math"
)

// Forward Discrete Cosine Transform for 2d matrix.
func FDCT2D(b math.Mat2D) math.Mat2D {
	var (
		l, _ int = b.Size()
		a    math.Mat2D
	)
	a.Zeros(l, l)

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			var c float64
			if i == 0 {
				c = sysmath.Sqrt(1 / float64(l))
			} else {
				c = sysmath.Sqrt(2 / float64(l))
			}
			a.A[i*l+j] = c * sysmath.Cos((float64(j)+0.5)*sysmath.Pi*float64(i)/float64(l))
		}
	}
	return a.Dot(b).Dot(a.T())
}

// Inverse Discrete Cosine Transform for 2d matrix.
func IFDCT2D(b math.Mat2D) math.Mat2D {
	var (
		l, _ int = b.Size()
		a    math.Mat2D
	)
	a.Zeros(l, l)

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			var c float64
			if i == 0 {
				c = sysmath.Sqrt(1 / float64(l))
			} else {
				c = sysmath.Sqrt(2 / float64(l))
			}
			a.A[i*l+j] = c * sysmath.Cos((float64(j)+0.5)*sysmath.Pi*float64(i)/float64(l))
		}
	}
	return a.T().Dot(b).Dot(a)
}
