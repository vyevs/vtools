package vtools

// Set is a set data structure that keeps track of items.
// A zero-value Set is not ready to use, create one using NewSet.
type Set[T comparable] map[T]struct{}

// NewSet returns a Set with the specified capacity.
// The capacity may prevent allocations when working with the set.
func NewSet[T comparable](cap int) Set[T] {
	return make(map[T]struct{}, cap)
}

func (s Set[T]) Add(t T) {
	s[t] = struct{}{}
}

// Contains returns whether t exists in the set.
func (s Set[T]) Contains(t T) bool {
	_, found := s[t]
	return found
}

// Delete removes t from the set.
// Delete does nothing when called with an item that does not exist in the set.
func (s Set[T]) Delete(t T) {
	delete(s, t)
}
