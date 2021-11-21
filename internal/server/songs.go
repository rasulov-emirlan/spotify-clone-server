package server

import (
	"net/http"
	"spotify-clone/server/internal/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *Server) handleSongsCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.FormValue("title")
		authorId, err := strconv.Atoi(c.FormValue("author_id"))

		if err != nil {
			s.Error(http.StatusBadRequest, "unable to read author_id", err, c)
			return err
		}

		audio, err := c.FormFile("audio")
		if err != nil {
			s.Error(http.StatusBadRequest, "unable to read audio file", err, c)
			return err
		}
		song := &models.Song{
			Title: title,
			Author: models.User{
				ID: authorId,
			},
			URL: "database/audio/" + title + ".mp3",
		}

		if err := s.store.Song().Create(song); err != nil {
			s.Error(http.StatusInternalServerError, "unable to save info into database", err, c)
			return err
		}

		if err := song.UploadSong(audio); err != nil {
			s.Error(http.StatusInternalServerError, "unable to save audio file into database", err, c)
			return err
		}
		c.JSON(http.StatusOK, "we uploaded your song")
		return nil
	}
}

func (s *Server) handleSongsFindByID() echo.HandlerFunc {
	type response struct {
		Title      string `json:"title"`
		AuthorName string `json:"author_name"`
		URL        string `json:"url"`
	}
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			s.Error(http.StatusBadRequest, "it is not a proper id", err, c)
			return err
		}

		song, err := s.store.Song().FindByID(id)
		if err != nil {
			s.Error(http.StatusBadRequest, "it is not a proper id", err, c)
			return err
		}
		resp := response{
			Title: song.Title,
			URL:   song.URL,
		}
		c.JSON(http.StatusOK, resp)
		return nil
	}
}
