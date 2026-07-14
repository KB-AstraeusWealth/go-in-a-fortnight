// Package day08 covers PostgreSQL with pgx v5 (pgxpool). INTEGRATION-ONLY: needs a
// running Postgres and DATABASE_URL set (see ../README.md). Run `go mod tidy` first.
//
// YOUR JOB: implement the stubs so `go test ./day08_postgres/` passes with a DB up.
package day08

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Item is the stored resource.
type Item struct {
	ID   string
	Name string
}

// Store wraps a pgx connection pool. (Struct provided; methods are yours.)
type Store struct {
	pool *pgxpool.Pool
}

// NewStore opens a connection pool against connString.
// HINT: pool, err := pgxpool.New(ctx, connString); if err != nil { return nil, err };
//       return &Store{pool: pool}, nil. (pgxpool.New is lazy; the first query connects.)
func NewStore(ctx context.Context, connString string) (*Store, error) {
	panic("TODO: implement NewStore")
}

// Close releases the pool. HINT: s.pool.Close()
func (s *Store) Close() { panic("TODO: implement Close") }

// Migrate creates the items table if it doesn't exist.
// HINT: _, err := s.pool.Exec(ctx,
//   `CREATE TABLE IF NOT EXISTS items (id TEXT PRIMARY KEY, name TEXT NOT NULL)`); return err
func (s *Store) Migrate(ctx context.Context) error { panic("TODO: implement Migrate") }

// Put upserts an item. Note the $1/$2 placeholders — never string-concat SQL.
// HINT: _, err := s.pool.Exec(ctx,
//   `INSERT INTO items (id, name) VALUES ($1, $2)
//    ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name`, it.ID, it.Name); return err
func (s *Store) Put(ctx context.Context, it Item) error { panic("TODO: implement Put") }

// Get fetches an item by id; returns (Item{}, false, nil) when it doesn't exist.
// HINT:
//   var it Item
//   err := s.pool.QueryRow(ctx, `SELECT id, name FROM items WHERE id = $1`, id).Scan(&it.ID, &it.Name)
//   if errors.Is(err, pgx.ErrNoRows) { return Item{}, false, nil }   // add pgx + errors imports
//   if err != nil { return Item{}, false, err }
//   return it, true, nil
func (s *Store) Get(ctx context.Context, id string) (Item, bool, error) {
	panic("TODO: implement Get")
}
