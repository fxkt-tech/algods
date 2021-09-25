package dct

import (
	sysmath "math"

	"fxkt.tech/algods/math"
)

func FDCT(b math.Mat2D) math.Mat2D {
	var sidelen int = len(b)
	var a math.Mat2D
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
