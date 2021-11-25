package models

type Playlist struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author User   `json:"author"`
}

func TestPlaylist() *Playlist {
	return &Playlist{
		ID:     1,
		Name:   "Lo-Fi",
		Author: *TestUser(),
	}
}
