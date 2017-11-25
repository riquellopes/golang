package problemOne

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGet23(t *testing.T) {
	assert.Equal(t, Multiple(3, 5, 10), 23)
}
