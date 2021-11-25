package server

import (
	"net/http"
	"spotify-clone/server/internal/models"

	"github.com/labstack/echo/v4"
)

// handleUserRegistration godoc
// @Summary      Register user
// @Description  Registers a new user and returns his token
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  "json web token"
// @Router       /users/{id} [get]
func (s *Server) handleUserRegistration() echo.HandlerFunc {
	type request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		Token string `json:"token"`
	}

	return func(c echo.Context) error {
		req := &request{}

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

		c.JSON(http.StatusOK, response{Token: token})
		return nil
	}
}

func (s *Server) handleUserLogin() echo.HandlerFunc {
	type request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type response struct {
		Token string `json:"token"`
	}
	return func(c echo.Context) error {
		req := &request{}

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

		c.JSON(http.StatusOK, response{Token: token})
		return nil
	}
}
