package server

import (
	"net/http"
	"spotify-clone/server/internal/models"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (s *Server) handlePlaylistsCreate() echo.HandlerFunc {
	type request struct {
		Name string `json:"name"`
	}
	type response struct {
		message string `json:"message"`
	}
	return func(c echo.Context) error {
		req := &request{}
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
		return c.JSON(http.StatusOK, response{message: "we created your playlist! 🥳"})
	}
}

func (s *Server) handlePlaylistsAddSong() echo.HandlerFunc {
	type response struct {
		Message string `json:"message"`
	}
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		playlistID, err := strconv.Atoi(c.QueryParam("playlist_id"))
		if err != nil {
			s.Error(http.StatusBadRequest, "incorrect query parameters", err, c)
			return err
		}
		songID, err := strconv.Atoi(c.QueryParam("song_id"))

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
		playlistID, err := strconv.Atoi(c.Param("playlist_id"))
		if err != nil {
			s.Error(http.StatusBadRequest, "playlist_id is not a number!", err, c)
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
