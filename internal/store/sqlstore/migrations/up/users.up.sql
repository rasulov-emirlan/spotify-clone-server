create table if not exists users(
	id serial UNIQUE,
	name varchar(100) UNIQUE,
	password varchar(100),
	email varchar(120) UNIQUE
);