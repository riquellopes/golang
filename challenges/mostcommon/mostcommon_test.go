package mostcommon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet322forBac(t *testing.T) {
	m := &Most{Char: "aabbbccde"}
	m.Shake()

	cholate, _ := m.Chocolate()
	strawberry, _ := m.Strawberry()
	vanilla, _ := m.Vanilla()

	assert.Equal(t, cholate, Common{"b", 3})
	assert.Equal(t, strawberry, Common{"a", 2})
	assert.Equal(t, vanilla, Common{"c", 2})
}

func TestGet4forA(t *testing.T) {
	m := &Most{Char: "aaaa"}
	m.Shake()

	cholate, _ := m.Chocolate()
	_, strawberror := m.Strawberry()
	_, vanillerror := m.Vanilla()

	assert.Equal(t, cholate, Common{"a", 4})
	assert.NotNil(t, strawberror)
	assert.NotNil(t, vanillerror)
}
