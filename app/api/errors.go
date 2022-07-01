package api

import (
	"net/http"
	"ushas/models"

	"github.com/labstack/echo/v4"
)

type errorResponse struct {
	Why    string      `json:"why"`
	Origin interface{} `json:"request"`
}

func newErrResponse(status int, why string, origin interface{}) *errorResponse {
	return &errorResponse{
		Why:    why,
		Origin: origin,
	}
}

// httpErrorHandler : Custom error handler.
func httpErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if !ok {
		// Errors not use *echo.HTTPError, such as panic.
		c.JSON(
			http.StatusInternalServerError,
			newErrResponse(http.StatusInternalServerError, "panic", nil),
		)
		return
	}

	switch he.Message.(type) {
	case error:
		if e, ok := he.Message.(*models.AppError); ok {
			c.Logger().Errorf("%+v", e.Err)
			c.JSON(e.Code, newErrResponse(e.Code, e.Why, e.Origin))
			return
		}
		// If `error` type is returned in controller(handler).
		e := he.Message.(error)
		c.Logger().Errorf("%+v", e)
		c.JSON(he.Code, newErrResponse(he.Code, e.Error(), ""))
	case string:
		// If the error happened other than controller(handler), such as URL not found.
		c.JSON(he.Code, newErrResponse(he.Code, he.Message.(string), nil))
	default:
		// Unreachable
		c.JSON(http.StatusInternalServerError, "Unknown error")
	}
}
