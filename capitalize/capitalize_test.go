package capitalize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCap_1(t *testing.T) {
	cap := Capitalize("hello world")

	assert.Equal(t, cap, "Hello World")
}

func TestCap_2(t *testing.T) {
	cap := Capitalize("12abc")

	assert.Equal(t, cap, "12abc")
}

func TestCap_3(t *testing.T) {
	cap := Capitalize("1 2 3 4 5 6 8")

	assert.Equal(t, cap, "1 2 3 4 5 6 8")
}

func TestCap_4(t *testing.T) {
	cap := Capitalize("q w e r t y u i o p a s d f g h j  k l z x c v b n m Q W E R T Y U I O P A S D F G H J  K L Z X C V B N M")

	assert.Equal(t, cap, "Q W E R T Y U I O P A S D F G H J  K L Z X C V B N M Q W E R T Y U I O P A S D F G H J  K L Z X C V B N M")
}
