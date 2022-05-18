package models

import (
	"errors"

	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

var (
	// ErrNilReceiver : Throw when the receiver is nil
	ErrNilReceiver = xerrors.New("Receiver is nil")

	// ErrNoSuchData : Thorw when the requested data not found
	ErrNoSuchData = xerrors.New("No such data")
)

// APIError : Struct for internal error.
type APIError struct {
	Code   int         `json:"-"`
	Err    error       `json:"-"`
	Why    string      `json:"error"`
	Origin interface{} `json:"result"`
}

// Error : Returns error message.
func (e *APIError) Error() string {
	return e.Why
}

// NewAPIError : Returns new APIError
func NewAPIError(err error, why string, origin interface{}) *APIError {
	return &APIError{
		Err:    err,
		Why:    why,
		Origin: origin,
	}
}

// translateGormError : Translates gorm error to internal common error.
func translateGormError(err error, origin interface{}) *APIError {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		msg := "Cannot find requested resource."
		e := NewAPIError(xerrors.Errorf(msg+": %w", err), msg, origin)
		e.Code = 404
		return e
	}
	if errors.Is(err, gorm.ErrInvalidTransaction) {
		msg := "Failed to establish transaction."
		e := NewAPIError(xerrors.Errorf(msg+": %w", err), msg, origin)
		e.Code = 503
		return e
	}
	if errors.Is(err, gorm.ErrNotImplemented) {
		msg := "Method you called is not implemented yet."
		e := NewAPIError(xerrors.Errorf(msg+": %w", err), msg, origin)
		e.Code = 501
		return e
	}
	if errors.Is(err, gorm.ErrMissingWhereClause) {
		msg := "Filtering parameters are missing."
		e := NewAPIError(xerrors.Errorf(msg+": %w", err), msg, origin)
		e.Code = 400
		return e
	}
	msg := "Unknown error."
	return NewAPIError(xerrors.Errorf(msg+": %w", err), msg, origin)
}
