package sqlstore

import (
	"database/sql"
	"log"
	"spotify-clone/server/internal/store"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func TestDB(t *testing.T, databaseURL string) (store.Store, func()) {
	t.Helper()
	m, err := migrate.New(
		"file://../../../migrations/",
		databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Down(); err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	m.Close()

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return &Store{
			db:             db,
			userRepository: &UserRepository{db: db},
			songRepository: &SongRepository{db: db},
		}, func() {
			db.Close()
			m, err := migrate.New(
				"file://../../../migrations",
				databaseURL)
			if err != nil {
				log.Fatal(err)
			}
			if err := m.Down(); err != nil {
				log.Fatal(err)
			}
			m.Close()
		}
}
