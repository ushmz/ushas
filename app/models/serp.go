package models

import "gorm.io/gorm"

// SearchPage : Each of search result pages.
type SearchPage struct {
	gorm.Model

	// PageId : ID of search page.
	PageId int `db:"id" json:"id"`

	// Title : The title of each search result page.
	Title string `db:"title" json:"title"`

	// Url : Url of each search result page.
	Url string `db:"url" json:"url"`

	// Snippet : Snippet of each search result page.
	Snippet string `db:"snippet" json:"snippet"`
}

// SearchPageWithIcon : The list of this type struct will be returned as a response of `serp` endpoint.
type SearchPageWithIcon struct {
	gorm.Model

	SearchPage `gorm:"embedded"`

	// Linked : Users' behavioral data that probably leaked to third party.
	// For more detail, see `Linked` type.
	Linked []LinkedPage `json:"linked"`
}

// SearchPageWithLinkedPageRow : `SearchPage` with `LinkedPage` query result row struct
type SearchPageWithLinkedPageRow struct {
	gorm.Model

	// PageId : ID of search result page.
	PageId int `db:"page_id" json:"page"`

	// Id : ID of linked page.
	Id int `db:"id" json:"id"`

	// Title : The title of linked page.
	Title string `db:"title" json:"title"`

	// Url : Url of the linked page.
	Url string `db:"url" json:"url"`

	// Icon : Favicon url of the page.
	Icon string `db:"icon_path" json:"icon"`

	// Category : Category name of linked page.
	Category string `db:"category" json:"category"`
}

// SearchPageWithRatio : The list of this type struct will be returned as a response of `serp` endpoint.
type SearchPageWithRatio struct {
	gorm.Model

	SearchPage `gorm:"embedded"`

	// Total : Total number of linked pages.
	Total int `db:"total" json:"total"`

	// Distribution : Distribution information for each categories.
	Distribution []CategoryCount `db:"distribution" json:"distribution"`
}

// SearchPageWithLinkedPageRatioRow : `SearchPage` with `LinkedPage` query result row struct
type SearchPageWithLinkedPageRatioRow struct {
	gorm.Model

	// PageId : ID of search page.
	PageId int `db:"page_id" json:"page"`

	// Category : Linked page category name.
	Category string `db:"category"`

	// CategoryCount : Total number of linked pages in the category.
	CategoryCount int `db:"category_count"`
}
