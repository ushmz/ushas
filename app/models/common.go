package models

import "net/http"

// APIError : Struct for internal error.
type APIError struct {
	Err     error       `json:"-"`
	Message string      `json:"error"`
	Code    int         `json:"status"`
	Result  interface{} `json:"result"`
}

// Error : Returns error message.
func (e *APIError) Error() string {
	return e.Message
}

// RaiseNotFoundError : Returns new error of no resource found.
func RaiseNotFoundError(err error, msg string) error {
	return &APIError{
		Err:     err,
		Message: msg,
		Code:    http.StatusNotFound,
	}
}

// RaiseBadRequestError : Returns new error that means it failed to parse request body.
func RaiseBadRequestError(err error, msg string, result ...interface{}) error {
	return &APIError{
		Err:     err,
		Message: msg,
		Code:    http.StatusBadRequest,
		Result:  result,
	}
}

// RaiseInternalServerError : Returns new error of internal server error.
func RaiseInternalServerError(err error, msg string, result ...interface{}) error {
	return &APIError{
		Err:     err,
		Message: msg,
		Code:    http.StatusInternalServerError,
		Result:  result,
	}
}
