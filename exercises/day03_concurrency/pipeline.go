package day3

// A pipeline is stages connected by channels. KEY RULE for each stage: make the out
// channel, launch a GOROUTINE that runs the loop and closes out when done (defer
// close), and RETURN the channel immediately. Doing the loop in the calling goroutine
// before returning deadlocks — there's no receiver yet.

// gen emits nums on a channel and closes it when finished.
// HINT: go func(){ defer close(out); for _, n := range nums { out <- n } }()
func gen(nums ...int) <-chan int { panic("TODO: implement gen") }

// sq squares each value from in and forwards it, closing its output when in drains.
// HINT: go func(){ defer close(out); for n := range in { out <- n*n } }()
func sq(in <-chan int) <-chan int { panic("TODO: implement sq") }

// SumOfSquares wires gen -> sq -> sink and returns the total.
// HINT: for v := range sq(gen(nums...)) { total += v }
func SumOfSquares(nums ...int) int { panic("TODO: implement SumOfSquares") }
