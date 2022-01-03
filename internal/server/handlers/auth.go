package handlers

import (
	"net/http"
	"spotify-clone/server/internal/models"
	"spotify-clone/server/internal/store"

	"github.com/labstack/echo/v4"
)

type authRequest struct {
	UserName  string `json:"username" example:"Johnny"`
	FullName  string `json:"full_name" example:"Johny Cash"`
	BirthDate string `json:"birth_date" example:"2000-01-01"`
	Email     string `json:"email" example:"john@gmai.com"`
	Password  string `json:"password" example:"123456" omniemit`
}

type authResponse struct {
	Token string `json:"token"`
}

// @Summary      Register user
// @Description  Registers a new user and returns his token
// @Tags         auth
// @Accept	     json
// @Param		 param  body        authRequest     true "Authorization request"
// @Success		 201	{object}	authResponse
// @Produce      json
// @Router       /auth/register [post]
func UserRegistration(store store.Store, JWTkey string) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &authRequest{}

		if err := c.Bind(req); err != nil {
			throwError(http.StatusBadRequest, "could not decode json", err, c)
			return err
		}

		user := &models.User{
			UserName:  req.UserName,
			FullName:  req.FullName,
			BirthDate: req.BirthDate,
			Password:  req.Password,
			Email:     req.Email,
		}

		if err := user.BeforeCreate(); err != nil {
			throwError(http.StatusBadRequest, "could not decode json 😥", err, c)
			return err
		}

		if err := store.User().Create(user); err != nil {
			throwError(http.StatusBadRequest, "looks like our database did not like this user info 😥", err, c)
			return err
		}

		user, err := store.User().FindByEmail(req.Email)
		if err != nil {
			throwError(http.StatusNonAuthoritativeInfo, "looks like you are not registered yet 😠", err, c)
			return err
		}

		if err := user.ComparePasswords(req.Password); err != nil {
			throwError(http.StatusNonAuthoritativeInfo, "wrong password bozo! 😠", err, c)
			return err
		}

		token, err := user.GenerateJWT(JWTkey)
		if err != nil {
			throwError(http.StatusNonAuthoritativeInfo, "could not generate a token 😭", err, c)
			return err
		}

		c.JSON(http.StatusCreated, authResponse{Token: token})
		return nil
	}
}

// UserLogin godoc
// @Summary      Login user
// @Description  Returns a json web token if user is registered in database and enters correct data
// @Tags         auth
// @Accept	     json
// @Param		 param  body        authRequest     true "Authorization request"
// @Produce      json
// @Success		 201	{object}	authResponse
// @Router       /auth/login [post]
func UserLogin(store store.Store, JWTkey string) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &authRequest{}

		if err := c.Bind(req); err != nil {
			throwError(http.StatusBadRequest, "could not parse json", err, c)
			return err
		}

		user, err := store.User().FindByEmail(req.Email)
		if err != nil {
			throwError(http.StatusNonAuthoritativeInfo, "looks like you are not registered yet 😠", err, c)
			return err
		}

		if err := user.ComparePasswords(req.Password); err != nil {
			throwError(http.StatusNonAuthoritativeInfo, "wrong password bozo! 😠", err, c)
			return err
		}

		token, err := user.GenerateJWT(JWTkey)
		if err != nil {
			throwError(http.StatusNonAuthoritativeInfo, "could not generate a token 😭", err, c)
			return err
		}

		c.JSON(http.StatusAccepted, authResponse{Token: token})
		return nil
	}
}
