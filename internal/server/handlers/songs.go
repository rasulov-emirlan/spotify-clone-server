package handlers

import (
	"errors"
	"net/http"
	"spotify-clone/server/internal/fs"
	"spotify-clone/server/internal/fs/googlefs"
	"spotify-clone/server/internal/models"
	"spotify-clone/server/internal/store"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type songRequest struct {
	Title string `json:""`
}

var (
	errNotASinger = errors.New("Not a verified singer")
)

// SongsCreate docs
// @Tags		songs
// @Summary		Upload a song
// @Description	Uploads a song and its cover with all the info about that song
// @Accept		mpfd
// @Produce		json
// @Param		Authorization	header		string			true	"JWToken for auth"
// @Param		audio			formData	file			true    "The actual audiofile"
// @Param		cover			formData	file			true    "The cover for the song"
// @Param		name			formData	string			true    "The name for that song"
// @Success		200 	"we uploaded your song"
// @Fail		400		"not a verified singer"
// @Router		/songs	[post]
func SongsCreate(store store.Store, fs fs.FileSystem) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		claims := token.Claims.(jwt.MapClaims)
		// ifSinger := false
		// for _, role := range claims["roles"].([]interface{}) {
		// 	fmt.Println(role)
		// 	if role == "singer" {
		// 		ifSinger = true
		// 		break
		// 	}
		// }
		// if !ifSinger {
		// 	throwError(http.StatusBadRequest, "you are not a verified singer, thus cannot upload your songs", errNotASinger, c)
		// 	return errNotASinger
		// }

		name := c.FormValue("name")

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
			Name: name,
			Author: models.User{
				ID: int(claims["userid"].(float64)), // this is syntax for type assertion of interfaces
			},
		}

		a, err := audio.Open()
		if err != nil {
			throwError(http.StatusInternalServerError, "unable to decode audiofile", err, c)
			return err
		}
		defer a.Close()
		b, err := cover.Open()
		if err != nil {
			throwError(http.StatusInternalServerError, "unable to decode imagefile", err, c)
			return err
		}
		defer a.Close()

		songid, err := fs.UploadFile(song.Name, audio.Header["Content-Type"][0], a, googlefs.FolderSongs)
		if err != nil {
			throwError(http.StatusInternalServerError, "unable to save audiofile to server", err, c)
			return err
		}

		coverid, err := fs.UploadFile(song.Name, cover.Header["Content-Type"][0], b, googlefs.FolderCovers)
		if err != nil {
			throwError(http.StatusInternalServerError, "unable to save audiofile to server", err, c)
			return err
		}

		song.URL = songid
		song.CoverURL = coverid

		if err := store.Song().Create(&song); err != nil {
			throwError(http.StatusInternalServerError, "unable to save info into database", err, c)
			return err
		}

		respondWithData(http.StatusAccepted, "we uploaded your song", jsonMapper{
			"songID": songid,
		}, c)
		return nil
	}
}

// SongsFromAtoB docs
// @Tags		songs
// @Summary		Get songs
// @Description	Returns songs from some id to some id
// @Accept		json
// @Produce		json
// @Param		from			query	int			true    "from which id to start"
// @Param		to				query	int			true    "at which id to end"
// @Success		200				{object}	[]models.Song	"list of songs"
// @Router		/songs	[get]
func SongsFromAtoB(store store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		from, err := strconv.Atoi(c.QueryParam("from"))
		if err != nil {
			throwError(http.StatusBadRequest, "did not like path params", err, c)
			return err
		}
		to, err := strconv.Atoi(c.QueryParam("to"))
		if err != nil {
			throwError(http.StatusBadRequest, "did not like path params", err, c)
			return err
		}

		songs, err := store.Song().GetSongs(from, to)
		if err != nil {
			throwError(http.StatusBadRequest, "did not like path params", err, c)
			return err
		}
		return c.JSON(http.StatusOK, responseMessage{
			Code: http.StatusOK,
			Data: songs,
		})
	}
}

func SongsFindByID(store store.Store) echo.HandlerFunc {
	type response struct {
		Name       string `json:"name"`
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
			Name: song.Name,
			URL:  song.URL,
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
