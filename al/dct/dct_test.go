package dct

import (
	"testing"

	"fxkt.tech/algods/math"
)

func TestFDCT(t *testing.T) {
	var m math.Mat2D
	m.Array([][]float64{
		{61, 19, 50, 20},
		{82, 26, 61, 45},
		{89, 90, 82, 43},
		{93, 59, 53, 97},
	})
	n := FDCT(m)
	n.Show()
}

func BenchmarkFDCT(b *testing.B) {
	b.StopTimer()
	var m math.Mat2D
	m.Array([][]float64{
		{61, 19, 50, 20},
		{82, 26, 61, 45},
		{89, 90, 82, 43},
		{93, 59, 53, 97},
	})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FDCT(m)
	}
}
