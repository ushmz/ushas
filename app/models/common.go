package models

import "net/http"

// CategoryCount : Distribution information for each categories.
type CategoryCount struct {
	// Category : Category name.
	Category string `json:"category"`
	// Count : Total number of pages.
	Count int `json:"count"`
}

// GroupCounts : Struct for group count
type GroupCounts struct {
	// GroupId : The ID assigned to the pair of "task IDs" and "condition ID"
	GroupId int `db:"group_id" json:"groupId"`

	// Count : Shows how many users are assigned to this group.
	Count int `db:"count" json:"count"`
}

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
