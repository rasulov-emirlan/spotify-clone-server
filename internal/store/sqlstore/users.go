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
	insert into users (username, full_name, birth_date, password, email)
	values($1, $2, $3, $4, $5) returning id;
	`, u.UserName, u.FullName, u.BirthDate, u.Password, u.Email).Scan(&u.ID)
}

func (r *UserRepository) FindByID(id int) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow(`
	select id, username, password, email
	from users
	where id=$1;
	`, id).Scan(&u.ID, &u.UserName, &u.Email)
	return &u, err
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow(`
	select id, username, password, email
	from users
	where email = $1;
	`, email).Scan(&u.ID, &u.UserName, &u.Password, &u.Email)
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Query(`
	SELECT r.id, r.name
	FROM users_roles ur
	INNER JOIN roles r
		ON ur.role_id = r.id
	WHERE ur.user_id = $1; 
	`, u.ID)
	if err != nil {
		return nil, err
	}
	var roleid int
	var role string

	for rows.Next() {
		rows.Scan(
			&role,
			&roleid,
		)
		u.Roles = append(u.Roles, &models.Role{
			ID:   roleid,
			Name: role,
		})
	}
	return &u, err
}

func (r *UserRepository) BanByID(id int) error {
	return nil
}
