package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/store"
	"testing"
)

func TestDB(t *testing.T, databaseURL string, table string) (store.Store, func()) {
	t.Helper()
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
	db.Exec("TRUNCATE table %s cascade;", table)
	return &Store{
			db:             db,
			userRepository: &UserRepository{db: db},
			songRepository: &SongRepository{db: db},
		}, func() {

			db.Exec("TRUNCATE table %s cascade;", table)

			db.Close()
		}
}
