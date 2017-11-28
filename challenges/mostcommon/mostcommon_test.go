package mostcommon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet322forBac(t *testing.T) {
	m := MostCommon("aabbbccde")

	assert.Equal(t, m["b"], 3)
	assert.Equal(t, m["a"], 2)
	assert.Equal(t, m["c"], 2)
}

func TestGet4forA(t *testing.T) {
	m := MostCommon("aaaa")

	assert.Equal(t, m["a"], 4)
}
