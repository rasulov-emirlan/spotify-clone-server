package teststore

import (
	"spotify-clone/server/internal/models"
	"testing"
)

func TestPlaylistsRepository_Create(t *testing.T) {
	tstore, cleaner := NewTEST(t)
	defer cleaner()

	playlist := models.TestPlaylist()
	if err := tstore.Playlist().Create(playlist); err != nil {
		t.Error(err)
	}
}
