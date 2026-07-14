// Package day13 covers the functional-options pattern and lifecycle config. Stdlib
// only — no `go get` needed. YOUR JOB: implement so `go test ./day13_config/` passes.
package day13

import "time"

// Server holds configuration assembled from functional options. (Provided.)
type Server struct {
	Addr        string
	ReadTimeout time.Duration
	MaxConns    int
}

// Option mutates a Server during construction. (Provided.)
// This is the idiomatic Go alternative to a huge constructor or a config struct with
// dozens of optional fields: each option is a small closure that sets one thing.
type Option func(*Server)

// New builds a Server with sensible defaults, then applies each option in order.
// HINT:
//   s := &Server{Addr: ":8080", ReadTimeout: 5 * time.Second, MaxConns: 100}
//   for _, opt := range opts { opt(s) }
//   return s
func New(opts ...Option) *Server { panic("TODO: implement New") }

// WithAddr overrides the listen address.
// HINT: return func(s *Server) { s.Addr = addr }
func WithAddr(addr string) Option { panic("TODO: implement WithAddr") }

// WithReadTimeout overrides the read timeout.
func WithReadTimeout(d time.Duration) Option { panic("TODO: implement WithReadTimeout") }

// WithMaxConns overrides the max connection count.
func WithMaxConns(n int) Option { panic("TODO: implement WithMaxConns") }
