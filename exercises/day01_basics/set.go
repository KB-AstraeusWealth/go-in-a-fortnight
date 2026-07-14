// Package day1 covers Day 1: generics-lite, methods, and map/slice semantics.
//
// YOUR JOB: replace every panic("TODO") with a real implementation so that
// `go test ./day01_basics/` goes green. The Set struct and NewSet are provided.
package day1

// Set is a generic set built on the idiomatic map[T]struct{} pattern.
type Set[T comparable] struct {
	m map[T]struct{}
}

// NewSet returns a ready-to-use set. (Provided.)
func NewSet[T comparable](items ...T) *Set[T] {
	s := &Set[T]{m: make(map[T]struct{}, len(items))}
	for _, it := range items {
		s.m[it] = struct{}{}
	}
	return s
}

// Add inserts v (idempotent).
func (s *Set[T]) Add(v T) {
	s.m[v] = struct{}{}
}

// Remove deletes v (no-op if absent).
func (s *Set[T]) Remove(v T) {
	delete(s.m, v)
}

// Contains reports whether v is present.
func (s *Set[T]) Contains(v T) bool {
	_, ok := s.m[v]

	return ok
}

// Len returns the number of elements.
func (s *Set[T]) Len() int {
	return len(s.m)
}

// Items returns the elements in unspecified order.
func (s *Set[T]) Items() []T {
	items := make([]T, 0, len(s.m))

	for it := range s.m {
		items = append(items, it)
	}

	return items
}

// Union returns a new set containing elements from both s and other.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	newSet := NewSet[T]()

	for it := range s.m {
		newSet.Add(it)
	}

	for it := range other.m {
		newSet.Add(it)
	}

	return newSet
}
