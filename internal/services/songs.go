package services

import "spotify-clone/server/internal/models"

type SongsRepository interface {
	Create() error
	FindByID(id int) (*models.Song, error)
	FindByGenre(genre string) (*[]models.Song, error)
	FindByPlaylist(playlist *models.Playlist) (*[]models.Song, error)
}
