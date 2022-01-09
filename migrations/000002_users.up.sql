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

INSERT INTO users(username, full_name, password, birth_date, email, profile_picture_url)
VALUES('god', 'the one above all', '1', '2022-01-09', 'none', 'there is no picture');

-- we could not create constaint for added_by in countries
-- and languages because we did not have any users
-- so we create them now
-- countries and languages could not be created
-- without users to create them
-- so we create them here

ALTER TABLE countries
ADD CONSTRAINT fk_countries_added_by FOREIGN KEY(added_by)
		REFERENCES users(id);

ALTER TABLE languages
ADD CONSTRAINT fk_languages_added_by FOREIGN KEY(added_by)
		REFERENCES users(id);

INSERT INTO countries(name, added_by)
VALUES('USA', 1);

INSERT INTO languages(name, added_by)
VALUES('english', 1);