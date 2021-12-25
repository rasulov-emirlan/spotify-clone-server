CREATE TABLE IF NOT EXISTS denied_songs (
	song_id integer NOT NULL,
	user_id integer NOT NULL,
	CONSTRAINT fk_denied_songs_song_id FOREIGN KEY (song_id)
		REFERENCES songs(id),
	CONSTRAINT fk_denied_songs_user_id FOREIGN KEY (user_id)
		REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_denied_songs
ON denied_songs(user_id, song_id);