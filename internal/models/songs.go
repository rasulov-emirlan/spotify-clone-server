package models

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/gofrs/uuid"
)

type Song struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author User   `json:"author"`
	URL    string `json:"url"`
}

func TestSong() *Song {
	return &Song{
		Title: "Dancing Machine",
		URL:   "youtube.com",
	}
}

func (s *Song) UUIDurl(path string, extension string) error {
	name, err := uuid.NewV1()
	if err != nil {
		return err
	}
	s.URL = path + name.String() + extension
	return nil
}

func (s *Song) UploadSong(audiofile *multipart.FileHeader) error {
	src, err := audiofile.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create("../" + s.URL)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

func (s *Song) DeleteSong() error {
	return os.Remove(s.URL)
}
