package datastructures

type Queue[T any] struct {
	Queue []T
}

func (q *Queue[T]) Push(item T) {
	q.Queue = append(q.Queue, item)
}

func (q *Queue[T]) Pop() T {
	if len(q.Queue) == 0 {
		return *new(T)
	}
	item := q.Queue[0]
	q.Queue = q.Queue[1:]
	return item
}
