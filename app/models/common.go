package models

import (
	"errors"

	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

// InternalError : Struct for internal error.
type InternalError struct {
	Code   int         `json:"-"`
	Err    error       `json:"-"`
	Why    string      `json:"error"`
	Origin interface{} `json:"result"`
}

// Error : Returns error message.
func (e *InternalError) Error() string {
	return e.Why
}

// NewInternalError : Returns new APIError
func NewInternalError(err error, why string, origin interface{}) *InternalError {
	return &InternalError{
		Err:    err,
		Why:    why,
		Origin: origin,
	}
}

// translateGormError : Translates gorm error to internal common error.
func translateGormError(err error, origin interface{}) *InternalError {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		msg := "Cannot find requested resource."
		e := NewInternalError(xerrors.Errorf(msg+": %w", err), msg, origin)
		e.Code = 404
		return e
	}
	if errors.Is(err, gorm.ErrInvalidTransaction) {
		msg := "Failed to establish transaction."
		e := NewInternalError(xerrors.Errorf(msg+": %w", err), msg, origin)
		e.Code = 503
		return e
	}
	if errors.Is(err, gorm.ErrNotImplemented) {
		msg := "Method you called is not implemented yet."
		e := NewInternalError(xerrors.Errorf(msg+": %w", err), msg, origin)
		e.Code = 501
		return e
	}
	if errors.Is(err, gorm.ErrMissingWhereClause) {
		msg := "Filtering parameters are missing."
		e := NewInternalError(xerrors.Errorf(msg+": %w", err), msg, origin)
		e.Code = 400
		return e
	}
	msg := "Unknown error."
	return NewInternalError(xerrors.Errorf(msg+": %w", err), msg, origin)
}
