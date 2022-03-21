package models

import (
	"errors"

	"gorm.io/gorm"
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
// Error message looks good if `why` parameter doesn't contain period. ;)
func translateGormError(err error, why string, origin interface{}) *APIError {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		e := NewAPIError(err, why+"; Cannot find requested resource.", origin)
		e.Code = 404
		return e
	}
	if errors.Is(err, gorm.ErrInvalidTransaction) {
		return NewAPIError(err, why+"; Failed to establish transaction.", origin)
	}
	if errors.Is(err, gorm.ErrNotImplemented) {
		return NewAPIError(err, why+"; Method you called is not implemented yet.", origin)
	}
	if errors.Is(err, gorm.ErrMissingWhereClause) {
		return NewAPIError(err, why+"; Filtering parameters are missing.", origin)
	}
	return NewAPIError(err, why+"; Unknown error.", origin)
}
