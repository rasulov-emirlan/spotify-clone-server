package server

import (
	"fmt"
	"net/http"
	"spotify-clone/server/internal/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

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

func (s *Server) IsAllowedMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := jwt.Parse(c.Request().Header.Get("Token"), func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Token could not be decoded")
			}
			return []byte(s.jwtkey), nil

		})
		if err != nil {
			s.Error(http.StatusBadRequest, "looks like you are not logged in", err, c)
			return err
		}
		c.Set("user", token.Claims)
		return next(c)
	}
}
