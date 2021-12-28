CREATE TABLE IF NOT EXISTS users (
	id INT GENERATED ALWAYS AS IDENTITY UNIQUE,
	username varchar(100) NOT NULL,
	full_name varchar(200) NOT NULL,
	password varchar(256) NOT NULL,
	birth_date date,
	email varchar(120) NOT NULL UNIQUE,
	profile_picture_url varchar(500),
	country_code integer,
	language_id integer,
	CONSTRAINT pk_users_id PRIMARY KEY(id),
	CONSTRAINT fk_users_country_code FOREIGN KEY(country_code) REFERENCES countries(id),
	CONSTRAINT fk_users_language_code FOREIGN KEY(language_id) REFERENCES languages(id)
);

CREATE INDEX IF NOT EXISTS idx_users_id
ON users(id, email);