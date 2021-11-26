package server

import (
	"net/http"
	"spotify-clone/server/internal/models"

	"github.com/labstack/echo/v4"
)

type authRequest struct {
	Name     string `json:"name" example:"Johny Cash"`
	Email    string `json:"email" example:"john@gmai.com"`
	Password string `json:"password" example:"123456"`
}

type authResponse struct {
	Token string `json:"token"`
}

// @Summary      Register user
// @Description  Registers a new user and returns his token
// @Tags         auth
// @Accept	     json
// @Param		 param  body        authRequest     true "Authorization request"
// @Success		 200 	{object}	authResponse
// @Produce      json
// @Success      200  "json web token"
// @Router       /auth/register [post]
func (s *Server) handleUserRegistration() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &authRequest{}

		if err := c.Bind(req); err != nil {
			s.Error(http.StatusBadRequest, "could not decode json", err, c)
			return err
		}

		user := &models.User{
			Name:     req.Name,
			Password: req.Password,
			Email:    req.Email,
		}

		if err := user.BeforeCreate(); err != nil {
			s.Error(http.StatusBadRequest, "could not decode json 😥", err, c)
			return err
		}

		if err := s.store.User().Create(user); err != nil {
			s.Error(http.StatusBadRequest, "looks like our database did not like this user info 😥", err, c)
			return err
		}

		user, err := s.store.User().FindByEmail(req.Email)
		if err != nil {
			s.Error(http.StatusNonAuthoritativeInfo, "looks like you are not registered yet 😠", err, c)
			return err
		}

		if err := user.ComparePasswords(req.Password); err != nil {
			s.Error(http.StatusNonAuthoritativeInfo, "wrong password bozo! 😠", err, c)
			return err
		}

		token, err := user.GenerateJWT(s.jwtkey)
		if err != nil {
			s.Error(http.StatusNonAuthoritativeInfo, "could not generate a token 😭", err, c)
			return err
		}

		c.JSON(http.StatusOK, authResponse{Token: token})
		return nil
	}
}

// handleUserLogin godoc
// @Summary      Login user
// @Description  Returns a json web token if user is registered in database and enters correct data
// @Tags         auth
// @Accept	     json
// @Param		 param  body        authRequest     true "Authorization request"
// @Success		 200 	{object}	authResponse
// @Produce      json
// @Success      200  "json web token"
// @Router       /auth/login [post]
func (s *Server) handleUserLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &authRequest{}

		if err := c.Bind(req); err != nil {
			s.Error(http.StatusBadRequest, "could not parse json", err, c)
			return err
		}

		user, err := s.store.User().FindByEmail(req.Email)
		if err != nil {
			s.Error(http.StatusNonAuthoritativeInfo, "looks like you are not registered yet 😠", err, c)
			return err
		}

		if err := user.ComparePasswords(req.Password); err != nil {
			s.Error(http.StatusNonAuthoritativeInfo, "wrong password bozo! 😠", err, c)
			return err
		}

		token, err := user.GenerateJWT(s.jwtkey)
		if err != nil {
			s.Error(http.StatusNonAuthoritativeInfo, "could not generate a token 😭", err, c)
			return err
		}

		c.JSON(http.StatusOK, authResponse{Token: token})
		return nil
	}
}
