package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskController struct{}

func NewTaskController() *TaskController {
	return new(TaskController)
}

func (tc *TaskController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"OK",
	))
}
