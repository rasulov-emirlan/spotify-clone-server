package server

import (
	"net/http"
	"spotify-clone/server/internal/models"
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

// handlePlaylistsCreate docs
// @Tags		playlists
// @Summary		Create a playlist
// @Description	Creates a new playlist that can be accesed by anyone but only you can edit it
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string			true	"JWToken for auth"
// @Param       name            body        playlistCreateRequest          true     "The name of the playlist"
// @Success		200 	"we created your playlist"
// @Router		/playlists	[post]
func (s *Server) handlePlaylistsCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &playlistCreateRequest{}
		if err := c.Bind(req); err != nil {
			s.Error(http.StatusBadRequest, "could not decode json", err, c)
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

		if err := s.store.Playlist().Create(playlist); err != nil {
			s.Error(http.StatusInternalServerError, "our database did not like your info", err, c)
			return err
		}
		return c.JSON(http.StatusOK, playlistCreateResponse{message: "we created your playlist! 🥳"})
	}
}

// handlePlaylistsAddSong docs
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
func (s *Server) handlePlaylistsAddSong() echo.HandlerFunc {
	type response struct {
		Message string `json:"message"`
	}
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		playlistID, err := strconv.Atoi(c.QueryParam("playlist"))
		if err != nil {
			s.Error(http.StatusBadRequest, "incorrect query parameters", err, c)
			return err
		}
		songID, err := strconv.Atoi(c.QueryParam("song"))

		if err != nil {
			s.Error(http.StatusBadRequest, "incorrect query parameters", err, c)
			return err
		}

		playlists, err := s.store.Playlist().UsersPlaylists(int(claims["userid"].(float64)))
		if err != nil {
			s.Error(http.StatusInternalServerError, "database did not find any records", err, c)
			return err
		}

		for i := 0; i < len(playlists); i++ {
			if playlists[i].ID == playlistID {
				if err := s.store.Playlist().AddSong(playlistID, songID); err != nil {
					s.Error(http.StatusInternalServerError, "database did not like your data", err, c)
					return err
				}
				return c.JSON(http.StatusOK, response{Message: "we added this song to your playlist! 😁"})
			}
		}

		s.Error(http.StatusBadRequest, "we dont know what went wrong", err, c)
		return err
	}
}

// handlePlaylistsGetSongsFromPlaylist docs
// @Tags		playlists
// @Summary		Get Songs from playlist
// @Description	Gives you an array of json with songs from a playlist you want
// @Accept		json
// @Produce		json
// @Param       id              path       int          true   "The id for the playlist"
// @Success		200 	"we created your playlist"
// @Router		/playlists/{id}	[get]
func (s *Server) handlePlaylistsGetSongsFromPlaylist() echo.HandlerFunc {
	type songs struct {
		ID         int    `json:"song_id"`
		Name       string `json:"name"`
		AuthorName string `json:"author_name"`
		URL        string `json:"url"`
		CoverURL   string `json:"cover_url"`
	}
	type response []songs
	return func(c echo.Context) error {
		playlistID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			s.Error(http.StatusBadRequest, "id is not a number!", err, c)
			return err
		}

		songs, err := s.store.Playlist().GetSongsFromPlaylist(playlistID)
		if err != nil {
			s.Error(http.StatusBadRequest, "database did not like your data!", err, c)
			return err
		}

		return c.JSON(http.StatusOK, songs)
	}
}
