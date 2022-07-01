package models

import (
	"ushas/database"

	"github.com/pkg/errors"
)

var (
	resultsPerSERP = 10
)

// SearchPage : Each of search result pages.
type SearchPage struct {
	// PageID : ID of search page.
	PageID int `gorm:"primaryKey;not null;column:id" json:"id"`

	// Title : The title of each search result page.
	Title string `gorm:"not null; column:title" json:"title"`

	// URL : URL of each search result page.
	URL string `gorm:"not null; column:url" json:"url"`

	// Snippet : Snippet of each search result page.
	Snippet string `gorm:"not null; column:snippet" json:"snippet"`
}

// ListSERP :
func ListSERP(taskID, offset int) ([]SearchPage, error) {
	pages := []SearchPage{}
	db := database.GetDB()

	query := db.Where("task_id = ?", taskID).
		Limit(resultsPerSERP).
		Offset(offset * resultsPerSERP).
		Find(&pages)

	if err := query.Error; err != nil {
		e := translateGormError(err, nil)
		e.Err = errors.WithStack(e.Err)
		return pages, e
	}
	return pages, nil
}

// SearchPageWithIcon : The list of this type struct will be returned as a response of `serp` endpoint.
type SearchPageWithIcon struct {
	SearchPage `gorm:"embedded"`

	// Linked : Users' behavioral data that probably leaked to third party.
	// For more detail, see `Linked` type.
	Linked []LinkedPage `json:"linked"`
}

// SearchPageWithLinkedPageRow : `SearchPage` with `LinkedPage` query result row struct
type SearchPageWithLinkedPageRow struct {
	// PageID : ID of search result page.
	PageID int `db:"page_id" json:"page"`

	// ID : ID of linked page.
	ID int `db:"id" json:"id"`

	// Title : The title of linked page.
	Title string `db:"title" json:"title"`

	// URL : URL of the linked page.
	URL string `db:"url" json:"url"`

	// Icon : Favicon url of the page.
	Icon string `db:"icon_path" json:"icon"`

	// Category : Category name of linked page.
	Category string `db:"category" json:"category"`
}

// CategoryCount : Distribution information for each categories.
type CategoryCount struct {
	// Category : Category name.
	Category string `gorm:"column:category" json:"category"`
	// Count : Total number of pages.
	Count int `gorm:"column:category_count" json:"count"`
}

// SearchPageWithLinkedPageRatioRow : `SearchPage` with `LinkedPage` query result row struct
type SearchPageWithLinkedPageRatioRow struct {
	// PageID : ID of search page.
	PageID int `db:"page_id" json:"page"`

	// Category : Linked page category name.
	Category string `db:"category"`

	// CategoryCount : Total number of linked pages in the category.
	CategoryCount int `db:"category_count"`
}
