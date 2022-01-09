CREATE TABLE IF NOT EXISTS users_roles (
	user_id integer NOT NULL,
	role_id integer NOT NULL,
	created_at date DEFAULT CURRENT_DATE,
	approved_by integer,
	CONSTRAINT fk_users_roles_user_id FOREIGN KEY (user_id)
		REFERENCES users(id),
	CONSTRAINT fk_users_roles_role_id FOREIGN KEY (role_id)
		REFERENCES roles(id),
	CONSTRAINT fk_users_roles_approved_by FOREIGN KEY(approved_by)
		REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_users_roles
ON users_roles(user_id);

INSERT INTO users_roles(user_id, role_id)
VALUES(1, 1);