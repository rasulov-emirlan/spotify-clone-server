package teststore

import (
	"log"
	"spotify-clone/server/config"
	"spotify-clone/server/internal/store/sqlstore"
)

type Store struct {
	store *sqlstore.Store
}

func New() *sqlstore.Store {
	config, err := config.NewTESTSQLDBconfig()
	if err != nil {
		log.Fatal(err)
	}
	s, err := sqlstore.New(config)
	if err != nil {
		log.Fatal(err)
	}
	return s
}
