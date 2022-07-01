package models

import (
	"errors"

	"gorm.io/gorm"
)

var (
	// ErrDBOperationFailed : DB operation failed
	ErrDBOperationFailed = errors.New("DB operation failed")

	// ErrDBConnectionFailed : DB connection fauled
	ErrDBConnectionFailed = errors.New("DB connection fauled")

	// ErrNilReceiver : Throw when the receiver is nil
	ErrNilReceiver = errors.New("Receiver is nil")

	// ErrBadRequest : HTTP request body or argument is invalid
	ErrBadRequest = errors.New("Invalid request")

	// ErrNoSuchData : Thorw when the requested data not found
	ErrNoSuchData = errors.New("No such data")

	// ErrInternal : Internal errors that don't have to tell users in detail
	ErrInternal = errors.New("Internal server error")
)

// AppError : Struct for internal error.
type AppError struct {
	Code   int         `json:"-"`
	Err    error       `json:"-"`
	Why    string      `json:"error"`
	Origin interface{} `json:"request"`
}

// Error : Returns error message.
func (e *AppError) Error() string {
	return e.Why
}

// NewAppError : Returns new APIError
func NewAppError(err error, code int, why string, origin interface{}) *AppError {
	return &AppError{
		Code:   code,
		Err:    err,
		Why:    why,
		Origin: origin,
	}
}

// translateGormError : Translates gorm error to internal common error.
func translateGormError(err error, origin interface{}) *AppError {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		msg := "Cannot find requested resource."
		e := NewAppError(ErrNoSuchData, 404, msg, origin)
		return e
	}
	if errors.Is(err, gorm.ErrInvalidTransaction) {
		msg := "Failed to establish transaction."
		e := NewAppError(ErrInternal, 503, msg, origin)
		return e
	}
	if errors.Is(err, gorm.ErrNotImplemented) {
		msg := "Method you called is not implemented yet."
		e := NewAppError(ErrInternal, 501, msg, origin)
		return e
	}
	if errors.Is(err, gorm.ErrMissingWhereClause) {
		msg := "Filtering parameters are missing."
		e := NewAppError(ErrInternal, 400, msg, origin)
		return e
	}
	msg := "Internal error."
	e := NewAppError(ErrInternal, 503, msg, origin)
	return e
}
