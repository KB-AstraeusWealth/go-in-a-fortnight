package day1

// Implement the higher-order helpers you'll miss from Scala. Package-level generic
// FUNCTIONS may introduce their own type parameters (A, B); methods may not.

// Map applies f to each element, returning a new slice (input untouched).
// HINT: out := make([]B, len(in)); out[i] = f(v). A and B may differ (e.g. []int->[]string).
func Map[A, B any](in []A, f func(A) B) []B { panic("TODO: implement Map") }

// Filter returns the elements for which pred is true.
// HINT: out := make([]A, 0, len(in)); append when pred(v) is true.
func Filter[A any](in []A, pred func(A) bool) []A { panic("TODO: implement Filter") }

// Reduce folds in from the left, starting at init.
// HINT: acc := init; acc = f(acc, v) each step; return acc. Empty input returns init.
// This is Scala's foldLeft — thread the accumulator through, don't restart from init.
func Reduce[A, B any](in []A, init B, f func(B, A) B) B { panic("TODO: implement Reduce") }
