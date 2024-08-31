package datastructures

type Stack[T any] struct {
	Stack []T
}

func (q *Stack[T]) Push(item T) {
	q.Stack = append(q.Stack, item)
}

func (q *Stack[T]) Pop() T {
	if len(q.Stack) == 0 {
		return *new(T)
	}
	item := q.Stack[len(q.Stack)-1]
	q.Stack = q.Stack[:len(q.Stack)-1]
	return item
}
