CREATE TABLE IF NOT EXISTS users_roles (
	user_id integer NOT NULL,
	role_id integer NOT NULL,
	CONSTRAINT fk_users_roles_user_id FOREIGN KEY (user_id)
		REFERENCES users(id),
	CONSTRAINT fk_users_roles_role_id FOREIGN KEY (role_id)
		REFERENCES roles(id)
);

CREATE INDEX IF NOT EXISTS idx_users_roles
ON users_roles(user_id);