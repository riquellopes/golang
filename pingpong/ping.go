package pingpong

import (
	"fmt"
	"strconv"
)

func ping(number int) string {
	var ping, pong, num string = "", "", ""

	if divisible(number, 3) {
		ping = "Ping"
	}

	if divisible(number, 5) {
		pong = "Pong"
	}

	if len(ping) == 0 && len(pong) == 0 {
		num = strconv.Itoa(number)
	}

	return fmt.Sprintf("%s%s%s", ping, pong, num)
}

func divisible(number, divider int) bool {
	if (number % divider) == 0 {
		return true
	}
	return false
}
