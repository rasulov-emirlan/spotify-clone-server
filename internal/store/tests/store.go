package teststore

import (
	"log"
	"spotify-clone/server/config"
	"spotify-clone/server/internal/store"
	"spotify-clone/server/internal/store/sqlstore"
	"testing"
)

type Store struct {
	store *sqlstore.Store
}

func New(t *testing.T) (store.Store, func(string)) {
	config, err := config.NewTESTSQLDBconfig()
	if err != nil {
		log.Fatal(err)
	}

	return sqlstore.TestDB(t, config)
}
