package vtools

// Set is a set data structure that keeps track of items.
// A zero-value Set is not ready to use, create one using NewSet.
type Set[T comparable] map[T]struct{}

// NewSet returns a Set with the specified capacity.
// The capacity may prevent allocations when working with the set.
func NewSet[T comparable](cap int) Set[T] {
	return make(map[T]struct{}, cap)
}

// Add adds 1 or more items to the set.
func (s Set[T]) Add(items ...T) {
	for _, it := range items {
		s[it] = struct{}{}
	}
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

// SetFromSlice returns a Set with all items in s present in the set. A utility to simplify set creation + initialization.
func SetFromSlice[T comparable](items []T) Set[T] {
	s := NewSet[T](len(items))
	s.Add(items...)
	return s
}
