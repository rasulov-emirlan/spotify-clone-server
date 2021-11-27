package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/models"
)

type SongRepository struct {
	db *sql.DB
}

func (r *SongRepository) Create(s *models.Song) error {
	return r.db.QueryRow(`
	WITH Y as (INSERT INTO songs(title, author_id, url)
	VALUES($1, $2, $3)
	RETURNING id)
	INSERT INTO songs_covers(song_id, url)
	SELECT id, $4
	FROM Y
	RETURNING song_id;
	`, s.Title, s.Author.ID, s.URL, s.CoverURL).Scan(&s.ID)
}

func (r *SongRepository) FindByID(id int) (*models.Song, error) {
	rows, err := r.db.Query(`
	SELECT id, title, author_id, url
	FROM songs
	WHERE id = $1;
	`, id)
	if err != nil {
		return nil, err
	}

	var result models.Song
	if rows.Next() {
		rows.Scan(
			&result.ID,
			&result.Title,
			&result.Author.ID,
			&result.URL,
		)
	}
	return &result, nil
}

func (song *SongRepository) DeleteByID(id int) error {
	return nil
}

func (r *SongRepository) GetSongs(from, to int) ([]models.Song, error) {
	limit := to - from
	rows, err := r.db.Query(`
	SELECT  s.id, s.title, s.url, sc.url as cover_url, s.author_id, u.name as username
	FROM songs as s
	INNER JOIN users as u
	ON u.id = s.author_id
	INNER JOIN songs_covers as sc
		ON sc.song_id = s.id
		Limit $1 offset $2;
	`, limit, from)

	if err != nil {
		return nil, err
	}

	var songs []models.Song
	var id, userid int
	var title, username, url, coverUrl string

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&title,
			&url,
			&coverUrl,
			&userid,
			&username,
		); err != nil {
			return nil, err
		}
		songs = append(songs, models.Song{
			ID:       id,
			Title:    title,
			URL:      url,
			CoverURL: coverUrl,
			Author: models.User{
				ID:   userid,
				Name: username,
			},
		})
	}
	return songs, nil
}
