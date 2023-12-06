package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Geovane Cavalcante", "geovane@gmail.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Geovane Cavalcante", user.Name)
	assert.Equal(t, "geovane@gmail.com", user.Email)
	assert.NotEmpty(t, user.Password)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("Geovane Cavalcante", "geovane@gmail.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, "123456", user.Password)
}
