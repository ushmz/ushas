package models

import (
	"ushas/database"
)

// LinkedPage : Linked page information with icon URL.
type LinkedPage struct {
	// ID : ID of linked page.
	ID int `gorm:"unique;not null;column:id"`

	// Title : The title of linked page.
	Title string `gorm:"not null;column:title"`

	// URL : URL of the linked page.
	URL string `gorm:"not null;column:url"`

	// Icon : Favicon url of the page.
	Icon string `gorm:"not null;column:icon_path"`

	// Category : Category name of linked page.
	Category string `gorm:"not null;column:category"`
}

// TableName : Overrides table name used by `LinkedPage` struct.
func (*LinkedPage) TableName() string {
	return "similarweb_pages"
}

// CreateLinkedPage : Add new record for table.
func CreateLinkedPage(l *LinkedPage) error {
	db := database.GetDB()
	err := db.Create(l).Error
	if err != nil {
		return translateGormError(err, l)
	}
	return nil
}

// GetLinkedPageByID  : Gets a record from table by ID.
func GetLinkedPageByID(id int) (*LinkedPage, error) {
	l := new(LinkedPage)
	db := database.GetDB()
	if err := db.Where("id = ?", id).First(l).Error; err != nil {
		return l, translateGormError(err, id)
	}
	return l, nil
}

// GetLinkedPagesByIDs : Gets multiple records from table by given IDs.
func GetLinkedPagesByIDs(ids []int) (*LinkedPage, error) {
	l := new(LinkedPage)
	db := database.GetDB()
	if err := db.Where("id IN ?", ids).Find(l).Error; err != nil {
		return l, translateGormError(err, ids)
	}
	return l, nil
}

// ListLinkedPage : Gets all records from table.
func ListLinkedPage() ([]LinkedPage, error) {
	lp := []LinkedPage{}
	db := database.GetDB()
	if err := db.Find(&lp).Error; err != nil {
		return lp, translateGormError(err, nil)
	}
	return lp, nil
}

// UpdateLinkedPageByID : Updates a record with give ID in table.
func UpdateLinkedPageByID(l *LinkedPage) error {
	db := database.GetDB()
	if err := db.Save(l).Error; err != nil {
		return translateGormError(err, l)
	}
	return nil
}

// DeleteLinkedPageByID : Deletes a record with given ID from table.
func DeleteLinkedPageByID(id int) error {
	db := database.GetDB()
	if err := db.Delete(&LinkedPage{}, id).Error; err != nil {
		return translateGormError(err, id)
	}
	return nil
}
