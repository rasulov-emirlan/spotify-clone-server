package server

import (
	"net/http"
	"spotify-clone/server/internal/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *Server) handleGenresCreate() echo.HandlerFunc {
	type request struct {
		Name string `json:"name"`
	}
	type response struct {
		message string `json:"message"`
	}
	return func(c echo.Context) error {

		req := &request{}

		if err := c.Bind(req); err != nil {
			s.Error(http.StatusBadRequest, "we could not decode you json", err, c)
			return err
		}

		genre := &models.Genre{
			Name: req.Name,
		}

		if err := s.store.Genre().Create(genre); err != nil {
			s.Error(http.StatusBadRequest, "our database did not like your data", err, c)
			return err
		}
		return c.JSON(http.StatusOK, response{message: "we have added your genre"})
	}
}

func (s *Server) handleGenresAddSong() echo.HandlerFunc {
	type response struct {
		Message string `json:"message"`
	}
	return func(c echo.Context) error {
		genreID, err := strconv.Atoi(c.QueryParam("genre"))
		if err != nil {
			s.Error(http.StatusBadRequest, "could not get correct query parametrs", err, c)
			return err
		}

		songID, err := strconv.Atoi(c.QueryParam("song"))
		if err != nil {
			s.Error(http.StatusBadRequest, "could not get correct query parametrs", err, c)
			return err
		}

		if err := s.store.Genre().AddSong(songID, genreID); err != nil {
			s.Error(http.StatusBadRequest, "our database did not like your data", err, c)
			return err
		}
		return c.JSON(http.StatusOK, response{Message: "we have added your song to this genre! :)"})
	}
}

func (s *Server) handleGenresSongs() echo.HandlerFunc {
	return func(c echo.Context) error {
		genreID, err := strconv.Atoi(c.Param("genre"))
		if err != nil {
			s.Error(http.StatusBadRequest, "could not get correct url parameter ", err, c)
			return err
		}

		songs, err := s.store.Genre().GetSongs(genreID)
		if err != nil {
			s.Error(http.StatusBadRequest, "our database did not like your request ", err, c)
			return err
		}
		return c.JSON(http.StatusOK, songs)
	}
}