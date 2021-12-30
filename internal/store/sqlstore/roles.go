package sqlstore

import "database/sql"

type RolesRepository struct {
	db *sql.DB
}

func (r *RolesRepository) AddRole(userID, roleID int) error {
	_, err := r.db.Query(`
	INSERT INTO users_roles(user_id, role_id)
	VALUES($1, $2);`, userID, roleID)
	return err
}
