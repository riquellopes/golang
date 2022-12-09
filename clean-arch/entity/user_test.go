package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	u, _ := NewUser("Henrique Lopes", 40)
	assert.Equal(t, u.Name, "Henrique Lopes")
	assert.Equal(t, u.Age, 40)
}

func TestCreateAnInvalidUser(t *testing.T) {
	_, err := NewUser("Henrique Lopes", 0)
	assert.NotNil(t, err)
}
