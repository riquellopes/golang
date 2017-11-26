package main

import (
      "testing"
      "github.com/stretchr/testify/assert"
)

func TestCalc_1(t *testing.T) {
  ari := Arithmetic{3, 2}
  a, b, c := ari.ToCalc()

  assert.Equal(t, a, 5)
  assert.Equal(t, b, 1)
  assert.Equal(t, c, 6)
}

func TestCalc_2(t *testing.T) {
  ari := Arithmetic{5, 2}
  a, b, c := ari.ToCalc()

  assert.Equal(t, a, 7)
  assert.Equal(t, b, 3)
  assert.Equal(t, c, 10)
}
