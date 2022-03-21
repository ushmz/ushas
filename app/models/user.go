package models

import (
	"ushas/database"
)

// UserParam : Struct for request of `/signup` endpoint
type UserParam struct {
	// UID : User name/ID for label.
	UID string `json:"uid"`
}

// UserSimple : Struct for user information w/o secret.
type UserSimple struct {
	// ID : The ID of user.
	ID int `gorm:"primaryKey;column:id" json:"id"`

	// UID : External user_id (like crowdsourcing site).
	UID string `gorm:"unique;not null;column:uid" json:"uid"`
}

// User : Struct for user information.
type User struct {
	// ID : The ID of user.
	ID int `gorm:"primaryKey;column:id" json:"id"`

	// UID : External user_id (like crowdsourcing site).
	UID string `gorm:"unique;not null;column:uid" json:"uid"`

	// Secret : Generated secret string.
	Secret string `gorm:"unique;not null;column:generated_secret" json:"-"`
}

// CreateUser : Create new record into table.
func CreateUser(u *User) error {
	db := database.GetDB()
	err := db.Create(u).Error
	if err != nil {
		return translateGormError(err, "Failed to create new user", u)
	}
	return nil
}

// GetUserByID : Gets record from table by ID
func GetUserByID(id int) (*User, error) {
	u := new(User)
	db := database.GetDB()
	err := db.Where("id = ?", id).First(u).Error
	if err != nil {
		return u, translateGormError(err, "Failed to fetch user", id)
	}
	return u, nil
}

// GetUserByUID : Gets a record from table by UID.
func GetUserByUID(uid string) (*User, error) {
	u := new(User)
	db := database.GetDB()
	err := db.Where("uid = ?", uid).First(u).Error
	if err != nil {
		return u, translateGormError(err, "Failed to fetch user", uid)
	}
	return u, nil
}

// ListUsers : Gets all records from table.
func ListUsers() ([]User, error) {
	us := []User{}
	db := database.GetDB()
	err := db.Find(&us).Error
	if err != nil {
		return us, translateGormError(err, "Failed to fetch all user", nil)
	}
	return us, nil
}

// UpdateUser : Updates a record with given ID in table.
func UpdateUser(u *User) error {
	// To throw 404 error when requested user ID not found, we select the row with given ID at first.
	// If we use `Updates()`, it returns no error even if when given ID not found.
	// This also catch when the row actually not changed(e.g. same "uid" value requested), so we don't use this.
	// if rs.RowsAffected <= 0 {
	// 	return RaiseNotFoundError(rs.Error, fmt.Sprintf("User for ID %d is not found", u.ID))
	// }
	user := new(User)
	db := database.GetDB()
	if err := db.Where("id = ?", u.ID).First(user).Error; err != nil {
		return translateGormError(err, "Failed to update user", u)
	}

	user.UID = u.UID

	if err := db.Save(user).Error; err != nil {
		return translateGormError(err, "Failed to Update user", u)
	}

	return nil
}

// DeleteUser : Delete a record with given ID from table.
func DeleteUser(id int) error {
	db := database.GetDB()
	if err := db.Delete(&User{}, id).Error; err != nil {
		return translateGormError(err, "Failed to delete user", id)
	}
	return nil
}
