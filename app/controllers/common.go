package controllers

import (
	"net/http"
	"ushas/models"

	"github.com/labstack/echo/v4"
)

// response : Struct for API response.
type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// newResponse : Return new response.
func newResponse(status int, message string, result interface{}) *response {
	return &response{status, message, result}
}

// new500Response : Return new 500 error response.
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

// newErrResponse : Return new error response.
func newErrResponse(context echo.Context, status int, err error, request interface{}) error {
	if e, ok := err.(*models.APIError); ok {
		return context.JSON(e.Code, newResponse(e.Code, e.Message, e.Result))
	}
	return context.JSON(status, newResponse(status, http.StatusText(status), request))
}
