package datastructures

// Stack is a generic stack implementation
type Stack[T any] struct {
	Stack []T
}

// Push adds an item to the stack
func (q *Stack[T]) Push(item T) {
	q.Stack = append(q.Stack, item)
}

// Pop removes and returns the last item from the stack
func (q *Stack[T]) Pop() T {
	if len(q.Stack) == 0 {
		return *new(T)
	}
	item := q.Stack[len(q.Stack)-1]
	q.Stack = q.Stack[:len(q.Stack)-1]
	return item
}
