package capitalize

import "strings"

// Capitalize -
func Capitalize(text string) string {
	data := strings.Split(text, " ")

	for pos, value := range data {
		data[pos] = strings.Title(value)
	}

	return strings.Join(data, " ")
}
