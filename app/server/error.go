package server

import (
	"net/http"
	"ushas/models"

	"github.com/labstack/echo/v4"
)

type errorResponse struct {
	Status int         `json:"status"`
	Why    string      `json:"why"`
	Origin interface{} `json:"request"`
}

func newErrResponse(status int, why string, origin interface{}) *errorResponse {
	return &errorResponse{
		Status: status,
		Why:    why,
		Origin: origin,
	}
}

// HTTPErrorHandler : Custom error handler.
func HTTPErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if !ok {
		// Errors not use *echo.HTTPError, such as panic.
		c.JSON(http.StatusInternalServerError, newErrResponse(http.StatusInternalServerError, "panic", nil))
		return
	}

	switch he.Message.(type) {
	case error:
		// If `error` type is returned in controller(handler).
		if e, ok := he.Message.(*models.APIError); ok {
			if e.Code > 0 {
				c.JSON(e.Code, newErrResponse(e.Code, e.Why, e.Origin))
			} else {
				c.JSON(he.Code, newErrResponse(he.Code, e.Why, e.Origin))
			}
			return
		}
		c.JSON(he.Code, newErrResponse(he.Code, he.Message.(error).Error(), nil))
	case string:
		// If the error happened other than controller(handler), such as URL not found.
		c.JSON(he.Code, newErrResponse(he.Code, he.Message.(string), nil))
	default:
		// Unreachable
		c.JSON(http.StatusInternalServerError, "Unknown error")
	}
}
