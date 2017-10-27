package main

import (
      "testing"
      "github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
  a, b, c := Arithmetic{3, 2}

  assert.Equal(t, a, 5)
  assert.Equal(t, b, 1)
  assert.Equal(t, c, 6)
}
