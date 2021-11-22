package teststore

import (
	"spotify-clone/server/internal/models"
	"testing"
)

func TestSongRepository_Create(t *testing.T) {
	tstore, destructor := New(t)
	defer destructor("songs")
	song := models.TestSong()
	_, err := tstore.Song().Create(song)
	if err != nil {
		t.Error(err)
	}
}

func TestSongRepository_FindByID(t *testing.T) {
	tstore, destructor := New(t)
	defer destructor("songs")
	song := models.TestSong()
	id, err := tstore.Song().Create(song)
	if err != nil {
		t.Error(err)
	}
	song, err = tstore.Song().FindByID(id)
	if err != nil {
		t.Error(err)
	}
	if song == models.TestSong() {
		t.Error(err)
	}
}
