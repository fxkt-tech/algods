package sort

import (
	"golang.org/x/exp/constraints"
)

type AnySlice[T constraints.Ordered] []T

func (x AnySlice[T]) Len() int           { return 0 }
func (x AnySlice[T]) Less(i, j int) bool { return x[i] < x[j] }
func (x AnySlice[T]) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
