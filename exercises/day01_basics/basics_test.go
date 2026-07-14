package day1

import (
	"sort"
	"testing"
)

func TestSet(t *testing.T) {
	s := NewSet(1, 2, 2, 3)
	if got := s.Len(); got != 3 {
		t.Fatalf("Len = %d, want 3", got)
	}
	if !s.Contains(2) {
		t.Error("expected Contains(2)")
	}
	s.Remove(2)
	if s.Contains(2) {
		t.Error("expected !Contains(2) after Remove")
	}
	s.Add(9)
	if !s.Contains(9) {
		t.Error("expected Contains(9) after Add")
	}
}

func TestSetUnion(t *testing.T) {
	a := NewSet("x", "y")
	b := NewSet("y", "z")
	u := a.Union(b)
	got := u.Items()
	sort.Strings(got) // Items() order is unspecified, so sort before comparing
	want := []string{"x", "y", "z"}
	if len(got) != len(want) {
		t.Fatalf("Union len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Union[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestMapFilterReduce(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	doubled := Map(nums, func(n int) int { return n * 2 })
	if doubled[4] != 10 {
		t.Errorf("Map last = %d, want 10", doubled[4])
	}

	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	if len(evens) != 2 {
		t.Errorf("Filter len = %d, want 2", len(evens))
	}

	sum := Reduce(nums, 0, func(acc, n int) int { return acc + n })
	if sum != 15 {
		t.Errorf("Reduce sum = %d, want 15", sum)
	}

	// type-changing Map: []int -> []string
	strings := Map(nums, func(n int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	})
	if strings[0] != "odd" || strings[1] != "even" {
		t.Errorf("Map to string = %v", strings)
	}
}
