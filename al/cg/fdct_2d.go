package cg

import (
	sysmath "math"

	"fxkt.tech/algods/math"
)

func FDCT2D(b math.Mat2D) math.Mat2D {
	var (
		sidelen int = len(b)
		a       math.Mat2D
	)
	a.Zeros(sidelen, sidelen)

	for i := 0; i < sidelen; i++ {
		for j := 0; j < sidelen; j++ {
			var c float64
			if i == 0 {
				c = sysmath.Sqrt(1 / float64(sidelen))
			} else {
				c = sysmath.Sqrt(2 / float64(sidelen))
			}
			a[i][j] = c * sysmath.Cos((float64(j)+0.5)*sysmath.Pi*float64(i)/float64(sidelen))
		}
	}
	return a.Dot(b).Dot(a.T())
}

func FDCT2DNew(b math.Mat2DNew) math.Mat2DNew {
	var (
		sidelen int = int(sysmath.Sqrt(float64(len(b))))
		a       math.Mat2DNew
	)
	a.Zeros(sidelen * sidelen)

	for i := 0; i < sidelen; i++ {
		for j := 0; j < sidelen; j++ {
			var c float64
			if i == 0 {
				c = sysmath.Sqrt(1 / float64(sidelen))
			} else {
				c = sysmath.Sqrt(2 / float64(sidelen))
			}
			a[i*sidelen+j] = c * sysmath.Cos((float64(j)+0.5)*sysmath.Pi*float64(i)/float64(sidelen))
		}
	}
	return a.Dot(b).Dot(a.T())
}
