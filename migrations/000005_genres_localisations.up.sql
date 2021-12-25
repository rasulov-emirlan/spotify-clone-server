CREATE TABLE IF NOT EXISTS genres_localisations (
	genre_id integer NOT NULL,
	name varchar(500) NOT NULL,
	language_id integer NOT NULL,
	CONSTRAINT fk_genres_localisations_language_id FOREIGN KEY (language_id) REFERENCES languages(id)
);

CREATE INDEX IF NOT EXISTS idx_genres_localisations_id
ON genres_localisations(genre_id);