package teststore

import (
	"spotify-clone/server/internal/models"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	tstore := New()
	user := &models.User{}
	err := tstore.User().Create(user)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_FindByID(t *testing.T) {
	tstore := New()
	user := &models.User{}
	err := tstore.User().Create(user)
	if err != nil {
		t.Error(err)
	}
	user, err = tstore.User().FindByID(1)
	if user == nil {
		t.Error(err)
	}
}
