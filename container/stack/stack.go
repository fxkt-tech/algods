package stack

type Stack[T any] interface {
	Push(T)
	Pop() (T, bool)
	IsEmpty() bool
}
