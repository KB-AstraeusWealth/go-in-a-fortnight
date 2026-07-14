// Package day3 covers Day 3: goroutines, channels, and the canonical patterns.
//
// YOUR JOB: implement the stubs so `go test -race ./day03_concurrency/` passes.
package day3

import "sync"

// WorkerPool applies fn to every input using n concurrent workers (fan-out), then
// collects the results (fan-in). Output order must match input order. Hint: carry
// each item's index so distinct workers write distinct slice indices (race-free).
// You'll want sync.WaitGroup.

func WorkerPool[A, B any](n int, inputs []A, fn func(A) B) []B {
	type job struct {
		idx int
		val A
	}
	jobs := make(chan job)

	results := make([]B, len(inputs))

	var wg sync.WaitGroup

	wg.Add(n)

	for w := 0; w < n; w++ { // <-- exactly n goroutines, regardless of len(inputs)\
		go func() {
			defer wg.Done()
			for j := range jobs {
				res := fn(j.val)
				results[j.idx] = res
			}
		}()
	}

	for i, input := range inputs {
		jobs <- job{i, input}
	}
	close(jobs)

	wg.Wait()
	return results
}
