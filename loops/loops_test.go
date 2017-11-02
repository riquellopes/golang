package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_number_1(t *testing.T) {
	result := Calculate(1)

	assert.Equal(t, result, 1.0)
}

func Test_number_4(t *testing.T) {
	result := Calculate(4)

	assert.Equal(t, result, 16.0)
}
