package vtools

import (
	"bufio"
	"fmt"
	"iter"
	"os"
	"time"
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

// Cycle returns an iterator that endlessly loops over s.
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

// Filter returns a slice containing only elements of s for which shouldKeep returns true.
func Filter[T any](s []T, shouldKeep func(T) bool) []T {
	out := make([]T, 0, len(s))
	for _, item := range s {
		if shouldKeep(item) {
			out = append(out, item)
		}
	}
	return out
}

// Map returns a slice of the items in s with to(item) called on each one.
func Map[T any, E any](s []T, to func(T) E) []E {
	out := make([]E, 0, len(s))
	for _, item := range s {
		out = append(out, to(item))
	}
	return out
}

// ReadLines reads fPath and returns all the non-empty lines.
func ReadLines(fPath string) ([]string, error) {
	f, err := os.Open(fPath)
	if err != nil {
		return nil, fmt.Errorf("os.Open failed: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	lines := make([]string, 0, 32)
	for sc.Scan() {
		line := sc.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("error reading lines: %v", err)
	}

	return lines, nil
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
