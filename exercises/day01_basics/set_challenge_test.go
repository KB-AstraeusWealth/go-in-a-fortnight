package day1

import (
	"sort"
	"testing"
)

func TestIntersect(t *testing.T) {
	got := NewSet(1, 2, 3).Intersect(NewSet(2, 3, 4)).Items()
	sort.Ints(got)
	want := []int{2, 3}
	if len(got) != len(want) {
		t.Fatalf("Intersect len = %d, want %d (%v)", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Intersect[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestDifference(t *testing.T) {
	// elements in a that are not in b -> {1}
	got := NewSet(1, 2, 3).Difference(NewSet(2, 3, 4)).Items()
	sort.Ints(got)
	if len(got) != 1 || got[0] != 1 {
		t.Fatalf("Difference = %v, want [1]", got)
	}
}
