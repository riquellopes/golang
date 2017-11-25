package problemOne

// Multiple -
func Multiple(below int) int {
	theSum := 0

	for i := 1; i < below; i++ {
		if i%3 == 0 || i%5 == 0 {
			theSum += i
		}
	}

	return theSum
}
