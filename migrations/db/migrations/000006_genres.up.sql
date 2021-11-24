create table if not exists genres (
	id serial PRIMARY KEY,
	name varchar(255) UNIQUE not null
);