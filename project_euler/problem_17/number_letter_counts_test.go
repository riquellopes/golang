package problem17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneToFive(t *testing.T) {
	lettersUsed := LettersUsed{1, 5}

	assert.Equal(t, lettersUsed.Total(), 19)
}

func TestOneToTwo(t *testing.T) {
	lettersUsed := LettersUsed{1, 2}

	assert.Equal(t, lettersUsed.Total(), 6)
}

func TestShouldGetOnlyWords(t *testing.T) {
	lettersUsed := LettersUsed{1, 2}

	assert.Equal(t, lettersUsed.IntToWords(342), "threehundredfortytwo")
}

func TestOneToOneThousand(t *testing.T) {
	lettersUsed := LettersUsed{1, 1000}

	assert.Equal(t, lettersUsed.Total(), 18451)
}
