package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/models"
)

type PlaylistRepository struct {
	db *sql.DB
}

func (r *PlaylistRepository) Create(p *models.Playlist) error {
	return r.db.QueryRow(`
	INSERT INTO playlists(
	name, user_id)
	VALUES ($1, $2);
	`, p.Name, p.Author.ID).Err()
}

func (r *PlaylistRepository) UsersPlaylists(userID int) ([]models.Playlist, error) {
	rows, err := r.db.Query(`
	SELECT p.id, p.name, u.name as username
	FROM playlists AS p
	INNER JOIN users AS u
	ON p.user_id = u.id
	WHERE user_id = $1;
	`, userID)

	if err != nil {
		return nil, err
	}

	var playlists []models.Playlist

	var id int
	var name, username string

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&name,
			&username,
		); err != nil {
			return nil, err
		}
		playlists = append(playlists, models.Playlist{
			ID:   id,
			Name: name,
			Author: models.User{
				ID:   userID,
				Name: username,
			},
		})
	}
	return playlists, nil
}

func (r *PlaylistRepository) AddSong(songID int, playlistID int) error {
	return r.db.QueryRow(`
	INSERT INTO songs_playlists(
	song_id, playlist_id)
	VALUES ($1, $2);
	`, songID, playlistID).Err()
}

func (r *PlaylistRepository) ListAll() ([]models.Playlist, error) {
	rows, err := r.db.Query(`
	SELECT p.id, p.name, p.user_id, u.name as username
	FROM playlists as p
	INNER JOIN users as u
	ON p.user_id = u.id
	;
	`)

	if err != nil {
		return nil, err
	}

	var playlists []models.Playlist
	var id, userID int
	var name, username string

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&name,
			&userID,
			&username,
		); err != nil {
			return nil, err
		}
		playlists = append(playlists, models.Playlist{
			ID:   id,
			Name: name,
			Author: models.User{
				ID:   id,
				Name: username,
			},
		})
	}
	return playlists, nil
}
func (r *PlaylistRepository) GetSongsFromPlaylist(playlistID int) (*[]models.Song, error) {
	rows, err := r.db.Query(`
	SELECT s.id, s.title, s.author_id, s.url, sc.url as cover_url, u.name as username
	FROM songs AS s
	INNER JOIN songs_covers AS sc
	ON sc.song_id = s.id
	INNER JOIN users as u
	ON s.author_id = u.id
	WHERE s.id IN(
		SELECT ps.song_id
		FROM songs_playlists AS ps
		WHERE ps.playlist_id = $1
	);
	`, playlistID)

	if err != nil {
		return nil, err
	}

	var songs []models.Song

	var id, authorId int
	var title, url, coverUrl, username string

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&title,
			&authorId,
			&url,
			&coverUrl,
			&username,
		); err != nil {
			return nil, err
		}
		songs = append(songs, models.Song{
			ID:       id,
			Title:    title,
			Author:   models.User{ID: authorId, Name: username},
			URL:      url,
			CoverURL: coverUrl})
	}

	return &songs, nil
}
