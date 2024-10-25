package vtools

// Queue is a generic queue data structure. Its zero-value is ready to use.
type Queue[T any] []T

// NewQueue returns a queue with an initialized capacity. Use of cap prevents later memory allocations.
func NewQueue[T any](cap int) Queue[T] {
	return make([]T, 0, cap)
}

// Push pushes items to the back of the queue.
func (q *Queue[T]) Push(items ...T) {
	*q = append(*q, items...)
}

// Pop removes and returns the item at the front of the queue.
// Pop panics if called on an empty queue.
func (q *Queue[T]) Pop() T {
	if len(*q) == 0 {
		panic("Pop called on empty queue")
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}
