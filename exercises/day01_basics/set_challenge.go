package day1

// Challenge (Day 1 "Extend it"): implement these to make set_challenge_test.go pass.

// Intersect returns a new set of the elements present in BOTH s and other.
// HINT: new set; for each element of s, keep it if other.Contains(it).
func (s *Set[T]) Intersect(other *Set[T]) *Set[T] { panic("TODO: implement Set.Intersect") }

// Difference returns a new set of the elements in s that are NOT in other.
// HINT: new set; for each element of s, keep it if !other.Contains(it).
func (s *Set[T]) Difference(other *Set[T]) *Set[T] { panic("TODO: implement Set.Difference") }
