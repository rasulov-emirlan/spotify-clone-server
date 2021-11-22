package store

import "spotify-clone/server/internal/models"

type Store interface {
	Song() SongRepository
	User() UserRepository
}

type SongRepository interface {
	Create(s *models.Song) error
	FindByID(id int) (*models.Song, error)
}

type UserRepository interface {
	Create(u *models.User) (int, error)
	FindByID(id int) (*models.User, error)
}
