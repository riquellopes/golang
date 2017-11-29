package mostcommon

// MostCommon -
func MostCommon(str string) map[string]int {
	// m := map[string]int{}
	m := make(map[string]int)

	for _, char := range str {
		if _, ok := m[string(char)]; ok {
			m[string(char)]++
		} else {
			m[string(char)] = 1
		}
	}

	return m
}
