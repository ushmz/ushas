package models

import (
	"strings"
	"ushas/database"

	"github.com/pkg/errors"
)

// LinkedPage : Linked page information with icon URL.
type LinkedPage struct {
	// ID : ID of linked page.
	ID int `gorm:"primaryKey;not null;column:id"`

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
		e := translateGormError(err, l)
		e.Err = errors.WithStack(e.Err)
		return e
	}
	return nil
}

// GetLinkedPageByID  : Gets a record from table by ID.
func GetLinkedPageByID(id int) (*LinkedPage, error) {
	l := new(LinkedPage)
	db := database.GetDB()
	if err := db.Where("id = ?", id).First(l).Error; err != nil {
		e := translateGormError(err, id)
		e.Err = errors.WithStack(e.Err)
		return l, e
	}
	return l, nil
}

type LinkedPageWithPageID struct {
	PageID int `gorm:"not null; column:page_id"`

	LinkedPage
}

// GetLinkedPagesByIDs : Gets multiple records from table by given IDs.
func GetLinkedPagesByIDs(pageIDs []int, taskID, top int) ([]LinkedPageWithPageID, error) {
	l := []LinkedPageWithPageID{}
	db := database.GetDB()

	params := make([]interface{}, 0, len(pageIDs)+2)
	for _, v := range pageIDs {
		params = append(params, v)
	}
	params = append(params, taskID)
	params = append(params, top)

	query := db.Raw(`
		SELECT
			rel.page_id page_id, 
			similarweb_pages.id id,
			similarweb_pages.title title,
			similarweb_pages.url url,
			similarweb_pages.icon_path icon_path,
			similarweb_pages.category category
		FROM (
			SELECT
				page_id,
				similarweb_id,
				idf_rank
		FROM (
			SELECT
				page_id,
				similarweb_id,
				ROW_NUMBER() OVER (PARTITION BY page_id ORDER BY idf DESC) idf_rank
			FROM
				search_page_similarweb_relation
			WHERE
				page_id IN( ?`+strings.Repeat(", ?", len(pageIDs)-1)+`)
			AND
				task_id = ?
			ORDER BY
				page_id ASC
		) idf
		WHERE idf_rank <= ?) rel
		LEFT JOIN
			similarweb_pages ON rel.similarweb_id = similarweb_pages.id
		LEFT JOIN
			similarweb_categories ON similarweb_pages.category = similarweb_categories.id`, params...)

	rs, err := query.Rows()
	if err != nil {
		e := translateGormError(err, pageIDs)
		e.Err = errors.WithStack(e.Err)
		return l, e
	}

	defer rs.Close()
	for rs.Next() {
		lp := LinkedPageWithPageID{}
		if err := db.ScanRows(rs, &lp); err != nil {
			e := translateGormError(err, pageIDs)
			e.Err = errors.WithStack(e.Err)
			return l, e
		}
		l = append(l, lp)
	}

	return l, nil
}

type CategoryCountWithPageID struct {
	PageID int `gorm:"not null; column:page_id"`

	CategoryCount
}

// If we don't need total count, use this query
// SELECT
// 	*
// FROM (
// 	SELECT
// 		*,
// 		ROW_NUMBER() OVER(PARTITION BY categories.page_id) top
// 	FROM (
// 		SELECT DISTINCT
// 			rel.page_id,
// 			sc.category,
// 			COUNT(*) OVER (PARTITION BY rel.page_id, sp.category) category_count
// 		FROM
// 			search_page_similarweb_relation rel
// 		LEFT JOIN
// 			similarweb_pages sp ON rel.similarweb_id = sp.id
// 		LEFT JOIN
// 			similarweb_categories sc ON sp.category = sc.id
// 		WHERE
// 			page_id IN ( ? `+strings.Repeat(", ?", len(pageIDs)-1)+` )
// 		AND
// 			task_id = ?
// 		ORDER BY
// 			page_id, category_count DESC
// 	) categories
// ) categories_top
// WHERE
// 	top <= ?

// GetLinkedPagesRatioByIDs : Gets multiple records from table by given IDs.
func GetLinkedPagesRatioByIDs(pageIDs []int, taskID int) ([]CategoryCountWithPageID, error) {
	l := []CategoryCountWithPageID{}
	db := database.GetDB()

	params := make([]interface{}, 0, len(pageIDs)+2)
	for _, v := range pageIDs {
		params = append(params, v)
	}
	params = append(params, taskID)
	query := db.Raw(`
		SELECT DISTINCT
			rel.page_id,
			sc.category,
			COUNT(*) OVER (PARTITION BY rel.page_id, sp.category) category_count
		FROM
			search_page_similarweb_relation rel
		LEFT JOIN
			similarweb_pages sp ON rel.similarweb_id = sp.id
		LEFT JOIN
			similarweb_categories sc ON sp.category = sc.id
		WHERE
			page_id IN ( ? `+strings.Repeat(", ?", len(pageIDs)-1)+` )
		AND
			task_id = ?
		ORDER BY
			page_id, category_count DESC
	`, params...)

	rs, err := query.Rows()
	if err != nil {
		e := translateGormError(err, pageIDs)
		e.Err = errors.WithStack(e.Err)
		return l, e
	}

	defer rs.Close()
	for rs.Next() {
		lp := CategoryCountWithPageID{}
		if err := db.ScanRows(rs, &lp); err != nil {
			e := translateGormError(err, pageIDs)
			e.Err = errors.WithStack(e.Err)
			return l, e
		}
		l = append(l, lp)
	}

	return l, nil
}

// ListLinkedPage : Gets all records from table.
func ListLinkedPage() ([]LinkedPage, error) {
	lp := []LinkedPage{}
	db := database.GetDB()
	if err := db.Find(&lp).Error; err != nil {
		e := translateGormError(err, nil)
		e.Err = errors.WithStack(e.Err)
		return lp, e
	}
	return lp, nil
}

// UpdateLinkedPageByID : Updates a record with give ID in table.
func UpdateLinkedPageByID(l *LinkedPage) error {
	db := database.GetDB()
	if err := db.Save(l).Error; err != nil {
		e := translateGormError(err, l)
		e.Err = errors.WithStack(e.Err)
		return e
	}
	return nil
}

// DeleteLinkedPageByID : Deletes a record with given ID from table.
func DeleteLinkedPageByID(id int) error {
	db := database.GetDB()
	if err := db.Delete(&LinkedPage{}, id).Error; err != nil {
		e := translateGormError(err, id)
		e.Err = errors.WithStack(e.Err)
		return e
	}
	return nil
}
