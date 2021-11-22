package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) Create(u *models.User) (int, error) {
	row, err := r.db.Query(`
	insert into users (name, password, email)
	values($1, $2, $3) returning id;
	`, u.Name, u.Password, u.Email)

	var id int

	if row.Next() {
		row.Scan(&id)
	}
	return id, err
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
