package store

import (
	"spotify-clone/server/internal/models"
)

type Store interface {
	Song() SongRepository
	User() UserRepository
	Playlist() PlaylistRepository
}

type SongRepository interface {
	Create(s *models.Song) error
	FindByID(id int) (*models.Song, error)
	// GetSongs(from int, to int) (*[]models.Song, error)
	// DeleteByID(id int) error // new functionalaty
}

type UserRepository interface {
	Create(u *models.User) error
	FindByID(id int) (*models.User, error)
	// BanByID(id int) error // new functionalaty
}

type PlaylistRepository interface {
	Create(p *models.Playlist) error
	GetSongsFromPlaylist(id int) (*[]models.Song, error)
}
