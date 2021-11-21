package models

type Playlist struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author User   `json:"author"`
}

type Playlits []Playlist
