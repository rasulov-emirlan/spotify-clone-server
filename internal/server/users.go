package server

import (
	"net/http"
	"spotify-clone/server/internal/models"

	"github.com/labstack/echo/v4"
)

func (s *Server) handleUserRegistration() echo.HandlerFunc {
	type request struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	type resposne struct {
		Message string `json:"messsage"`
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

		msg := resposne{
			Message: "we just registered you 😃",
		}
		return c.JSON(http.StatusOK, msg)
	}
}

func (s *Server) handleUserLogin() echo.HandlerFunc {
	type request struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	return func(c echo.Context) error {
		return nil
	}
}
