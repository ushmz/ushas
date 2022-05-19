package models

import "ushas/database"

var (
	resultsPerSERP = 10
)

// SearchPage : Each of search result pages.
type SearchPage struct {
	// PageID : ID of search page.
	PageID int `db:"id" json:"id"`

	// Title : The title of each search result page.
	Title string `db:"title" json:"title"`

	// URL : URL of each search result page.
	URL string `db:"url" json:"url"`

	// Snippet : Snippet of each search result page.
	Snippet string `db:"snippet" json:"snippet"`
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
		return pages, translateGormError(err, nil)
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
	Category string `json:"category"`
	// Count : Total number of pages.
	Count int `json:"count"`
}

// SearchPageWithRatio : The list of this type struct will be returned as a response of `serp` endpoint.
type SearchPageWithRatio struct {
	SearchPage `gorm:"embedded"`

	// Total : Total number of linked pages.
	Total int `db:"total" json:"total"`

	// Distribution : Distribution information for each categories.
	Distribution []CategoryCount `db:"distribution" json:"distribution"`
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
