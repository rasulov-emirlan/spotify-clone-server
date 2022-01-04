package handlers

import (
	"net/http"
	"spotify-clone/server/internal/models"
	"spotify-clone/server/internal/store"
	"strconv"

	"github.com/labstack/echo/v4"
)

type countriesCreateRquest struct {
	Name string `json:"name"`
}

// CountriesCreate docs
// @Tags		countries
// @Summary		Create country
// @Description	Adds a new country to our database
// @Accept		json
// @Produce		json
// @Param		Authorization	header	string					true	"Bearer jwt"
// @Param		param			body	countriesCreateRquest	true	"request"
// @Success		201 	"201 if we added your country"
// @Router		/countries	[post]
func CountriesCreate(store store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &countriesCreateRquest{}

		if err := c.Bind(req); err != nil {
			throwError(http.StatusBadRequest, "could not decode json", err, c)
			return err
		}

		country := &models.Country{
			Name: req.Name,
		}

		if err := store.Country().Create(country); err != nil {
			throwError(http.StatusBadRequest, "our database did not like your data", err, c)
			return err
		}

		return c.JSON(http.StatusCreated, "we have added your language")
	}
}

// CountriesListAll docs
// @Tags		countries
// @Summary		List countries
// @Description	Returns all the countries in our database
// @Accept		json
// @Produce		json
// @Success		201		{object} []models.Country	"Success"
// @Router		/countries	[get]
func CountriesListAll(s store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		countries, err := s.Country().ListAll()
		if err != nil {
			throwError(http.StatusInternalServerError, "could not find find thme", err, c)
			return err
		}
		return c.JSON(http.StatusOK, countries)
	}
}

// CountriesDelete docs
// @Tags		countries
// @Summary		Delete country
// @Description	Deletes a country from our database
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string	true				"Bearer jwt"
// @Param		id				path		int		true				"country id"
// @Success		200				"we have deleted that country"
// @Router		/countries/{id}	[delete]
func CountriesDelete(s store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		countryID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			throwError(http.StatusBadRequest, "could not parse id for path param", err, c)
			return err
		}

		if err := s.Country().Delete(countryID); err != nil {
			throwError(http.StatusInternalServerError, "looks like our database did not your data", err, c)
			return err
		}

		respondWithData(http.StatusOK, "we have deleted that country", nil, c)
		return nil
	}
}
