package teststore

import (
	"spotify-clone/server/config"
	"spotify-clone/server/internal/models"
	"spotify-clone/server/internal/store/sqlstore"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	config, err := config.NewTESTSQLDBconfig()
	if err != nil {
		t.Error(err)
	}
	tstore, destructor := sqlstore.TestDB(t, config, "users")
	defer destructor()

	user := models.TestUser()
	err = tstore.User().Create(user)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestUserRepository_FindByID(t *testing.T) {
	config, err := config.NewTESTSQLDBconfig()
	if err != nil {
		t.Error(err)
		return
	}
	tstore, destructor := sqlstore.TestDB(t, config, "users")
	defer destructor()

	user := models.TestUser()
	user, err = tstore.User().FindByID(user.ID)
	if err != nil {
		t.Error(err)
		return
	}
}
