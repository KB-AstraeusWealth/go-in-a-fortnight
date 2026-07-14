package day3

// A pipeline is stages connected by channels. KEY RULE for each stage: make the out
// channel, launch a GOROUTINE that runs the loop and closes out when done (defer
// close), and RETURN the channel immediately. Doing the loop in the calling goroutine
// before returning deadlocks — there's no receiver yet.

// gen emits nums on a channel and closes it when finished.
func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}

		close(out)
	}()

	return out
}

// sq squares each value from in and forwards it, closing its output when in drains.
func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}

		close(out)
	}()

	return out
}

// SumOfSquares wires gen -> sq -> sink and returns the total.
func SumOfSquares(nums ...int) int {

	data := gen(nums...)
	acc := 0

	for n := range sq(data) {
		acc += n
	}

	return acc
}
