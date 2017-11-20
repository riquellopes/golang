package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberthree(t *testing.T) {
	sequencie := prints(3)

	assert.Equal(t, sequencie, "123")
}

func TestNumberfive(t *testing.T) {
	sequencie := prints(5)

	assert.Equal(t, sequencie, "12345")
}
