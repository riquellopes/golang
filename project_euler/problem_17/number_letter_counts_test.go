package problem17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneToFive(t *testing.T) {
	letters_used := LettersUsed{1, 5}

	assert.Equal(t, letters_used.Total(), 19)
}
