package queue

type ArrayQueue struct {
	arr []interface{} // abstract stack.
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{}
}

func (q *ArrayQueue) Push(val interface{}) {
	q.arr = append(q.arr, val)
}

func (q *ArrayQueue) Pop() (interface{}, bool) {
	if len(q.arr) == 0 {
		return 0, false
	}
	elem := q.arr[0]
	q.arr = q.arr[1:]
	return elem, true
}

func (q *ArrayQueue) Size() int {
	return len(q.arr)
}

func (q *ArrayQueue) IsEmpty() bool {
	return len(q.arr) == 0
}
