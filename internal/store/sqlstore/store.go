package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/store"
	"testing"

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

func TestDB(t *testing.T, databaseURL string) (store.Store, func(string)) {
	t.Helper()
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
	return &Store{db: db}, func(tables string) {
		if len(tables) > 0 {
			db.Exec("TRUNCATE %s CASCADE", tables)
		}
		db.Close()
	}
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
