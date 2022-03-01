package models

// LinkedPage : Linked page information with icon URL.
type LinkedPage struct {
	// Id : ID of linked page.
	Id int `gorm:"unique;not null;column:id"`

	// Title : The title of linked page.
	Title string `gorm:"not null;column:title"`

	// Url : Url of the linked page.
	Url string `gorm:"not null;column:url"`

	// Icon : Favicon url of the page.
	Icon string `gorm:"not null;column:icon_path"`

	// Category : Category name of linked page.
	Category string `gorm:"not null;column:category"`
}
