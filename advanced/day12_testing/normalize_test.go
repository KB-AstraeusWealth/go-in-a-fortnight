package day12

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeName(t *testing.T) {
	cases := []struct{ in, want string }{
		{"  Alice   Smith ", "alice smith"},
		{"BOB", "bob"},
		{"\tcarol\ndan", "carol dan"},
		{"", ""},
	}
	for _, c := range cases {
		assert.Equal(t, c.want, NormalizeName(c.in), "input %q", c.in)
	}
}

// FuzzNormalizeName asserts properties that must hold for ANY input, rather than
// specific outputs — the sweet spot for fuzzing. Run: go test -fuzz=FuzzNormalizeName
func FuzzNormalizeName(f *testing.F) {
	f.Add("  Alice  Smith ")
	f.Add("x")
	f.Add("")
	f.Fuzz(func(t *testing.T, s string) {
		got := NormalizeName(s)
		assert.Equal(t, got, NormalizeName(got), "must be idempotent")
		for _, r := range got {
			assert.False(t, unicode.IsUpper(r), "output must have no uppercase")
		}
	})
}
