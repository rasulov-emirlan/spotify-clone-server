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
	INSERT INTO songs(name, author_id, cover_picture_url, song_url)
	VALUES($1, $2, $3, $4)
	RETURNING id;
	`, s.Name, s.Author.ID, s.CoverURL, s.URL).Scan(&s.ID)
}

func (r *SongRepository) FindByID(id int) (*models.Song, error) {
	rows, err := r.db.Query(`
	SELECT id, name, author_id, song_url
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
			&result.Name,
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
	SELECT  s.id, s.name, s.song_url, s.cover_picture_url as cover_url, s.author_id, u.username
	FROM songs as s
	INNER JOIN users as u
	ON u.id = s.author_id
		Limit $1 offset $2;
	`, limit, from)

	if err != nil {
		return nil, err
	}

	var songs []models.Song
	var id, userid int
	var name, username, url, coverUrl string

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&name,
			&url,
			&coverUrl,
			&userid,
			&username,
		); err != nil {
			return nil, err
		}
		songs = append(songs, models.Song{
			ID:       id,
			Name:     name,
			URL:      url,
			CoverURL: coverUrl,
			Author: models.User{
				ID:       userid,
				UserName: username,
			},
		})
	}
	return songs, nil
}
