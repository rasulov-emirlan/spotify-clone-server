CREATE TABLE IF NOT EXISTS favorite_authors (
	user_id integer NOT NULL,
	author_id integer NOT NULL,
	CONSTRAINT fk_favorite_authors_user_id FOREIGN KEY (user_id)
		REFERENCES users(id),
	CONSTRAINT fk_favorite_authors_author_id FOREIGN KEY (author_id)
		REFERENCES users(id),
	CONSTRAINT favorite_authors_user_id_cannot_be_author_id
		CHECK (user_id <> author_id)
);

CREATE INDEX IF NOT EXISTS idx_favorite_authors
ON favorite_authors(user_id, author_id);