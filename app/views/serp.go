package views

import (
	"ushas/models"
)

// SERPWithIcon : The list of this type struct will be returned as a response of `serp` endpoint.
type SERPWithIcon struct {
	models.SearchPage

	// Linked : Users' behavioral data that probably leaked to third party.
	// For more detail, see `Linked` type.
	Linked []models.LinkedPage `json:"linked"`
}

// NewSERPWithIconView :
func NewSERPWithIconView(pages []models.SearchPage, linkedPages []models.LinkedPage) []SERPWithIcon {
	// serp := []SERPWithIcon{}
	//
	// for _, v := range pages {
	//
	// }
	return nil
}

// SERPWithRatio : The list of this type struct will be returned as a response of `serp` endpoint.
type SERPWithRatio struct {
	models.SearchPage

	// Total : Total number of linked pages.
	Total int `db:"total" json:"total"`

	// Distribution : Distribution information for each categories.
	Distribution []models.CategoryCount `db:"distribution" json:"distribution"`
}

// NewSERPWithRatioView :
func NewSERPWithRatioView(pages []models.SearchPage, linkedPages []models.LinkedPage) []SERPWithRatio {
	return nil
}
