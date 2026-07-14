package day08

import (
	"context"
	"os"
	"testing"
)

// testStore is integration-only: it fails (not skips) when DATABASE_URL is unset, per
// the chosen strategy. Start Postgres and export DATABASE_URL (see ../README.md).
func testStore(t *testing.T) *Store {
	t.Helper()
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Fatal("DATABASE_URL not set — start Postgres and export it (see advanced/README.md)")
	}
	ctx := context.Background()
	s, err := NewStore(ctx, dsn)
	if err != nil {
		t.Fatalf("NewStore: %v", err)
	}
	if err := s.Migrate(ctx); err != nil {
		t.Fatalf("Migrate: %v", err)
	}
	return s
}

func TestPutGet(t *testing.T) {
	s := testStore(t)
	defer s.Close()
	ctx := context.Background()

	if err := s.Put(ctx, Item{ID: "1", Name: "widget"}); err != nil {
		t.Fatalf("Put: %v", err)
	}
	got, ok, err := s.Get(ctx, "1")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if !ok {
		t.Fatal("expected item to exist")
	}
	if got.Name != "widget" {
		t.Errorf("Name = %q, want widget", got.Name)
	}
}

func TestUpsertAndMissing(t *testing.T) {
	s := testStore(t)
	defer s.Close()
	ctx := context.Background()

	_ = s.Put(ctx, Item{ID: "2", Name: "old"})
	if err := s.Put(ctx, Item{ID: "2", Name: "new"}); err != nil {
		t.Fatalf("upsert: %v", err)
	}
	got, _, _ := s.Get(ctx, "2")
	if got.Name != "new" {
		t.Errorf("after upsert Name = %q, want new", got.Name)
	}

	_, ok, err := s.Get(ctx, "does-not-exist")
	if err != nil {
		t.Fatalf("Get missing returned error: %v", err)
	}
	if ok {
		t.Error("expected not found for missing id")
	}
}
