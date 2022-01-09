CREATE TABLE IF NOT EXISTS playlists_songs (
	song_id integer,
	playlist_id integer,
	created_at date DEFAULT CURRENT_DATE,
	CONSTRAINT fk_playlists_songs_song_id FOREIGN KEY(song_id) REFERENCES songs(id),
	CONSTRAINT fk_playlists_songst_playlist_id FOREIGN KEY(playlist_id) REFERENCES playlists(id)
);

CREATE INDEX IF NOT EXISTS idx_playlists_songs
ON playlists_songs(song_id, playlist_id);