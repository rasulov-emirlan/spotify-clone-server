CREATE TABLE IF NOT EXISTS users_playlists_history (
	playlist_id integer,
	user_id integer,
	created_at date DEFAULT CURRENT_DATE,
	CONSTRAINT fk_users_playlists_history_playlist_id FOREIGN KEY (playlist_id)
		REFERENCES playlists(id),
	CONSTRAINT fk_users_playlists_history_user_id FOREIGN KEY (user_id)
		REFERENCES users(id)
);