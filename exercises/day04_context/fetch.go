// Package day4 covers Day 4: context, cancellation, and correctness.
//
// YOUR JOB: implement FetchAll so `go test -race ./day04_context/` passes.
package day4

import (
	"context"
	"sync"
)

// FetchFunc simulates fetching a resource; a real one MUST honor ctx. (Provided.)
type FetchFunc func(ctx context.Context, key string) (string, error)

// FetchAll fetches all keys concurrently and returns results keyed by input. As soon
// as any fetch fails, cancel the others and return that first error.
// HINT:
//   - ctx, cancel := context.WithCancel(ctx); defer cancel().
//   - one goroutine per key; wg.Add(len(keys)).
//   - pass the DERIVED ctx into fetch, or cancel won't stop the others.
//   - on error: once.Do(func(){ firstErr = err; cancel() }); return.
//   - on success: lock ONLY the map write (mu around results[k]=v). NEVER hold the
//     lock across fetch — that serializes everything and breaks the timing test.
//   - wg.Wait(); return firstErr!=nil ? (nil,err) : (results,nil).
func FetchAll(ctx context.Context, keys []string, fetch FetchFunc) (map[string]string, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var once sync.Once
	var firstError error

	newCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	results := make(map[string]string, len(keys))

	wg.Add(len(keys))

	for _, k := range keys {
		go func() {
			defer wg.Done()
			res, err := fetch(newCtx, k)
			if err != nil {
				once.Do(func() {
					firstError = err
					cancel()
				})

				return
			}
			mu.Lock()
			results[k] = res
			mu.Unlock()
		}()
	}

	wg.Wait()
	return results, firstError
}
