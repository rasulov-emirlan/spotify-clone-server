package store

import (
	"spotify-clone/server/internal/models"
)

type Store interface {
	Song() SongRepository
	User() UserRepository
	Playlist() PlaylistRepository
	Genre() GenresRepository
	Language() LanguageRepository
	Country() CountryRepository
}

type SongRepository interface {
	Create(s *models.Song) error
	FindByID(id int) (*models.Song, error)
	GetSongs(from int, to int) ([]*models.Song, error)
	// DeleteByID(id int) error // new functionalaty
}

type UserRepository interface {
	Create(u *models.User) error
	FindByID(id int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	AddFavoriteSong(songID, userID int) error
	ListFavoriteSongs(userID, limit, offset int) ([]*models.Song, error)
	RemoveFromFavoriteSongs(userID, songID int) error
	// BanByID(id int) error // new functionalaty
}

type PlaylistRepository interface {
	Create(p *models.Playlist) error
	ListAll() ([]*models.Playlist, error)
	UsersPlaylists(userID int) ([]*models.Playlist, error)
	AddSong(songID int, playlistID int) error
	GetSongsFromPlaylist(id int) ([]*models.Song, error)
}

type GenresRepository interface {
	Create(g *models.Genre, languageID int) error
	AddLocalization(genreID, languageID int, name string) error
	ListAll() ([]*models.Genre, error)
	AddSong(songID, genreID int) error
	GetSongs(genderID int) ([]*models.Song, error)
}

type LanguageRepository interface {
	Create(l *models.Language) error
	ListAll() ([]*models.Language, error)
}

type CountryRepository interface {
	Create(c *models.Country) error
	ListAll() ([]*models.Country, error)
	Delete(countryID int) error
}
