package go_koans

func divideFourBy(i int) int {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok && err != nil {
				println("Are you dividing by zero, you little shit???")
			}
		}
	}()
	return 4 / i
}

const __divisor__ = 0

func aboutPanics() {
	// panics are exceptional errors at runtime

	n := divideFourBy(2)
	assert(n == 2) // panics are exceptional errors at runtime
}
