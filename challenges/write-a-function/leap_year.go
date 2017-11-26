package leapyear

// isLeap -
func isLeap(year int) bool {
	divisibleByFour := year%4 == 0

	divisibleByHundred := year%100 != 0

	if divisibleByFour && divisibleByHundred || year%400 == 0 {
		return true
	}
	return false
}
