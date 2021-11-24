package teststore

import (
	"spotify-clone/server/internal/models"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	tstore, destructor := NewTEST(t)
	defer destructor()

	user := models.TestUser()
	err := tstore.User().Create(user)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_FindByID(t *testing.T) {
	tstore, destructor := NewTEST(t)
	defer destructor()
	user := models.TestUser()
	err := tstore.User().Create(user)
	if err != nil {
		t.Error(err)
	}
	user, err = tstore.User().FindByID(user.ID)
	if err != nil {
		t.Error(err)
	}
	if user == nil {
		t.Error("user is null")
	}
}
