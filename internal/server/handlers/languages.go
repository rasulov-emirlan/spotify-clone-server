package handlers

import (
	"net/http"
	"spotify-clone/server/internal/models"
	"spotify-clone/server/internal/store"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type languagesCreateRequest struct {
	Name string `json:"name"`
}

// LanguagesCreate docs
// @Tags		languages
// @Summary		Create a new language
// @Description	Creates a new language
// @Accept		json
// @Produce		json
// @Param       Authorization	header		string								true   "Bearer JWT Token"
// @Param       param			body		languagesCreateRequest		true   "A name for new genre"
// @Success		201 	"we created your language"
// @Router		/languages	[post]
func LanguagesCreate(store store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		claims := token.Claims.(jwt.MapClaims)
		check := false
		for _, v := range claims["roles"].([]interface{}) {
			if v == "admin" {
				check = true
				break
			}
		}
		if !check {
			throwError(http.StatusForbidden, "you have to be an admin to use this endpoint", nil, c)
			return nil
		}

		req := &languagesCreateRequest{}
		if err := c.Bind(req); err != nil {
			throwError(http.StatusBadRequest, "could not decode json", err, c)
			return err
		}

		language := &models.Language{
			Name: req.Name,
		}

		if err := store.Language().Create(language); err != nil {
			throwError(http.StatusBadRequest, "our database did not like your data", err, c)
			return err
		}

		respondWithData(http.StatusCreated, "we have created new language", nil, c)
		return nil
	}
}

// LanguagesListAll docs
// @Tags		languages
// @Summary		Get all languages
// @Description	Returns an array of all languages in our database
// @Accept		json
// @Produce		json
// @Success		200 	"we will give you an array of languages"
// @Router		/languages	[get]
func LanguagesListAll(store store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {

		languages, err := store.Language().ListAll()
		if err != nil {
			throwError(http.StatusInternalServerError, "", err, c)
			return err
		}
		return c.JSON(http.StatusOK, languages)
	}
}
