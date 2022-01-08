CREATE TABLE IF NOT EXISTS users_roles_applications (
	user_id integer NOT NULL,
	role_id integer NOT NULL,
	motivation_letter varchar(1500),
	application_date timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_users_roles_applications_user_id FOREIGN KEY(user_id)
		REFERENCES users(id),
	CONSTRAINT fk_users_roles_applications_role_id FOREIGN KEY(role_id)
		REFERENCES roles(id)
);