package problem17

import (
	"log"
	"regexp"

	"github.com/divan/num2words"
)

// LettersUsed -
type LettersUsed struct {
	S int
	E int
}

// Total -
func (l *LettersUsed) Total() int {
	numberStr := ""

	for i := l.S; i <= l.E; i++ {
		numberStr += l.IntToWords(i)
	}

	return len(numberStr)
}

// IntToWords -
func (l *LettersUsed) IntToWords(number int) string {
	words := num2words.Convert(number)
	reg, err := regexp.Compile("[^a-zA-Z]+")

	if err != nil {
		log.Fatal(err)
	}

	return reg.ReplaceAllString(words, "")
}
