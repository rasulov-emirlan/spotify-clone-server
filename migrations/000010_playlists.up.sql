CREATE TABLE IF NOT EXISTS playlists (
	id INT GENERATED ALWAYS AS IDENTITY,
	name varchar(500) NOT NULL,
	author_id integer NOT NULL,
	is_album boolean DEFAULT FALSE NOT NULL,
	cover_picture_url varchar(500),
	created_at date DEFAULT CURRENT_DATE,
	real_creation_date date,
	CONSTRAINT pk_playlists_id PRIMARY KEY (id),
	CONSTRAINT fk_playlists_author_id FOREIGN KEY (id) REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_playlists_id
ON playlists(id);