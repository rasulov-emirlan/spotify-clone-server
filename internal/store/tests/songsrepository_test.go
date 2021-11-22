package teststore

import (
	"spotify-clone/server/internal/models"
	"testing"
)

func TestSongRepository_Create(t *testing.T) {
	tstore := New()
	song := &models.Song{}
	err := tstore.Song().Create(song)
	if err != nil {
		t.Error(err)
	}
}

func TestSongRepository_FindByID(t *testing.T) {
	tstore := New()
	song, err := tstore.Song().FindByID(1)
	if err != nil {
		t.Error(err)
	}
	if song == nil {
		t.Error(err)
	}
}
