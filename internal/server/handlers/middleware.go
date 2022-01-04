package handlers

import (
	"net/http"
	"spotify-clone/server/internal/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func MiddlewareCheckRole(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user")
			token := user.(*jwt.Token)

			claims := token.Claims.(jwt.MapClaims)
			check := make([]bool, len(roles))

			for _, v := range claims["roles"].([]interface{}) {
				for j := 0; j < len(roles); j++ {
					if v == models.Admin {
						return next(c)
					}
					if v == roles[j] {
						check[j] = true
					}
				}
			}

			for i := 0; i < len(check); i++ {
				if !check[i] {
					throwError(http.StatusForbidden, "you are not authorized to use this endpoint", nil, c)
					return nil
				}
			}

			return next(c)
		}
	}
}
