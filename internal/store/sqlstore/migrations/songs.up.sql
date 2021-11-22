create table if not exists songs(
	id serial,
	title varchar(1000),
	author_id integer,
	url text,
	CONSTRAINT songs_pkey PRIMARY KEY (id),
    CONSTRAINT uk_title UNIQUE (title)
        INCLUDE(title)
);