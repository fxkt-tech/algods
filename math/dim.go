package math

import "constraints"

// 整型绝对值
func Abs[T constraints.Integer | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Clip[T constraints.Ordered](x, a, b T) T {
	return Min(Max(x, a), b)
}
