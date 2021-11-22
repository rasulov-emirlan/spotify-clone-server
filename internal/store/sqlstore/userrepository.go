package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) Create(u *models.User) error {
	return nil
}

func (r *UserRepository) FindByID(id int) (*models.User, error) {
	return nil, nil
}
