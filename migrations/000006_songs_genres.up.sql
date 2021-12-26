CREATE TABLE IF NOT EXISTS songs_genres (
	song_id integer,
	genre_id integer,
	CONSTRAINT fk_songs_genres_song_id FOREIGN KEY(song_id) REFERENCES songs(id),
	CONSTRAINT fk_songs_genres_genre_id FOREIGN KEY(genre_id) REFERENCES genres(id)
);

CREATE INDEX IF NOT EXISTS idx_songs_genres
ON songs_genres(song_id, genre_id);