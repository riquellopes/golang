package mostcommon

import "fmt"

// MostCommon -
func MostCommon(str string) map[string]int {
	m := map[string]int{}
	// m := make(map[string]int)

	for key, _ := range str {
		if _, ok := m[""]; ok {
			fmt.Printf("%s\n", key)
		}
	}

	m["b"] = 3
	m["a"] = 2
	m["c"] = 2
	return m
}
