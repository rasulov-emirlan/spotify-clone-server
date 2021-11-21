package store

import (
	"database/sql"
	"spotify-clone/server/config"

	_ "github.com/lib/pq"
)

type Store struct {
	db             *sql.DB
	songRepository *SongRepository
}

func New() (*Store, error) {
	s, err := config.NewDBConfig()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", s)
	if err != nil {
		return nil, err
	}
	return &Store{
		db: db,
	}, nil
}

func (s *Store) Song() *SongRepository {
	if s.songRepository != nil {
		return s.songRepository
	}
	return &SongRepository{
		db: s.db,
	}
}
