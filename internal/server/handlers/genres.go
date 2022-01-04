package handlers

import (
	"log"
	"net/http"
	"spotify-clone/server/internal/fs"
	"spotify-clone/server/internal/fs/googlefs"
	"spotify-clone/server/internal/models"
	"spotify-clone/server/internal/store"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GenresCreate docs
// @Tags		genres
// @Summary		Add a localization
// @Description	Adds a new name for genre in a different language
// @Accept		mpfd
// @Produce		json
// @Param       Authorization	header		string		true   "Bearer JWT Token"
// @Param       name			formData		string		true   "A name for new genre"
// @Param       languageID		formData		int			true   "LanguageID to specify language of a name"
// @Param       cover			formData		file		true   "Cover image for the genre"
// @Success		201 	"we created your genre"
// @Router		/genres	[post]
func GenresCreate(store store.Store, fs fs.FileSystem) echo.HandlerFunc {
	return func(c echo.Context) error {
		cover, err := c.FormFile("cover")
		if err != nil {
			throwError(http.StatusBadRequest, "something is wrong with your file for cover", err, c)
			return err
		}
		coverfile, err := cover.Open()
		if err != nil {
			throwError(http.StatusBadRequest, "could not open file for cover", err, c)
			return err
		}

		name := c.FormValue("name")
		languageID, err := strconv.Atoi(c.FormValue("languageID"))
		if err != nil {
			throwError(http.StatusBadRequest, "could not convert languageID into integer value", err, c)
			return err
		}

		genre := &models.Genre{
			Name: name,
		}

		coverlink, err := fs.UploadFile(genre.Name, cover.Header["Content-Type"][0], coverfile, googlefs.FolderCovers)
		if err != nil {
			throwError(http.StatusInternalServerError, "could not upload your image into our cloud", err, c)
			return err
		}

		genre.CoverURL = coverlink

		if err := store.Genre().Create(genre, languageID); err != nil {
			throwError(http.StatusBadRequest, "our database did not like your data", err, c)
			return err
		}
		return c.JSON(http.StatusCreated, responseMessage{
			Code:    http.StatusCreated,
			Message: "we have created your genre",
			Data:    genre,
		})
	}
}

type genresAddLocalizationRequest struct {
	GenreID    int    `json:"genre_id"`
	Name       string `json:"name"`
	LanugageID int    `json:"lanugage_id"`
}

// GenresCreate docs
// @Tags		genres
// @Summary		Create a new genre
// @Description	Creates a new genre
// @Accept		json
// @Produce		json
// @Param       Authorization	header		string								true   "Bearer JWT Token"
// @Param       param			body		genresAddLocalizationRequest		true   "A name for new genre"
// @Success		201 	"we created your genre"
// @Router		/genres	[patch]
func GenresAddLocalization(store store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &genresAddLocalizationRequest{}

		if err := c.Bind(req); err != nil {
			throwError(http.StatusBadRequest, "Could not decode JSON", err, c)
			return err
		}

		if err := store.Genre().AddLocalization(req.GenreID, req.LanugageID, req.Name); err != nil {
			throwError(http.StatusInternalServerError, "Could not save data in our database", err, c)
			return err
		}

		return c.JSON(http.StatusOK, responseMessage{
			Code:    http.StatusOK,
			Message: "we have added new localization to your genre",
			Data:    nil,
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
