// Package day5 covers Day 5: net/http, the 1.22 mux, middleware, and JSON.
//
// YOUR JOB: implement the stubs so `go test ./day05_http/` passes. You'll add the
// encoding/json and log/slog imports yourself.
package day5

import (
	"net/http"
	"sync"
)

// Item is the resource. (Provided.)
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Store is a thread-safe in-memory store. (Struct + NewStore provided.)
type Store struct {
	mu    sync.RWMutex
	items map[string]Item
}

func NewStore() *Store {
	return &Store{items: make(map[string]Item)}
}

// Get returns the item and whether it existed.
// HINT: RLock for reads (many readers OK); defer RUnlock; comma-ok on the map.
func (s *Store) Get(id string) (Item, bool) { panic("TODO: implement Store.Get") }

// Put stores the item.
// HINT: Lock (exclusive) for writes; defer Unlock.
func (s *Store) Put(it Item) { panic("TODO: implement Store.Put") }

// NewRouter wires the routes and logging middleware.
// HINT:
//   - Register routes ONCE here, not inside a handler. Use the 1.22 mux with method+path:
//     mux.HandleFunc("GET /items/{id}", ...) and mux.HandleFunc("POST /items", ...).
//     r.PathValue("id") reads the {id} wildcard.
//   - GET (guard-clause): store.Get; if !ok -> http.Error(w, "not found", 404); else set
//     Content-Type: application/json, then json.NewEncoder(w).Encode(item).
//   - POST: json.NewDecoder(r.Body).Decode(&item); if err or item.ID=="" -> 400; else
//     store.Put(item); w.WriteHeader(http.StatusCreated).
//   - Middleware is func(http.Handler) http.Handler: log method/path with slog, THEN call
//     next.ServeHTTP(w, r) (forgetting that call is the classic bug). Return logging(mux).
func NewRouter(store *Store) http.Handler { panic("TODO: implement NewRouter") }
