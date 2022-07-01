package api

import (
	"fmt"
	"net/http"
	"strings"
	"ushas/models"

	"github.com/go-playground/validator"
)

// Validator : Struct for validation.
type Validator struct {
	validator *validator.Validate
}

// Validate : Validate request body.
func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		results := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			results = append(results, translateValidationError(e))
		}

		return models.NewAppError(
			err,
			http.StatusBadRequest,
			strings.Join(results, ";"),
			i,
		)
	}
	return nil
}

func translateValidationError(e validator.FieldError) string {
	f := e.Field()
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", f)
	case "max":
		return fmt.Sprintf("%s must be lower than %s", f, e.Param())
	case "min":
		return fmt.Sprintf("%s must be greater then %s", f, e.Param())
	}
	return fmt.Sprintf("%s is not valid %s", e.Field(), e.Value())
}
