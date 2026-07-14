// Package day1 covers Day 1: generics-lite, methods, and map/slice semantics.
//
// YOUR JOB: replace every panic("TODO") with a real implementation so that
// `go test ./day01_basics/` goes green. The Set struct and NewSet are provided.
package day1

// Set is a generic set built on the idiomatic map[T]struct{} pattern.
// HINT: struct{} is a zero-width value, so the map stores keys only — which is
// exactly what a set is. `comparable` is required because map keys must be comparable.
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
// HINT: `s.m[v] = struct{}{}` — writing the same key twice is naturally idempotent.
func (s *Set[T]) Add(v T) { panic("TODO: implement Set.Add") }

// Remove deletes v (no-op if absent).
// HINT: the builtin delete(s.m, v) already no-ops when the key is absent.
func (s *Set[T]) Remove(v T) { panic("TODO: implement Set.Remove") }

// Contains reports whether v is present.
// HINT: comma-ok on the map: `_, ok := s.m[v]; return ok`.
func (s *Set[T]) Contains(v T) bool { panic("TODO: implement Set.Contains") }

// Len returns the number of elements.
// HINT: len() works directly on a map.
func (s *Set[T]) Len() int { panic("TODO: implement Set.Len") }

// Items returns the elements in unspecified order.
// HINT: pre-size with make([]T, 0, len(s.m)), then range the map collecting keys.
// Map order is randomized; the test sorts before comparing, so don't fight it.
func (s *Set[T]) Items() []T { panic("TODO: implement Set.Items") }

// Union returns a new set containing elements from both s and other.
// HINT: make a new set, add all of s's elements, then all of other's. Reuse Add.
func (s *Set[T]) Union(other *Set[T]) *Set[T] { panic("TODO: implement Set.Union") }
