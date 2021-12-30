CREATE TABLE IF NOT EXISTS countries (
	id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
	name varchar(255) NOT NULL,
	-- language_id integer NOT NULL,
	CONSTRAINT pk_countries_id PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS languages (
	id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
	name varchar(255) NOT NULL,
	-- country_id integer NOT NULL,
	CONSTRAINT pk_languages_id PRIMARY KEY(id)
);

-- ALTER TABLE countries
-- ADD CONSTRAINT fk_countries_languege_id
-- FOREIGN KEY (language_id)
-- REFERENCES languages(id);

-- ALTER TABLE languages
-- ADD CONSTRAINT fk_languages_country_id
-- FOREIGN KEY (country_id)
-- REFERENCES countries(id);

INSERT INTO countries(name)
VALUES('USA');

INSERT INTO languages(name)
VALUES('english');