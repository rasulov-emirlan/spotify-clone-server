package models

import (
	"errors"
	"io"
	"mime/multipart"
	"os"

	"github.com/gofrs/uuid"
)

type Song struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Author    User   `json:"author"`
	СreatedAt string `json:"created_at"`
	Length    int    `json:"length"`
	URL       string `json:"url"`
	CoverURL  string `json:"cover_url"`
	LikeCount int    `json:"like_count"`
	IsLiked   bool   `json:"is_liked"`
}

func TestSong() *Song {
	return &Song{
		Name: "Dancing Machine",
		URL:  "youtube.com",
	}
}

func (s *Song) UUIDurl() error {
	name, err := uuid.NewV1()
	if err != nil {
		return err
	}
	s.URL = "database/audio/" + name.String() + ".mp3"
	s.CoverURL = "database/covers/" + name.String() + ".png"
	return nil
}

var (
	errWrongFileType = errors.New("we do not support this type of file")
)

func (s *Song) UploadSong(audiofile, imagefile *multipart.FileHeader) error {

	if imagefile.Header["Content-Type"][0] != "image/png" {
		return errWrongFileType
	}
	if audiofile.Header["Content-Type"][0] != "audio/mpeg" && audiofile.Header["Content-Type"][0] != "audio/mp3" && audiofile.Header["Content-Type"][0] != "audio/wav" {
		return errWrongFileType
	}
	// here is logic for saving covers for songs
	coversrc, err := imagefile.Open()
	if err != nil {
		return err
	}
	defer coversrc.Close()

	coversdst, err := os.Create("../" + s.CoverURL)
	if err != nil {
		return err
	}
	defer coversdst.Close()

	if _, err := io.Copy(coversdst, coversrc); err != nil {
		return err
	}

	// here is logic for audio file itself
	audiosrc, err := audiofile.Open()
	if err != nil {
		return err
	}
	defer audiosrc.Close()

	audiodst, err := os.Create("../" + s.URL)
	if err != nil {
		return err
	}
	defer audiodst.Close()

	if _, err := io.Copy(audiodst, audiosrc); err != nil {
		return err
	}
	return nil
}

func (s *Song) DeleteSong() error {
	return os.Remove(s.URL)
}
