package handlers

import (
	"net/http"
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
// @Param		Authorization	header		string			true	"JWToken for auth"
// @Param       name            body        playlistCreateRequest          true     "The name of the playlist"
// @Success		200 	"we created your playlist"
// @Router		/playlists	[post]
func PlaylistsCreate(store store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &playlistCreateRequest{}
		if err := c.Bind(req); err != nil {
			throwError(http.StatusBadRequest, "could not decode json", err, c)
			return err
		}

		user := c.Get("user")
		token := user.(*jwt.Token)

		claims := token.Claims.(jwt.MapClaims)

		playlist := &models.Playlist{
			Name: req.Name,
			Author: models.User{
				ID: int(claims["userid"].(float64)),
			},
		}

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
// @Param       playlist_id            query        int          true     "The id for the playlist"
// @Param       song_id            query        int          true     "The id for the song"
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
				if err := store.Playlist().AddSong(playlistID, songID); err != nil {
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
