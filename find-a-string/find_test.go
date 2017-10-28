package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_should_get_2(t *testing.T) {
	response := CountSubstring("ABCDCDC", "CDC")
	assert.Equal(t, response, 2)
}

func Test_should_get_3(t *testing.T) {
	response := CountSubstring("ABCDCDCAAAACDC", "CDC")
	assert.Equal(t, response, 3)
}
