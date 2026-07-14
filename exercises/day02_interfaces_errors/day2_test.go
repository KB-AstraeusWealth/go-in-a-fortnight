package day2

import (
	"errors"
	"testing"
)

func TestLookupOK(t *testing.T) {
	s := NewStore()
	id, err := s.Lookup("alice")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != 1 {
		t.Errorf("id = %d, want 1", id)
	}
}

func TestLookupNotFound(t *testing.T) {
	s := NewStore()
	_, err := s.Lookup("carol")
	if err == nil {
		t.Fatal("expected an error")
	}
	// errors.Is unwraps the %w chain to find the sentinel.
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("errors.Is(err, ErrNotFound) = false; err = %v", err)
	}
}

func TestLookupValidation(t *testing.T) {
	s := NewStore()
	_, err := s.Lookup("")
	var ve *ValidationError
	// errors.As finds the first error in the chain matching the target type.
	if !errors.As(err, &ve) {
		t.Fatalf("expected *ValidationError, got %T", err)
	}
	if ve.Field != "name" {
		t.Errorf("Field = %q, want name", ve.Field)
	}
}

func TestDescribe(t *testing.T) {
	cases := []struct {
		name  string
		shape Shape
		want  string
	}{
		{"circle", Circle{R: 2}, "circle r=2"},
		{"rect", Rectangle{W: 3, H: 4}, "rectangle 3x4"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Describe(tc.shape); got != tc.want {
				t.Errorf("Describe = %q, want %q", got, tc.want)
			}
		})
	}
}
