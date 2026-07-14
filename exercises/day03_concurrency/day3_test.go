package day3

import "testing"

func TestWorkerPoolPreservesOrder(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8}
	got := WorkerPool(3, in, func(n int) int { return n * n })
	want := []int{1, 4, 9, 16, 25, 36, 49, 64}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestWorkerPoolEmpty(t *testing.T) {
	got := WorkerPool(4, []int{}, func(n int) int { return n })
	if len(got) != 0 {
		t.Errorf("len = %d, want 0", len(got))
	}
}

func TestSumOfSquares(t *testing.T) {
	if got := SumOfSquares(1, 2, 3); got != 14 { // 1 + 4 + 9
		t.Errorf("SumOfSquares = %d, want 14", got)
	}
}
