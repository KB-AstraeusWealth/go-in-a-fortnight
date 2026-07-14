// Command api is the runnable Day 5 service: it demonstrates production-shaped
// startup and, crucially, GRACEFUL SHUTDOWN driven by an OS signal.
//
// Run:  go run ./cmd/api    (Ctrl-C to shut down cleanly)
package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	day5 "github.com/kb/go-in-a-week/day05_http"
)

func main() {
	// NotifyContext gives us a context that is canceled on SIGINT/SIGTERM.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           day5.NewRouter(day5.NewStore()),
		ReadHeaderTimeout: 5 * time.Second, // basic hardening
	}

	// Run the server in a goroutine so main can wait for the shutdown signal.
	go func() {
		slog.Info("listening", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server error", "err", err)
		}
	}()

	<-ctx.Done() // block until a signal arrives
	slog.Info("shutdown signal received; draining connections")

	// Give in-flight requests up to 10s to finish, then force close.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("graceful shutdown failed", "err", err)
		return
	}
	slog.Info("shutdown complete")
}
