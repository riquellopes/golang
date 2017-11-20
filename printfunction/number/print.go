package number

import (
	"strconv"
	"strings"
)

func prints(number int) string {
	list := make([]string, number)

	for i := 0; i < number; i++ {
		list[i] = strconv.Itoa(i + 1)
	}
	return strings.Join(list, "")
}
