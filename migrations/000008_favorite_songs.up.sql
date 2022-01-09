CREATE TABLE IF NOT EXISTS favorite_songs (
	song_id integer NOT NULL,
	user_id integer NOT NULL,
	created_at date DEFAULT CURRENT_DATE,
	CONSTRAINT fk_favorite_songs_song_id FOREIGN KEY (song_id)
		REFERENCES songs(id),
	CONSTRAINT fk_favorite_songs_user_id FOREIGN KEY (user_id)
		REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_favorite_songs
ON favorite_songs(song_id, user_id);