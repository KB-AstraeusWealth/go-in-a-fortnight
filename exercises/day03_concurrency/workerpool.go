// Package day3 covers Day 3: goroutines, channels, and the canonical patterns.
//
// YOUR JOB: implement the stubs so `go test -race ./day03_concurrency/` passes.
package day3

// WorkerPool applies fn to every input using n concurrent workers (fan-out), then
// collects the results (fan-in). Output order must match input order.
// HINT:
//   - jobs := make(chan job), where `job` carries {idx int; val A} so the index
//     travels WITH the work.
//   - launch EXACTLY n goroutines, each: `for j := range jobs { results[j.idx] = fn(j.val) }`.
//     distinct indices => race-free writes, no mutex needed (prove it with -race).
//   - after launching, send every input as a job, then close(jobs) so the ranges end.
//   - wg.Add(n) before launching; wg.Wait() before returning. (You'll want sync.WaitGroup.)
func WorkerPool[A, B any](n int, inputs []A, fn func(A) B) []B {
	panic("TODO: implement WorkerPool")
}
