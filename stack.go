package vtools

// Stack is a stack contianer. A zero value stack is ready to use.
type Stack[T any] []T

// NewStack returns a Stack with the specified capacity.
// This capacity may prevent memory allocations when working with the stack.
func NewStack[T any](stackCap int) Stack[T] {
	return make([]T, 0, stackCap)
}

// Push puts t at the top of the stack.
func (s *Stack[T]) Push(t ...T) {
	*s = append(*s, t...)
}

// Pop removes the item at the top of the stack and returns it.
// Pop panics if called on an empty stack.
func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		panic("Pop called on empty stack")
	}
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}
