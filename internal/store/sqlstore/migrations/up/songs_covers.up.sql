create table if not exists songs_covers (
	song_id int UNIQUE,
	url text not null,
	CONSTRAINT fk_songs_covers_song_id FOREIGN KEY(song_id) REFERENCES songs(id)
);