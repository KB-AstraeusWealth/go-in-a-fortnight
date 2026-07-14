// Package day2 covers Day 2: interfaces, type switches, and idiomatic errors.
//
// YOUR JOB: implement the stubs so `go test ./day02_interfaces_errors/` passes.
package day2

import (
	"errors"
	"fmt"
)

// ErrNotFound is a sentinel error. (Provided.)
var ErrNotFound = errors.New("not found")

// ValidationError is a custom typed error. (Struct provided; method is yours.)
type ValidationError struct {
	Field string
	Msg   string
}

// Error must satisfy the error interface (one method: Error() string).
// HINT: return a formatted string via fmt.Sprintf. (Add the "fmt" import.)
func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s %s", e.Field, e.Msg)
}

// Store is a tiny lookup table. (Provided.)
type Store struct {
	data map[string]int
}

// NewStore returns a seeded Store. (Provided.)
func NewStore() *Store {
	return &Store{data: map[string]int{"alice": 1, "bob": 2}}
}

// Lookup returns the id for name.
// HINT (guard-clause style):
//   - empty name  -> return 0, &ValidationError{Field: "name", Msg: ...}
//   - missing key -> wrap the sentinel so errors.Is works: fmt.Errorf("lookup %q: %w", name, ErrNotFound)
//   - otherwise   -> return id, nil
//
// The test uses errors.Is(err, ErrNotFound), errors.As for *ValidationError, and expects Field=="name".
func (s *Store) Lookup(name string) (int, error) {
	if name == "" {
		return 0, &ValidationError{Field: "name", Msg: "name is required"}
	}
	id, ok := s.data[name]
	if !ok {
		return 0, fmt.Errorf("lookup %q: %w", name, ErrNotFound)
	}
	return id, nil
}
