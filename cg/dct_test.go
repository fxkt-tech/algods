package cg

import (
	"fmt"
	"testing"

	"fxkt.tech/ringo/math"
)

var (
	fdct_2d_testdata = [][]float64{
		{61, 19, 50, 20},
		{82, 26, 61, 45},
		{89, 90, 82, 43},
		{93, 59, 53, 97},
	}
	idct_2d_testdata = [][]float64{
		{242.5, 32.16133964439072, 22.5, 33.22120154777878},
		{-61.82630491859174, 7.924621202458745, -10.73436691753627, 30.688077169749352},
		{-16.500000000000004, -14.754876046005235, 22.500000000000004, -6.877036634119447},
		{8.832214904162955, 16.68807716974934, -35.06099494993978, -6.924621202458751},
	}
	fdct_2d_8s_testdata = [][]float64{
		{0x7f, 0xf6, 0x01, 0x07, 0xff, 0x00, 0x00, 0x00},
		{0xf5, 0x01, 0xfa, 0x01, 0xfe, 0x00, 0x01, 0x00},
		{0x05, 0x05, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00},
		{0x01, 0xff, 0xf8, 0x00, 0x01, 0xff, 0x00, 0x00},
		{0x00, 0x01, 0x00, 0x01, 0x00, 0xff, 0xff, 0x00},
		{0xff, 0x0c, 0x00, 0x00, 0x00, 0x00, 0xff, 0x01},
		{0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00},
		{0x01, 0x00, 0x00, 0x01, 0xff, 0x01, 0x00, 0xfe},
	}
)

func TestFDCT2D(t *testing.T) {
	var m math.Mat2D
	m.Array(fdct_2d_testdata)
	n := FDCT2D(m)
	n.Show()
}

func BenchmarkFDCT2D(b *testing.B) {
	b.StopTimer()
	var m math.Mat2D
	m.Array(fdct_2d_8s_testdata)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FDCT2D(m)
	}
}

func TestIDCT2D(t *testing.T) {
	var m math.Mat2D
	m.Array(idct_2d_testdata)
	n := IFDCT2D(m)
	n.Show()
	k := n.ToInt()
	fmt.Println(k)
}

func BenchmarkIDCT2D(b *testing.B) {
	b.StopTimer()
	var m math.Mat2D
	m.Array(fdct_2d_8s_testdata)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IFDCT2D(m)
	}
}
