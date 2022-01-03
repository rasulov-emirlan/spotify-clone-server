package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/models"
)

type GenreRepository struct {
	db *sql.DB
}

func (r *GenreRepository) Create(g *models.Genre, languageID int) error {
	return r.db.QueryRow(`
	WITH tt as (
		INSERT INTO genres (cover_picture_url)
		VALUES ($1)
		RETURNING id
	)
	INSERT INTO genres_localizations(genre_id, name, language_id)
	VALUES((SELECT id FROM tt), $2, $3)
	RETURNING genre_id;
	`, g.CoverURL, g.Name, languageID).Scan(&g.ID)
}

func (r *GenreRepository) AddSong(songID, genreID int) error {
	return r.db.QueryRow(`
	INSERT INTO songs_genres (song_id, genre_id)
	VALUES($1, $2);
	`, songID, genreID).Err()
}

func (r *GenreRepository) AddLocalization(genreID, languageID int, name string) error {
	return r.db.QueryRow(`
	INSERT INTO genres_localizations(name, genre_id, language_id)
	VALUES($1, $2, $3);
	`, name, genreID, languageID).Err()
}

func (r *GenreRepository) GetSongs(genderID int) ([]*models.Song, error) {
	rows, err := r.db.Query(`
	SELECT s.id, s.name, u.username, s.author_id, s.song_url, s.cover_picture_url
	FROM songs_genres AS sg
	INNER JOIN songs AS s
	ON s.id = sg.song_id
	INNER JOIN users AS u
	ON u.id = s.author_id
	WHERE sg.genre_id = $1;
	`, genderID)

	if err != nil {
		return nil, err
	}

	var (
		songs                         []*models.Song
		id, authorID                  int
		username, name, url, coverURL string
	)

	for rows.Next() {
		if err := rows.Scan(
			&id, &name, &username, &authorID, &url, &coverURL,
		); err != nil {
			return nil, err
		}

		songs = append(songs,
			&models.Song{
				ID:       id,
				Name:     name,
				URL:      url,
				CoverURL: coverURL,
				Author: models.User{
					ID:       authorID,
					UserName: username,
				},
			},
		)
	}

	return songs, nil
}

func (r *GenreRepository) ListAll() ([]*models.Genre, error) {
	rows, err := r.db.Query(`
	SELECT g.id, gl.name
	FROM genres g
	INNER JOIN genres_localizations gl
		ON g.id = gl.genre_id
	WHERE language_id = 1;
	`)

	if err != nil {
		return nil, err
	}

	var (
		genres []*models.Genre
		id     int
		name   string
	)

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&name,
		); err != nil {
			return nil, err
		}
		genres = append(genres, &models.Genre{
			ID:   id,
			Name: name,
		})
	}
	return genres, nil
}
