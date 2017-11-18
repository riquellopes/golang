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

func Test_1900_is_true(t *testing.T) {
	assert.False(t, isLeap(1900))
}

func Test_2100_is_true(t *testing.T) {
	assert.False(t, isLeap(2100))
}

func Test_2200_is_true(t *testing.T) {
	assert.False(t, isLeap(2200))
}

func Test_2300_is_true(t *testing.T) {
	assert.False(t, isLeap(2300))
}

func Test_2500_is_true(t *testing.T) {
	assert.False(t, isLeap(2500))
}
