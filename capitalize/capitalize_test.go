package capitalize

import (
    "testing"
    "github.com/stretchr/testify/assert"
)


func TestCap_1(t *testing.T){
    cap := Capitalize("hello world")

    assert.Equal(t, cap, "Hello World")
}
