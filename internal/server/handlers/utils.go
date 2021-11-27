package handlers

import (
	"log"

	"github.com/labstack/echo/v4"
)

type responseMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func throwError(code int, msg string, data interface{}, c echo.Context) {
	c.JSON(code, responseMessage{
		Code:    code,
		Message: msg,
		Data:    data,
	})
	log.Println(data)
}