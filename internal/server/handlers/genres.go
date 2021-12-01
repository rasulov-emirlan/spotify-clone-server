package handlers

import (
	"log"
	"net/http"
	"spotify-clone/server/internal/models"
	"spotify-clone/server/internal/store"
	"strconv"

	"github.com/labstack/echo/v4"
)

type genresCreateRequest struct {
	Name string `json:"name"`
}

// GenresCreate docs
// @Tags		genres
// @Summary		Create a new genre
// @Description	Creates a new genre
// @Accept		json
// @Produce		json
// @Param       name              body       genresCreateRequest          true   "A name for new genre"
// @Success		200 	"we created your genre"
// @Router		/genres	[post]
func GenresCreate(store store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {

		req := &genresCreateRequest{}

		if err := c.Bind(req); err != nil {
			throwError(http.StatusBadRequest, "we could not decode you json", err, c)
			return err
		}

		genre := &models.Genre{
			Name: req.Name,
		}

		if err := store.Genre().Create(genre); err != nil {
			throwError(http.StatusBadRequest, "our database did not like your data", err, c)
			return err
		}
		return c.JSON(http.StatusOK, responseMessage{
			Code:    http.StatusOK,
			Message: "we have created your genre",
			Data:    genre,
		})
	}
}

// ListAllGenres docs
// @Tags		genres
// @Summary		List genres
// @Description	Lists all the genres in our database
// @Accept		json
// @Produce		json
// @Success		200 	body {object} []models.Genre
// @Router		/genres	[get]
func ListAllGenres(store store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		genres, err := store.Genre().ListAll()
		if err != nil {
			throwError(http.StatusInternalServerError, "database error", err, c)
			return err
		}
		log.Println(genres)
		return c.JSON(http.StatusOK, responseMessage{
			Code: http.StatusOK,
			Data: genres,
		})
	}
}

// GenresAddSong docs
// @Tags		genres
// @Summary		Add a song
// @Description	Adds a song to a genre
// @Accept		json
// @Produce		json
// @Param       genre              query       int          true   "id for a genre"
// @Param       song              query       int          true   "id for a song"
// @Success		200 	"we added a new song to the genre"
// @Router		/genres	[put]
func GenresAddSong(store store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		genreID, err := strconv.Atoi(c.QueryParam("genre"))
		if err != nil {
			throwError(http.StatusBadRequest, "could not get correct query parametrs", err, c)
			return err
		}

		songID, err := strconv.Atoi(c.QueryParam("song"))
		if err != nil {
			throwError(http.StatusBadRequest, "could not get correct query parametrs", err, c)
			return err
		}

		if err := store.Genre().AddSong(songID, genreID); err != nil {
			throwError(http.StatusBadRequest, "our database did not like your data", err, c)
			return err
		}
		return c.JSON(http.StatusOK, responseMessage{
			Code:    http.StatusOK,
			Message: "we have added your song to this genre! :)",
		})
	}
}

// GenresSongs docs
// @Tags		genres
// @Summary		Add a song
// @Description	Adds a song to a genre
// @Accept		json
// @Produce		json
// @Param       genre              path       int          true   "id for a genre"
// @Success		200 	"we added a new song to the genre"
// @Router		/genres/{genre}	[get]
func GenresSongs(store store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		genreID, err := strconv.Atoi(c.Param("genre"))
		if err != nil {
			throwError(http.StatusBadRequest, "could not get correct url parameter ", err, c)
			return err
		}

		songs, err := store.Genre().GetSongs(genreID)
		if err != nil {
			throwError(http.StatusBadRequest, "our database did not like your request ", err, c)
			return err
		}
		return c.JSON(http.StatusOK, responseMessage{
			Code: http.StatusOK,
			Data: songs,
		})
	}
}
