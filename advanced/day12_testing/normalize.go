// Package day12 covers testing technique: table tests + testify, fuzzing, and
// testcontainers. NormalizeName is offline (unit + fuzz, red→green). The Postgres test
// uses testcontainers and needs Docker. Run `go mod tidy` first.
package day12

// NormalizeName trims surrounding whitespace, lowercases, and collapses internal runs
// of whitespace to single spaces:  "  Alice   Smith " -> "alice smith".
// HINT: strings.Fields splits on any whitespace and drops empty tokens; join with " ";
// strings.ToLower the result. (Add the "strings" import.)
func NormalizeName(s string) string { panic("TODO: implement NormalizeName") }
