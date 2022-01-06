package handlers

import (
	"net/http"
	"spotify-clone/server/internal/store"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// UsersAddFavoriteSong docs
// @Tags		users
// @Summary		Add favorite song
// @Description	Adds a new favorite song for a certain user
// @Accept		json
// @Produce		json
// @Param		Authorization	header	string		true	"Bearer jwt"
// @Param		id				path	int			true	"Song id"
// @Success		201 	"201 if we added your country"
// @Router		/users/favorite/songs/{id}	[post]
func UsersAddFavoriteSong(s store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		claims := token.Claims.(jwt.MapClaims)

		songID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			throwError(http.StatusBadRequest, "could not turn songID into int", err, c)
			return err
		}

		if err := s.User().AddFavoriteSong(
			int(claims["userid"].(float64)),
			songID,
		); err != nil {
			throwError(http.StatusInternalServerError, "", err, c)
			return err
		}

		return c.JSON(http.StatusCreated, nil)
	}
}

// UsersListFavoriteSongs docs
// @Tags		users
// @Summary		Get Favorite Songs
// @Description	Returns an array of favorite songs
// @Accept		json
// @Produce		json
// @Param		user	query		int			true	"User id"
// @Param		limit	query		int			true	"how many songs you want"
// @Param		offset	query		int			true	"from which index to start"
// @Success		200 	{object}	[]models.Song	"array of favorite songs"
// @Router		/users/favorite/songs/	[get]
func UsersListFavoriteSongs(s store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.Atoi(c.QueryParam("user"))
		if err != nil {
			throwError(http.StatusBadRequest, "", err, c)
			return err
		}
		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil {
			throwError(http.StatusBadRequest, "", err, c)
			return err
		}
		offset, err := strconv.Atoi(c.QueryParam("offset"))
		if err != nil {
			throwError(http.StatusBadRequest, "", err, c)
			return err
		}

		songs, err := s.User().ListFavoriteSongs(userID, limit, offset)
		if err != nil {
			throwError(http.StatusInternalServerError, "", err, c)
			return err
		}

		return c.JSON(http.StatusOK, songs)
	}
}

// UsersRemoveFavoriteSong docs
// @Tags		users
// @Summary		Remove Favorite Song
// @Description	Removes a song from favorites
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string			true	"Bearer jwt"
// @Param		id				path		int				true	"song id"
// @Success		200 			"200 if deleted successfuly"
// @Router		/users/favorite/songs/{id}	[delete]
func UsersRemoveFavoriteSong(s store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		claims := token.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		if userID == 0 {
			throwError(http.StatusBadRequest, "", nil, c)
			return nil
		}

		songID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			throwError(http.StatusBadRequest, "", err, c)
			return err
		}

		if err := s.User().RemoveFromFavoriteSongs(userID, songID); err != nil {
			throwError(http.StatusInternalServerError, "", err, c)
			return err
		}

		respondWithData(http.StatusOK, "we have deleted the song from your favorites", nil, c)
		return nil
	}
}

// UsersAddFavoriteAuthor docs
// @Tags		users
// @Summary		Add Favorite Author
// @Description	Adds a favorite author
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string			true	"Bearer jwt"
// @Param		id				path		int				true	"author id"
// @Success		201 			"200 if added successfuly"
// @Router		/users/favorite/authors/{id}	[post]
func UsersAddFavoriteAuthor(s store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		claims := token.Claims.(jwt.MapClaims)

		userID, ok := claims["userid"].(float64)
		if !ok {
			throwError(http.StatusBadRequest, "could not turn userID into int", nil, c)
			return nil
		}

		authorID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			throwError(http.StatusBadRequest, "could not turn authorID into int", err, c)
			return err
		}

		if err := s.User().AddFavoriteAuthor(int(userID), authorID); err != nil {
			throwError(http.StatusInternalServerError, "looks like our database did not like your data", err, c)
			return err
		}
		return c.JSON(http.StatusCreated, nil)
	}
}

// UsersListFavoriteAuthors docs
// @Tags		users
// @Summary		List Favorite Authors
// @Description	Returns an array of authors from `offset` with length of `limit`
// @Accept		json
// @Produce		json
// @Param		id				query		int				true	"user id"
// @Param		limit			query		int				true	"limit"
// @Param		offset			query		int				true	"offset"
// @Success		200 			{object}	[]models.User	"all the favorite authors"
// @Router		/users/favorite/authors/	[get]
func UsersListFavoriteAuthors(s store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.Atoi(c.QueryParam("id"))
		if err != nil {
			throwError(http.StatusBadRequest, "could not turn userID into int", err, c)
			return err
		}

		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil {
			throwError(http.StatusBadRequest, "could not turn limit into int", err, c)
			return err
		}

		offset, err := strconv.Atoi(c.QueryParam("offset"))
		if err != nil {
			throwError(http.StatusBadRequest, "could not turn offset into int", err, c)
			return err
		}

		authors, err := s.User().ListFavoriteAuthors(userID, limit, offset)
		if err != nil {
			throwError(http.StatusInternalServerError, "looks like our database did not like your data", err, c)
			return err
		}

		return c.JSON(http.StatusOK, authors)
	}
}

// UsersListFavoriteAuthors docs
// @Tags		users
// @Summary		Delete Favorite Author
// @Description	Removes an authors from favorites
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string			true	"Bearer jwt"
// @Param		id				path		int				true	"author id"
// @Success		200 			{object}	[]models.User	"all the favorite authors"
// @Router		/users/favorite/authors/{id}	[delete]
func UsersRemoveFavoriteAuthor(s store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		claims := token.Claims.(jwt.MapClaims)

		userID, ok := claims["userid"].(float64)
		if !ok {
			throwError(http.StatusBadRequest, "could not turn userID into int", nil, c)
			return nil
		}

		authorID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			throwError(http.StatusBadRequest, "could not turn authorID into int", err, c)
			return err
		}

		if err := s.User().RemoveFromFavoriteAuthors(int(userID), authorID); err != nil {
			throwError(http.StatusInternalServerError, "looks like our database did not your data", err, c)
			return err
		}
		return c.JSON(http.StatusOK, nil)
	}
}
