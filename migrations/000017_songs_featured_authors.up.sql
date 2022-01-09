CREATE TABLE IF NOT EXISTS songs_featured_authors (
	song_id integer NOT NULL,
	user_id integer NOT NULL,
	CONSTRAINT fk_featured_authors_song_id FOREIGN KEY (song_id)
		REFERENCES songs(id),
	CONSTRAINT fk_featured_authors_user_id FOREIGN KEY (user_id)
		REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_songs_featured_authors
ON songs_featured_authors(song_id, user_id);