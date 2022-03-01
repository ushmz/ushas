package controllers

import (
	"net/http"
	"ushas/models"

	"github.com/labstack/echo/v4"
)

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func newResponse(status int, message string, result interface{}) *response {
	return &response{status, message, result}
}

func new500Response(context echo.Context, request interface{}, err error) error {
	if e, ok := err.(*models.APIError); ok {
		return context.JSON(e.Code, newResponse(e.Code, e.Message, e.Result))
	}
	return context.JSON(http.StatusInternalServerError, newResponse(
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
		request,
	))
}

func newErrResponse(context echo.Context, status int, err error, request interface{}) error {
	if e, ok := err.(*models.APIError); ok {
		return context.JSON(e.Code, newResponse(e.Code, e.Message, e.Result))
	}
	return context.JSON(status, newResponse(status, http.StatusText(status), request))
}
