package teststore

import (
	"spotify-clone/server/internal/models"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	tstore, destructor := NewTEST(t, "users")
	defer destructor()

	user := models.TestUser()
	err := tstore.User().Create(user)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_FindByID(t *testing.T) {
	tstore, destructor := NewTEST(t, "songs")
	defer destructor()

	user := models.TestUser()
	user, err := tstore.User().FindByID(user.ID)
	if err != nil {
		t.Error(err)
	}
}
