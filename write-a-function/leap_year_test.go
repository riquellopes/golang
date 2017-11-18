package leapyear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_2000_is_true(t *testing.T) {
	assert.True(t, isLeap(2000))
}

func Test_2400_is_true(t *testing.T) {
	assert.True(t, isLeap(2400))
}

func Test_1800_is_true(t *testing.T) {
	assert.False(t, isLeap(1800))
}
