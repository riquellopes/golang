package main

import "log"

// CountSubstring -
func CountSubstring(s, sub string) int {
	occurrences := 0

	for i := range s {
		length := len(sub) + i

		cut := s[i:length]

		if cut == sub {
			occurrences++
		}

		if length == len(s) {
			break
		}

	}

	return occurrences
}

func main() {
	log.Println(CountSubstring("ABCDCDC", "CDC"))

}
