package database

import (
	"testing"

	"github.com/GeovaneCavalcante/api-users/internal/entity"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("Geovane", "geovane@gmail.con", "123456")

	userDB := NewUser(db)
	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID.String()).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID.String(), userFound.ID.String())
	assert.Equal(t, user.Name, userFound.Name)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("Geovane", "geovane@gmail.com", "123456")

	userDB := NewUser(db)
	err = userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID.String(), userFound.ID.String())
	assert.Equal(t, user.Name, userFound.Name)
}
