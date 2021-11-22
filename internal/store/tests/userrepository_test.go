package teststore

import (
	"spotify-clone/server/internal/models"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	tstore, destructor, close := New(t)
	defer destructor("users")
	destructor("users")
	defer close()
	user := models.TestUser()
	_, err := tstore.User().Create(user)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_FindByID(t *testing.T) {
	tstore, destructor, close := New(t)
	defer destructor("users")
	destructor("users")
	defer close()
	user := models.TestUser()
	id, err := tstore.User().Create(user)
	if err != nil {
		t.Error(err)
	}
	user, err = tstore.User().FindByID(id)
	if err != nil {
		t.Error(err)
	}
	if user == nil {
		t.Error("user is nil")
	}
}
