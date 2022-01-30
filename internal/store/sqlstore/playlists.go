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
	name, author_id, cover_picture_url, is_album)
	VALUES ($1, $2, $3, $4);
	`, p.Name, p.Author.ID, p.CoverUrl, p.IsAlbum).Err()
}

func (r *PlaylistRepository) UsersPlaylists(userID int) ([]*models.Playlist, error) {
	rows, err := r.db.Query(`
	SELECT p.id, p.name, u.username, p.cover_picture_url, is_album
	FROM playlists AS p
	INNER JOIN users AS u
	ON p.author_id = u.id
	WHERE author_id = $1;
	`, userID)

	if err != nil {
		return nil, err
	}

	var (
		playlists                []*models.Playlist
		id                       int
		name, username, coverUrl string
		isAlbum                  bool
	)

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&name,
			&username,
			&coverUrl,
			&isAlbum,
		); err != nil {
			return nil, err
		}
		playlists = append(playlists, &models.Playlist{
			ID:   id,
			Name: name,
			Author: models.User{
				ID:       userID,
				UserName: username,
			},
			CoverUrl: coverUrl,
			IsAlbum:  isAlbum,
		})
	}
	return playlists, nil
}

func (r *PlaylistRepository) AddSong(songID int, playlistID int) error {
	return r.db.QueryRow(`
	INSERT INTO playlists_songs(
	song_id, playlist_id)
	VALUES ($1, $2);
	`, songID, playlistID).Err()
}

func (r *PlaylistRepository) ListAll() ([]*models.Playlist, error) {
	rows, err := r.db.Query(`
	SELECT p.id, p.name, p.author_id, u.username, is_album
	FROM playlists as p
	INNER JOIN users as u
	ON p.author_id = u.id
	;
	`)

	if err != nil {
		return nil, err
	}

	var (
		playlists      []*models.Playlist
		id, userID     int
		name, username string
		isAlbum        bool
	)

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&name,
			&userID,
			&username,
			&isAlbum,
		); err != nil {
			return nil, err
		}
		playlists = append(playlists, &models.Playlist{
			ID:   id,
			Name: name,
			Author: models.User{
				ID:       id,
				UserName: username,
			},
			IsAlbum: isAlbum,
		})
	}
	return playlists, nil
}

func (r *PlaylistRepository) GetSongsFromPlaylist(userID, playlistID int) ([]*models.Song, error) {
	rows, err := r.db.Query(`WITH tt AS (
		SELECT song_id, COUNT(*) as like_count
		FROM favorite_songs
		GROUP BY song_id
	)
	SELECT s.id, s.name, s.author_id, s.song_url, COALESCE(tt.like_count, 0), 
	CASE
		WHEN  NOT EXISTS (
		SELECT TRUE
		FROM favorite_songs fs
			WHERE fs.user_id = $1
			AND fs.song_id = s.id
			LIMIT 1
		) THEN TRUE
		ELSE FALSE
	END
	is_liked,
	s.cover_picture_url AS cover_url, u.username
	FROM songs AS s
	INNER JOIN users as u
		ON s.author_id = u.id
	LEFT JOIN tt
		ON tt.song_id = s.id
		WHERE s.id IN(
			SELECT ps.song_id
			FROM playlists_songs AS ps
			WHERE ps.playlist_id = $2
		);`, userID, playlistID)

	if err != nil {
		return nil, err
	}

	var (
		songs                         []*models.Song
		id, authorId, likeCount       int
		name, url, coverUrl, username string
		isLiked                       bool
	)

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&name,
			&authorId,
			&url,
			&likeCount,
			&isLiked,
			&coverUrl,
			&username,
		); err != nil {
			return nil, err
		}
		songs = append(songs, &models.Song{
			ID:        id,
			Name:      name,
			Author:    models.User{ID: authorId, UserName: username},
			URL:       url,
			CoverURL:  coverUrl,
			LikeCount: likeCount,
			IsLiked:   isLiked,
		})
	}

	return songs, nil
}

func (r *PlaylistRepository) DeletePlaylist(playlistID int) error {
	return r.db.QueryRow(`
	DELETE FROM playlists
		WHERE id = $1;
	`, playlistID).Scan()
}
