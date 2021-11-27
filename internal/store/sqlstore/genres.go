package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/models"
)

type GenreRepository struct {
	db *sql.DB
}

func (r *GenreRepository) Create(g *models.Genre) error {
	return r.db.QueryRow(`
	INSERT INTO genres (name)
	VALUES($1) RETURNING id;
	`, g.Name).Scan(&g.ID)
}

func (r *GenreRepository) AddSong(songID, genreID int) error {
	return r.db.QueryRow(`
	INSERT INTO songs_genres (song_id, genre_id)
	VALUES($1, $2);
	`, songID, genreID).Err()
}

func (r *GenreRepository) GetSongs(genderID int) ([]models.Song, error) {
	rows, err := r.db.Query(`
	SELECT s.id, s.title, u.name AS username, s.author_id, s.url, sc.url AS coverurl
	FROM songs_genres AS sg
	INNER JOIN songs AS s
	ON s.id = sg.song_id
	INNER JOIN users AS u
	ON u.id = s.author_id
	INNER JOIN songs_covers AS sc
	ON sc.song_id = s.id
	WHERE sg.genre_id = $1;
	`, genderID)

	if err != nil {
		return nil, err
	}

	var songs []models.Song

	var id, authorID int
	var username, title, url, coverURL string

	for rows.Next() {
		if err := rows.Scan(
			&id, &title, &username, &authorID, &url, &coverURL,
		); err != nil {
			return nil, err
		}

		songs = append(songs,
			models.Song{
				ID:       id,
				Title:    title,
				URL:      url,
				CoverURL: coverURL,
				Author: models.User{
					ID:   authorID,
					Name: username,
				},
			},
		)
	}

	return songs, nil
}

func (r *GenreRepository) ListAll() ([]models.Genre, error) {
	rows, err := r.db.Query(`
	SELECT id, name
	FROM genres;
	`)

	if err != nil {
		return nil, err
	}

	var genres []models.Genre
	var id int
	var name string

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&name,
		); err != nil {
			return nil, err
		}
		genres = append(genres, models.Genre{
			ID:   id,
			Name: name,
		})
	}
	return genres, nil
}
