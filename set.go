package vtools

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](cap int) Set[T] {
	return make(map[T]struct{}, cap)
}

func (s Set[T]) Add(t T) {
	s[t] = struct{}{}
}

func (s Set[T]) Contains(t T) bool {
	_, found := s[t]
	return found
}

func (s Set[T]) Delete(t T) {
	delete(s, t)
}
