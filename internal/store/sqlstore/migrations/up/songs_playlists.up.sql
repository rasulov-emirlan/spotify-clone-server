create table if not exists songs_playlists (
	song_id integer,
	playlist_id integer,
	CONSTRAINT fk_songs_playlists_song_id FOREIGN KEY(song_id) REFERENCES songs(id),
	CONSTRAINT fk_songs_playlistst_playlist_id FOREIGN KEY(playlist_id) REFERENCES playlists(id)
);