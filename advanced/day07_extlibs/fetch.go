// Package day07 covers external libraries: modules, `go get`/`go mod tidy`, semantic
// import versioning, and two ubiquitous ones — golang.org/x/sync/errgroup and
// github.com/google/uuid.
//
// YOUR JOB: implement the stubs so `go test -race ./day07_extlibs/` passes.
// Run `go mod tidy` from the advanced/ dir first to fetch the deps.
package day07

import "context"

// FetchFunc mirrors day 4. (Provided.)
type FetchFunc func(ctx context.Context, key string) (string, error)

// FetchAll is day 4's exercise again — but with errgroup instead of a hand-rolled
// WaitGroup + Once + cancel. Notice how much bookkeeping the library removes.
// HINT:
//   g, ctx := errgroup.WithContext(ctx)   // ctx auto-cancels on the first error
//   var mu sync.Mutex; results := make(map[string]string, len(keys))
//   for _, k := range keys {
//       k := k
//       g.Go(func() error {
//           v, err := fetch(ctx, k)
//           if err != nil { return err }          // returning err cancels the group
//           mu.Lock(); results[k] = v; mu.Unlock()
//           return nil
//       })
//   }
//   if err := g.Wait(); err != nil { return nil, err }
//   return results, nil
func FetchAll(ctx context.Context, keys []string, fetch FetchFunc) (map[string]string, error) {
	panic("TODO: implement FetchAll with golang.org/x/sync/errgroup")
}
