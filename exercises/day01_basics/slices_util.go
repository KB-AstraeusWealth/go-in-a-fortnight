package day1

// Implement the higher-order helpers you'll miss from Scala. Package-level generic
// FUNCTIONS may introduce their own type parameters (A, B); methods may not.

// Map applies f to each element, returning a new slice (input untouched).
func Map[A, B any](in []A, f func(A) B) []B {
	out := make([]B, len(in))

	for i := range in {
		out[i] = f(in[i])
	}

	return out
}

// Filter returns the elements for which pred is true.
func Filter[A any](in []A, pred func(A) bool) []A {
	out := make([]A, 0, len(in))

	for i := range in {
		if pred(in[i]) {
			out = append(out, in[i])
		}
	}

	return out
}

// FilterNot returns the elements for which pred is false.
func FilterNot[A any](in []A, pred func(A) bool) []A {
	out := make([]A, 0, len(in))

	for i := range in {
		if !pred(in[i]) {
			out = append(out, in[i])
		}
	}

	return out
}

// Reduce folds in from the left, starting at init.
func Reduce[A, B any](in []A, init B, f func(B, A) B) B {
	acc := init
	for i := range in {
		acc = f(acc, in[i])
	}

	return acc
}
