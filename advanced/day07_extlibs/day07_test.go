package day07

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestFetchAllSuccess(t *testing.T) {
	fetch := func(ctx context.Context, key string) (string, error) { return "v-" + key, nil }
	got, err := FetchAll(context.Background(), []string{"a", "b"}, fetch)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if len(got) != 2 || got["a"] != "v-a" {
		t.Fatalf("got %v", got)
	}
}

func TestFetchAllCancelsOnError(t *testing.T) {
	boom := errors.New("boom")
	fetch := func(ctx context.Context, key string) (string, error) {
		if key == "bad" {
			return "", boom
		}
		select {
		case <-time.After(2 * time.Second):
			return "ok", nil
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
	start := time.Now()
	_, err := FetchAll(context.Background(), []string{"bad", "slow"}, fetch)
	if !errors.Is(err, boom) {
		t.Fatalf("err = %v, want boom", err)
	}
	if time.Since(start) > time.Second {
		t.Fatalf("expected fast cancellation")
	}
}

func TestNewRequestID(t *testing.T) {
	id := NewRequestID()
	if _, err := uuid.Parse(id); err != nil {
		t.Fatalf("not a valid uuid: %q (%v)", id, err)
	}
	if id == NewRequestID() {
		t.Fatal("IDs should be unique")
	}
}
