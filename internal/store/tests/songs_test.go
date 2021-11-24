package teststore

import (
	"spotify-clone/server/internal/models"
	"testing"
)

func TestSongRepository_Create(t *testing.T) {
	tstore, destructor := NewTEST(t, "songs")
	defer destructor()

	user := models.TestUser()
	err := tstore.User().Create(user)
	if err != nil {
		t.Error(err)
	}
	song := models.TestSong()
	song.Author.ID = user.ID
	err = tstore.Song().Create(song)
	if err != nil {
		t.Error(err)
	}
}

func TestSongRepository_FindByID(t *testing.T) {
	tstore, destructor := NewTEST(t, "songs")
	defer destructor()

	song := models.TestSong()
	err := tstore.Song().Create(song)
	if err != nil {
		t.Error(err)
	}
	song, err = tstore.Song().FindByID(song.ID)
	if err != nil {
		t.Error(err)
	}
	if song == models.TestSong() {
		t.Error(err)
	}
}
