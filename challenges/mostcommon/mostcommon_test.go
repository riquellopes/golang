package mostcommon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet322forBac(t *testing.T) {
	m := &Most{Char: "aabbbccde"}
	m.Shake()

	assert.Equal(t, m.Chocolate(), Common{"b", 3})
	assert.Equal(t, m.Strawberry(), Common{"a", 2})
	assert.Equal(t, m.Vanilla(), Common{"c", 2})
}

func TestGet4forA(t *testing.T) {
	m := &Most{Char: "aaaa"}
	m.Shake()

	assert.Equal(t, m.Chocolate(), Common{"a", 4})
}
