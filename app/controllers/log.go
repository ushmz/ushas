package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type LogController struct{}

func NewLogController() *LogController {
	return new(LogController)
}

func (lc *LogController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"OK",
	))
}
