package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func (u *UserRepository) Create() error {
	return nil
}

func (u *UserRepository) FindByID(id int) (*models.User, error) {
	return nil, nil
}
