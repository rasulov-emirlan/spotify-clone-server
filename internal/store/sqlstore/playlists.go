package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/models"
)

type PlaylistRepository struct {
	db *sql.DB
}

func (r *PlaylistRepository) Create(p *models.Playlist) error {
	return nil
}

func (r *PlaylistRepository) GetSongsFromPlaylist(id int) (*[]models.Song, error) {
	return nil, nil
}
