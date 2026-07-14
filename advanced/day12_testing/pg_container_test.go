package day12

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

// TestPostgresRoundTrip spins a real Postgres in a container, so it needs Docker
// running. Reference pattern for integration tests that own their infrastructure:
// no shared DB, no manual setup, cleaned up automatically.
func TestPostgresRoundTrip(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping container test in -short mode")
	}
	ctx := context.Background()

	container, err := postgres.Run(ctx, "postgres:16",
		postgres.WithDatabase("test"),
		postgres.WithUsername("test"),
		postgres.WithPassword("test"),
	)
	require.NoError(t, err)
	defer func() { _ = container.Terminate(ctx) }()

	dsn, err := container.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)

	conn, err := pgx.Connect(ctx, dsn)
	require.NoError(t, err)
	defer conn.Close(ctx)

	_, err = conn.Exec(ctx, `CREATE TABLE names (v TEXT NOT NULL)`)
	require.NoError(t, err)
	_, err = conn.Exec(ctx, `INSERT INTO names (v) VALUES ($1), ($2)`, "a", "b")
	require.NoError(t, err)

	var count int
	require.NoError(t, conn.QueryRow(ctx, `SELECT count(*) FROM names`).Scan(&count))
	require.Equal(t, 2, count)
}
