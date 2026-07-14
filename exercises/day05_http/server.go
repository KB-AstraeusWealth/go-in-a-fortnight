// Package day5 covers Day 5: net/http, the 1.22 mux, middleware, and JSON.
//
// YOUR JOB: implement the stubs so `go test ./day05_http/` passes. You'll add the
// encoding/json and log/slog imports yourself.
package day5

import (
	"encoding/json"
	"log/slog"
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

// Get returns the item and whether it existed. Guard with the mutex.
func (s *Store) Get(id string) (Item, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item, ok := s.items[id]

	return item, ok
}

// Put stores the item. Guard with the mutex.
func (s *Store) Put(it Item) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items[it.ID] = it
}

// NewRouter wires:
//
//	GET  /items/{id}  -> 200 with the item as JSON, or 404
//	POST /items       -> decode JSON body, 400 if bad or missing id, else 201
//
// using the Go 1.22 mux (r.PathValue) and logging middleware.
func NewRouter(store *Store) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /items", func(w http.ResponseWriter, r *http.Request) {
		var item Item
		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			store.Put(item)
			w.WriteHeader(http.StatusCreated)
		}
	})

	mux.HandleFunc("GET /items/{id}", func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		item, ok := store.Get(idString)
		if !ok || item.ID == "" {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(item)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	})

	logging := func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			slog.Info("request", "method", r.Method, "path", r.URL.Path)
			handler.ServeHTTP(w, r)
		})
	}

	return logging(mux)
}
