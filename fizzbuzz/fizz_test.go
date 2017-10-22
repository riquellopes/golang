package fizzbuzz

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFizzBuzz_0(t *testing.T) {
    fizzbuzz := FizzBuzz(0)
    assert.Equal(t, fizzbuzz, "0")
}

func TestFizzBuzz_1(t *testing.T) {
    fizzbuzz := FizzBuzz(1)
    assert.Equal(t, fizzbuzz, "1")
}

func TestFizzBuzz_3(t *testing.T) {
    fizzbuzz := FizzBuzz(3)
    assert.Equal(t, fizzbuzz, "fizz")
}

func TestFizzBuzz_9(t *testing.T) {
    fizzbuzz := FizzBuzz(9)
    assert.Equal(t, fizzbuzz, "fizz")
}

func TestFizzBuzz_5(t *testing.T) {
    fizzbuzz := FizzBuzz(5)
    assert.Equal(t, fizzbuzz, "buzz")
}

func TestFizzBuzz_25(t *testing.T) {
    fizzbuzz := FizzBuzz(25)
    assert.Equal(t, fizzbuzz, "buzz")
}

func TestFizzBuzz_15(t *testing.T) {
    fizzbuzz := FizzBuzz(15)
    assert.Equal(t, fizzbuzz, "fizzbuzz")
}
