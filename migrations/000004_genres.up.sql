CREATE TABLE IF NOT EXISTS genres (
	id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
	added_by integer NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
	cover_picture_url varchar(500),
	CONSTRAINT pk_genre_id PRIMARY KEY (id),
	CONSTRAINT fk_genres_added_by FOREIGN KEY(added_by)
		REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_genres_id
ON genres(id);