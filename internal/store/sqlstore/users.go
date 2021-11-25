package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) Create(u *models.User) error {
	return r.db.QueryRow(`
	insert into users (name, password, email)
	values($1, $2, $3) returning id;
	`, u.Name, u.Password, u.Email).Scan(&u.ID)
}

func (r *UserRepository) FindByID(id int) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow(`
	select id, name, email
	from users
	where id=$1;
	`, id).Scan(&u.ID, &u.Name, &u.Email)
	return &u, err
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow(`
	select id, name, password, email
	from users
	where email = $1;
	`, email).Scan(&u.ID, &u.Name, &u.Password, &u.Email)
	return &u, err
}

func (r *UserRepository) BanByID(id int) error {
	return nil
}
