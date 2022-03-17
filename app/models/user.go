package models

import (
	"fmt"
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
	ID int `gorm:"unique;not null;column:id" json:"id"`

	// UID : External user_id (like crowdsourcing site).
	UID string `gorm:"unique;not null;column:uid" json:"uid"`
}

// User : Struct for user information.
type User struct {
	// ID : The ID of user.
	ID int `gorm:"unique;not null;column:id" json:"id"`

	// UID : External user_id (like crowdsourcing site).
	UID string `gorm:"unique;not null;column:uid" json:"uid"`

	// Secret : Generated secret string.
	Secret string `gorm:"unique;not null;column:generated_secret" json:"secret"`
}

// CreateUser : Create new record into table.
func CreateUser(u *User) error {
	db := database.GetDB()
	err := db.Create(u).Error
	if err != nil {
		return RaiseInternalServerError(err, "Failed to create new `User` resource")
	}
	return nil
}

// GetUserByID : Gets record from table by ID
func GetUserByID(id int) (*User, error) {
	u := new(User)
	db := database.GetDB()
	err := db.Where("id = ?", id).First(u).Error
	if err != nil {
		return u, RaiseNotFoundError(err, fmt.Sprintf("User for ID %d is not found", id))
	}
	return u, nil
}

// GetUserByUID : Gets a record from table by UID.
func GetUserByUID(uid string) (*User, error) {
	u := new(User)
	db := database.GetDB()
	err := db.Where("uid = ?", uid).First(u).Error
	if err != nil {
		return u, RaiseNotFoundError(err, fmt.Sprintf("User for ID %s is not found", uid))
	}
	return u, nil
}

// ListUsers : Gets all records from table.
func ListUsers() ([]User, error) {
	us := []User{}
	db := database.GetDB()
	err := db.Find(&us).Error
	if err != nil {
		return us, RaiseInternalServerError(err, "Failed to fetch all User resource")
	}
	return us, nil
}

// UpdateUser : Updates a record with given ID in table.
func UpdateUser(u *User) error {
	db := database.GetDB()
	err := db.Model(&User{}).Updates(u).Error
	if err != nil {
		return RaiseInternalServerError(
			err,
			fmt.Sprintf("Failed to Update User resource of ID %d", u.ID),
			u,
		)
	}
	return nil
}

// DeleteUser : Delete a record with given ID from table.
func DeleteUser(id int) error {
	db := database.GetDB()
	err := db.Delete(&User{}, id).Error
	if err != nil {
		return RaiseInternalServerError(
			err,
			fmt.Sprintf("Failed to delete User resource of ID %d", id),
		)
	}
	return nil
}
