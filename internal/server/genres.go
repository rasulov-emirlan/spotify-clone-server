package server

import (
	"net/http"
	"spotify-clone/server/internal/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

type genresCreateRequest struct {
	Name string `json:"name"`
}

// handleGenresCreate docs
// @Tags		genres
// @Summary		Create a new genre
// @Description	Creates a new genre
// @Accept		json
// @Produce		json
// @Param       name              body       genresCreateRequest          true   "A name for new genre"
// @Success		200 	"we created your genre"
// @Router		/gnres	[post]
func (s *Server) handleGenresCreate() echo.HandlerFunc {
	type response struct {
		message string `json:"message"`
	}
	return func(c echo.Context) error {

		req := &genresCreateRequest{}

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

// handleGenresAddSong docs
// @Tags		genres
// @Summary		Add a song
// @Description	Adds a song to a genre
// @Accept		json
// @Produce		json
// @Param       genre              query       int          true   "id for a genre"
// @Param       song              query       int          true   "id for a song"
// @Success		200 	"we added a new song to the genre"
// @Router		/gnres	[put]
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

// handleGenresSongs docs
// @Tags		genres
// @Summary		Add a song
// @Description	Adds a song to a genre
// @Accept		json
// @Produce		json
// @Param       genre              path       int          true   "id for a genre"
// @Success		200 	"we added a new song to the genre"
// @Router		/gnres/{genre}	[get]
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
