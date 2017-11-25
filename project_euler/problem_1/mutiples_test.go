package problemOne

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGet23(t *testing.T) {
	assert.Equal(t, Multiple(10), 23)
}

func TestShouldGet14(t *testing.T) {
	assert.Equal(t, Multiple(9), 14)
}

func TestShouldGet233168(t *testing.T) {
	assert.Equal(t, Multiple(1000), 233168)
}
