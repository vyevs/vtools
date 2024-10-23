package vtools

import (
	"fmt"
	"iter"
	"time"

	"golang.org/x/exp/constraints"
)

// Any returns true if f(item) is true for at least one item in seq, otherwise false.
func Any[T any](seq iter.Seq[T], f func(T) bool) bool {
	for v := range seq {
		if f(v) {
			return true
		}
	}
	return false
}

// AnySlice returns true if f(item) is true for at least one item in s, otherwise false.
func AnySlice[T any](s []T, f func(T) bool) bool {
	for _, item := range s {
		if f(item) {
			return true
		}
	}
	return false
}

// AllSlice returns true if f(item) returns true for all items in s, otherwise false.
func AllSlice[T any](s []T, f func(T) bool) bool {
	for _, it := range s {
		if !f(it) {
			return false
		}
	}
	return true
}

// StrBytes returns an iterator over the bytes in s.
func StrBytes(s string) iter.Seq[byte] {
	return func(yield func(b byte) bool) {
		for i := range len(s) {
			if !yield(s[i]) {
				return
			}
		}
	}
}

// Count returns the number of items in s that are equal to target.
func Count[T comparable](s iter.Seq[T], target T) int {
	var ct int
	for v := range s {
		if v == target {
			ct++
		}
	}
	return ct
}

// CountSlice returns the count of elements in s equal to target.
func CountSlice[T comparable](s []T, target T) int {
	var ct int
	for _, item := range s {
		if item == target {
			ct++
		}
	}
	return ct
}

// CountFunc returns the count of items in s for which shouldCount(item) returns true.
func CountFunc[T any](s []T, shouldCount func(T) bool) int {
	var ct int
	for _, item := range s {
		if shouldCount(item) {
			ct++
		}
	}
	return ct
}

// Counter returns a map whose keys are items in s and whose values are the counts of each item.
// e.g. Counter([]string{7, 1, 7, 9, 1, 3}) == map[string]int{1: 2, 7: 2, 3: 1, 9: 1}
func Counter[T comparable](s iter.Seq[T]) map[T]int {
	out := make(map[T]int, 8)
	for item := range s {
		out[item]++
	}
	return out
}

// CounterSlice returns a map whose keys are items in s and whose values are the counts of each item.
// e.g. Counter([]string{7, 1, 7, 9, 1, 3}) == map[string]int{1: 2, 7: 2, 3: 1, 9: 1}
func CounterSlice[T comparable](s []T) map[T]int {
	out := make(map[T]int, len(s))
	for _, item := range s {
		out[item]++
	}
	return out
}

// Cycle returns an iterator that endlessly loops over s. Analogous to python's itertools.cycle.
func Cycle[T any](s []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		var i int
		for {
			if !yield(s[i]) {
				return
			}
			i = (i + 1) % len(s)
		}
	}
}

// Enumerate returns an iterator where the 1st item in the pair is an index and the 2nd is the item in s. Analogous to python's enumerate.
func Enumerate[T any](s iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		var i int
		for v := range s {
			if !yield(i, v) {
				return
			}
			i++
		}
	}
}

// Filter returns a sequence of items from s for which shouldKeep(item) returns true.
func Filter[T any](s iter.Seq[T], shouldKeep func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s {
			if shouldKeep(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// FilterSlice returns a slice containing only elements of s for which shouldKeep returns true.
func FilterSlice[T any](s []T, shouldKeep func(T) bool) []T {
	out := make([]T, 0, len(s))
	for _, item := range s {
		if shouldKeep(item) {
			out = append(out, item)
		}
	}
	return out
}

// Map maps the items in s to a new sequence by calling to(s)
func Map[T any, E any](s iter.Seq[T], to func(T) E) iter.Seq[E] {
	return func(yield func(e E) bool) {
		for t := range s {
			if !yield(to(t)) {
				return
			}
		}
	}
}

// MapSlice returns a slice of the items in s with to(item) called on each one.
func MapSlice[T any, E any](s []T, to func(T) E) []E {
	out := make([]E, 0, len(s))
	for _, item := range s {
		out = append(out, to(item))
	}
	return out
}

// MaxIndex returns the max value in s along with it's index.
// If there are multiple max value occurrences, the index of the first one is returned.
func MaxIndex[T constraints.Ordered](s []T) (T, int) {
	if len(s) == 0 {
		panic("slice with length 0") // Same behavior as slices.Max
	}

	max, maxI := s[0], 0
	for i := 1; i < len(s); i++ {
		v := s[i]

		if v > max {
			max = v
			maxI = i
		}
	}

	return max, maxI
}

// Number is a constraint that contains all number types.
type Number interface {
	constraints.Integer | constraints.Float
}

// Sum returns the sum of a sequence of a numbers.
func Sum[T Number](s iter.Seq[T]) T {
	var sum T
	for v := range s {
		sum += v
	}
	return sum
}

// SumSlice returns the sum of a slice of numbers.
func SumSlice[T Number](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

// SplitWS returns s split on all types of whitespace.
func SplitWS(s string) []string {
	parts := make([]string, 0, 8)
	buf := make([]byte, 0, 16)

	for char := range StrBytes(s) {
		if char == ' ' || char == '\t' || char == '\n' {
			if len(buf) != 0 {
				parts = append(parts, string(buf))
				buf = buf[:0]
			}
		} else {

			buf = append(buf, char)
		}
	}

	return parts
}

// TimeIt prints to stdout the time some action took.
// Intended usage is defer TimeIt(time.Now(), "foo")
func TimeIt(start time.Time, action string) {
	fmt.Printf("%s took %s\n", action, time.Since(start))
}
