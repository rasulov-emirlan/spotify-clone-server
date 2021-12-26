CREATE TABLE IF NOT EXISTS songs(
	id INT GENERATED ALWAYS AS IDENTITY,
	name varchar(500) NOT NULL,
	author_id integer NOT NULL,
	length integer,
	cover_picture_url varchar(500),
	song_url varchar(500),
	CONSTRAINT pk_songs_id PRIMARY KEY (id),
	CONSTRAINT ch_songs_length CHECK (length > 0),
	CONSTRAINT fk_songs_author_id FOREIGN KEY(author_id) REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_songs_id
ON songs(id);