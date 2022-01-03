package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) Create(u *models.User) error {
	err := r.db.QueryRow(
		`
		INSERT INTO users (username, full_name, birth_date, password, email)
		VALUES($1, $2, $3, $4, $5) returning id;
		`,
		u.UserName,
		u.FullName,
		u.BirthDate,
		u.Password,
		u.Email,
	).Scan(&u.ID)
	u.Password = ""
	return err
}

func (r *UserRepository) FindByID(id int) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow(`
	SELECT id, username, password, email
	FROM users
		WHERE id=$1;
	`, id).Scan(&u.ID, &u.UserName, &u.Email)
	return &u, err
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow(`
	SELECT id, username, password, email
	FROM users
		WHERE email = $1;
	`, email).Scan(&u.ID, &u.UserName, &u.Password, &u.Email)
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Query(`
	SELECT r.name
	FROM users_roles ur
	INNER JOIN roles r
		ON ur.role_id = r.id
		WHERE ur.user_id = $1; 
	`, u.ID)
	if err != nil {
		return nil, err
	}
	var role string

	for rows.Next() {
		rows.Scan(
			&role,
		)
		u.Roles = append(u.Roles, role)
	}
	return &u, err
}

func (r *UserRepository) BanByID(id int) error {
	return nil
}

func (r *UserRepository) FindByName(name string) (*models.User, error) {
	var u models.User

	return &u, r.db.QueryRow(`
	SELECT id, username, profile_picture_url
	FROM users
		WHERE username @@ to_tsquery('english', $1)
	LIMIT 1;
	`, name).Scan(&u.ID, &u.UserName, &u.ProfilePictureURL)
}
