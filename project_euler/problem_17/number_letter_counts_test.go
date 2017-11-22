package problem17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneToFive(t *testing.T) {
	lettersUsed := LettersUsed{1, 5}

	assert.Equal(t, lettersUsed.Total(), 19)
}
