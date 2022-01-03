package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/store"

	_ "github.com/lib/pq"
)

type Store struct {
	db                 *sql.DB
	songRepository     *SongRepository
	userRepository     *UserRepository
	playlistRepository *PlaylistRepository
	genreRepository    *GenreRepository
	languageRepository *LanguageRepository
	countryRepository  *CountryRepository
}

func New(config string) (*Store, error) {
	db, err := sql.Open("postgres", config)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Store{
		db: db,
	}, nil
}

func (s *Store) Song() store.SongRepository {
	if s.songRepository != nil {
		return s.songRepository
	}
	return &SongRepository{
		db: s.db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	return &UserRepository{
		db: s.db,
	}
}

func (s *Store) Playlist() store.PlaylistRepository {
	if s.playlistRepository != nil {
		return s.playlistRepository
	}
	return &PlaylistRepository{
		db: s.db,
	}
}

func (s *Store) Genre() store.GenresRepository {
	if s.genreRepository != nil {
		return s.genreRepository
	}
	return &GenreRepository{
		db: s.db,
	}
}

func (s *Store) Language() store.LanguageRepository {
	if s.languageRepository != nil {
		return s.languageRepository
	}
	return &LanguageRepository{
		db: s.db,
	}
}

func (s *Store) Country() store.CountryRepository {
	if s.countryRepository != nil {
		return s.countryRepository
	}
	return &CountryRepository{
		db: s.db,
	}
}
