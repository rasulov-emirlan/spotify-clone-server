create table if not exists users(
	id serial,
	name varchar(100),
	password varchar(100),
	email varchar(120),
	 CONSTRAINT uk_email UNIQUE (email)
        INCLUDE(email)
);