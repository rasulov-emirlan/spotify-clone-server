package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/store"

	_ "github.com/lib/pq"
)

type Store struct {
	db             *sql.DB
	songRepository *SongRepository
	userRepository *UserRepository
}

func New(config string) (*Store, error) {
	db, err := sql.Open("postgres", config)
	if err != nil {
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
	if s.songRepository != nil {
		return s.userRepository
	}
	return &UserRepository{
		db: s.db,
	}
}
