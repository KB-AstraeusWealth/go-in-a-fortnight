package day4

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestFetchAllSuccess(t *testing.T) {
	fetch := func(ctx context.Context, key string) (string, error) {
		return "val-" + key, nil
	}
	got, err := FetchAll(context.Background(), []string{"a", "b", "c"}, fetch)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 3 {
		t.Fatalf("len = %d, want 3", len(got))
	}
	if got["a"] != "val-a" {
		t.Errorf("got[a] = %q, want val-a", got["a"])
	}
}

func TestFetchAllCancelsOnFirstError(t *testing.T) {
	boom := errors.New("boom")
	fetch := func(ctx context.Context, key string) (string, error) {
		if key == "bad" {
			return "", boom // fails immediately
		}
		// The slow keys should be canceled before this timer fires.
		select {
		case <-time.After(2 * time.Second):
			return "ok", nil
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}

	start := time.Now()
	_, err := FetchAll(context.Background(), []string{"bad", "slow1", "slow2"}, fetch)
	elapsed := time.Since(start)

	if !errors.Is(err, boom) {
		t.Fatalf("err = %v, want boom", err)
	}
	if elapsed > time.Second {
		t.Errorf("took %v; expected fast cancellation, not the 2s timeout", elapsed)
	}
}
