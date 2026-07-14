package day1

// Challenge (Day 1 "Extend it"): implement these to make set_challenge_test.go pass.

// Intersect returns a new set of the elements present in BOTH s and other.
func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	out := NewSet[T]()

	for it := range s.m {
		if other.Contains(it) {
			out.Add(it)
		}
	}

	return out
}

// Difference returns a new set of the elements in s that are NOT in other.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	out := NewSet[T]()

	for it := range s.m {
		if !other.Contains(it) {
			out.Add(it)
		}
	}

	return out
}
