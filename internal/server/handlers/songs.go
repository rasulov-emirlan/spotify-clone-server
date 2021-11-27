package handlers

import (
	"net/http"
	"spotify-clone/server/internal/models"
	"spotify-clone/server/internal/store"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type songRequest struct {
	Title string `json:""`
}

// SongsCreate docs
// @Tags		songs
// @Summary		Upload a song
// @Description	Uploads a song and its cover with all the info about that song
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string			true	"JWToken for auth"
// @Param		audio			formData	file			true    "The actual audiofile"
// @Param		cover			formData	file			true    "The cover for the song"
// @Param		title			formData	string			true    "The title for that song"
// @Success		200 	"we uploaded your song"
// @Router		/songs	[post]
func SongsCreate(store store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		claims := token.Claims.(jwt.MapClaims)
		title := c.FormValue("title")

		audio, err := c.FormFile("audio")
		if err != nil {
			throwError(http.StatusBadRequest, "unable to read audio file", err, c)
			return err
		}
		cover, err := c.FormFile("cover")
		if err != nil {
			throwError(http.StatusBadRequest, "unable to read image file", err, c)
			return err
		}
		song := models.Song{
			Title: title,
			Author: models.User{
				ID: int(claims["userid"].(float64)), // this is syntax for type assertion of interfaces
			},
		}

		song.UUIDurl()

		if err := store.Song().Create(&song); err != nil {
			throwError(http.StatusInternalServerError, "unable to save info into database", err, c)
			return err
		}

		if err := song.UploadSong(audio, cover); err != nil {
			throwError(http.StatusInternalServerError, "unable to save audio file into database", err, c)
			return err
		}
		c.JSON(http.StatusOK, "we uploaded your song")
		return nil
	}
}

func SongsFindByID(store store.Store) echo.HandlerFunc {
	type response struct {
		Title      string `json:"title"`
		AuthorName string `json:"author_name"`
		URL        string `json:"url"`
	}
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			throwError(http.StatusBadRequest, "it is not a proper id", err, c)
			return err
		}

		song, err := store.Song().FindByID(id)
		if err != nil {
			throwError(http.StatusBadRequest, "it is not a proper id", err, c)
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

// be carefull with this function
// if database is medium size this can slow down the whole server
func GetAllSongs(store store.Store) echo.HandlerFunc {
	type song struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		AuthorID int    `json:"author_id"`
		URL      string `json:""`
	}
	return func(c echo.Context) error {

		return nil
	}
}
