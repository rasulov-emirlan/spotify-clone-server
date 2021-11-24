create table if not exists songs(
	id serial,
	title varchar(1000) UNIQUE not null,
	author_id integer,
	url text,
	CONSTRAINT songs_pkey PRIMARY KEY (id),
	CONSTRAINT fk_songs_author_id FOREIGN KEY(author_id) REFERENCES users(id)
);