package queue

type ArrayQueue[T any] struct {
	arr []T // abstract stack.
}

func NewArrayQueue[T any]() *ArrayQueue[T] {
	return &ArrayQueue[T]{}
}

func (q *ArrayQueue[T]) Push(val T) {
	q.arr = append(q.arr, val)
}

func (q *ArrayQueue[T]) Pop() (t T, b bool) {
	if len(q.arr) == 0 {
		return
	}
	elem := q.arr[0]
	q.arr = q.arr[1:]
	return elem, true
}

func (q *ArrayQueue[T]) Size() int {
	return len(q.arr)
}

func (q *ArrayQueue[T]) IsEmpty() bool {
	return len(q.arr) == 0
}
