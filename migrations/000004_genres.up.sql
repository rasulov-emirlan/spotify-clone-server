CREATE TABLE IF NOT EXISTS genres (
	id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
	cover_picture_url varchar(500),
	CONSTRAINT pk_genre_id PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_genres_id
ON genres(id);