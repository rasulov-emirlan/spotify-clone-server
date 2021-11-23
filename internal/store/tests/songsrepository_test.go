package teststore

import (
	"log"
	"spotify-clone/server/config"
	"spotify-clone/server/internal/models"
	"spotify-clone/server/internal/store/sqlstore"
	"testing"
)

func TestSongRepository_Create(t *testing.T) {
	config, err := config.NewTESTSQLDBconfig()
	if err != nil {
		t.Error(err)
	}
	tstore, destructor := sqlstore.TestDB(t, config, "songs")
	defer destructor()
	song := models.TestSong()
	err = tstore.Song().Create(song)
	if err != nil {
		t.Error(err)
	}
}

func TestSongRepository_FindByID(t *testing.T) {
	config, err := config.NewTESTSQLDBconfig()
	if err != nil {
		t.Error(err)
	}
	tstore, destructor := sqlstore.TestDB(t, config, "songs")
	defer destructor()
	song := models.TestSong()
	err = tstore.Song().Create(song)
	if err != nil {
		t.Error(err)
	}
	log.Println(1)
	song, err = tstore.Song().FindByID(song.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println(2)
	if song == models.TestSong() {
		t.Error(err)
	}
}
