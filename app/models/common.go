package models

import "net/http"

type APIError struct {
	Err     error       `json:"-"`
	Message string      `json:"error"`
	Code    int         `json:"status"`
	Result  interface{} `json:"result"`
}

func (e *APIError) Error() string {
	return e.Message
}

func RaiseNotFoundError(err error, msg string) error {
	return &APIError{
		Err:     err,
		Message: msg,
		Code:    http.StatusNotFound,
	}
}

func RaiseInternalServerError(err error, msg string, result ...interface{}) error {
	return &APIError{
		Err:     err,
		Message: msg,
		Code:    http.StatusInternalServerError,
		Result:  result,
	}
}

func RaiseBadRequestError(err error, msg string, result ...interface{}) error {
	return &APIError{
		Err:     err,
		Message: msg,
		Code:    http.StatusBadRequest,
		Result:  result,
	}
}
