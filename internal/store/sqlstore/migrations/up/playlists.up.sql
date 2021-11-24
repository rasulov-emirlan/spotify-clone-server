create table if not exists playlists (
	id serial PRIMARY KEY,
	name varchar(100) not null UNIQUE,
	user_id integer
);
