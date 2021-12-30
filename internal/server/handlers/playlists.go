package handlers

import (
	"log"
	"net/http"
	"spotify-clone/server/internal/fs"
	"spotify-clone/server/internal/models"
	"spotify-clone/server/internal/store"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type playlistCreateRequest struct {
	Name string `json:"name" example:"my favorites"`
}
type playlistCreateResponse struct {
	message string `json:"message"`
}

// PlaylistsCreate docs
// @Tags		playlists
// @Summary		Create a playlist
// @Description	Creates a new playlist that can be accesed by anyone but only you can edit it
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string		true	"JWToken for auth"
// @Param       name            formData	string		true	"The name of the playlist"
// @Param       cover			formData	file		true	"The name of the playlist"
// @Success		200 	"we created your playlist"
// @Router		/playlists	[post]
func PlaylistsCreate(store store.Store, fs fs.FileSystem) echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user")
		token := user.(*jwt.Token)

		claims := token.Claims.(jwt.MapClaims)

		cover, err := c.FormFile("cover")
		if err != nil {
			throwError(http.StatusBadRequest, "could not decode image for cover", err, c)
			return err
		}
		coverfile, err := cover.Open()
		if err != nil {
			throwError(http.StatusBadRequest, "could not open image for cover", err, c)
			return err
		}

		name := c.FormValue("name")

		playlist := &models.Playlist{
			Name: name,
			Author: models.User{
				ID: int(claims["userid"].(float64)),
			},
		}

		// 😅 this esoteric looking string is an id for folder in  google drive💿
		coverlink, err := fs.UploadFile(playlist.Name, cover.Header["Content-Type"][0], coverfile, "1VAu4UO77e9OeXCfckOKT2mEGttoFgmeq")
		if err != nil {
			throwError(http.StatusInternalServerError, "could not upload cover to our server", err, c)
			return err
		}


		playlist.CoverUrl = coverlink

		if err := store.Playlist().Create(playlist); err != nil {
			throwError(http.StatusInternalServerError, "our database did not like your info", err, c)
			return err
		}
		return c.JSON(http.StatusOK, responseMessage{
			Code:    http.StatusOK,
			Message: "we created your playlist! 🥳"})
	}
}

// PlaylistsAddSong docs
// @Tags		playlists
// @Summary		Add a song
// @Description	Adds a song to whatever playlist you want to. But it has to be your playlist that you created
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string			true	"JWToken for auth"
// @Param       playlist		query		int          true     "The id for the playlist"
// @Param       song			query		int          true     "The id for the song"
// @Success		200 	"we created your playlist"
// @Router		/playlists	[put]
func PlaylistsAddSong(store store.Store) echo.HandlerFunc {
	type response struct {
		Message string `json:"message"`
	}
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		playlistID, err := strconv.Atoi(c.QueryParam("playlist"))
		if err != nil {
			throwError(http.StatusBadRequest, "incorrect query parameters", err, c)
			return err
		}
		songID, err := strconv.Atoi(c.QueryParam("song"))

		if err != nil {
			throwError(http.StatusBadRequest, "incorrect query parameters", err, c)
			return err
		}

		playlists, err := store.Playlist().UsersPlaylists(int(claims["userid"].(float64)))
		if err != nil {
			throwError(http.StatusInternalServerError, "database did not find any records", err, c)
			return err
		}

		for i := 0; i < len(playlists); i++ {
			if playlists[i].ID == playlistID {
				if err := store.Playlist().AddSong(songID, playlistID); err != nil {
					throwError(http.StatusInternalServerError, "database did not like your data", err, c)
					return err
				}
				return c.JSON(http.StatusOK, response{Message: "we added this song to your playlist! 😁"})
			}
		}

		throwError(http.StatusBadRequest, "we dont know what went wrong", err, c)
		return err
	}
}

// ListAllPlaylists docs
// @Tags		playlists
// @Summary		List playlists
// @Description	Lists all the playlists in our database
// @Accept		json
// @Produce		json
// @Success		200 	body {object} []models.Playlist
// @Router		/playlists	[get]
func ListAllPlaylists(store store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		playlists, err := store.Playlist().ListAll()
		if err != nil {
			throwError(http.StatusInternalServerError, "database error", err, c)
			return err
		}
		log.Println(playlists)
		return c.JSON(http.StatusOK, responseMessage{
			Code: http.StatusOK,
			Data: playlists,
		})
		return nil
	}
}

// GetSongsFromPlaylist docs
// @Tags		playlists
// @Summary		Get Songs from playlist
// @Description	Gives you an array of json with songs from a playlist you want
// @Accept		json
// @Produce		json
// @Param       id              path       int          true   "The id for the playlist"
// @Success		200 	"we created your playlist"
// @Router		/playlists/{id}	[get]
func GetSongsFromPlaylist(store store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		playlistID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			throwError(http.StatusBadRequest, "id is not a number!", err, c)
			return err
		}

		songs, err := store.Playlist().GetSongsFromPlaylist(playlistID)
		if err != nil {
			throwError(http.StatusBadRequest, "database did not like your data!", err, c)
			return err
		}

		return c.JSON(http.StatusOK,
			responseMessage{
				Code: http.StatusOK,
				Data: songs,
			})
	}
}
