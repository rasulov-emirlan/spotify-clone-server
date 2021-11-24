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
	rows, err := r.db.Query(`
	select id, name, email
	from users
	where id=$1;
	`, id)
	if err != nil {
		return nil, err
	}

	var result *models.User
	if rows.Next() {
		rows.Scan(
			&result.ID,
			&result.Name,
			&result.Email,
		)
	}
	return result, nil
}

func (r *UserRepository) BanByID(id int) error {
	return nil
}
