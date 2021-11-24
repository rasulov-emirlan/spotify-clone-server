create table if not exists users_avatars (
	user_id integer UNIQUE,
	url text not null,
	CONSTRAINT fk_users_avatars_user_id FOREIGN KEY(user_id) REFERENCES users(id)
);