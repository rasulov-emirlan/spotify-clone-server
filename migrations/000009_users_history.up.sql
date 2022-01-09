CREATE TABLE IF NOT EXISTS users_history (
	song_id integer NOT NULL,
	user_id integer NOT NULL,
	created_at date DEFAULT CURRENT_DATE,
	CONSTRAINT fk_users_history_song_id FOREIGN KEY (song_id)
		REFERENCES songs(id),
	CONSTRAINT fk_users_history_user_id FOREIGN KEY (user_id)
		REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_users_history
ON users_history(user_id, song_id);