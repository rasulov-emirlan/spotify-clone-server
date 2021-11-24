create table if not exists songs_genres (
	song_id integer,
	genre_id integer,
	CONSTRAINT fk_songs_genres_song_id FOREIGN KEY(song_id) REFERENCES songs(id),
	CONSTRAINT fk_songs_genres_genre_id FOREIGN KEY(genre_id) REFERENCES genres(id)
);