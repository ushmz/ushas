package models

import (
	"fmt"

	"ushas/database"
)

// Answer : Submitted answer for the tasks.
type Answer struct {
	// ID : The ID of answer.
	ID int `gorm:"unique;not null;column:id" json:"id"`

	// UserID : Means external ID.
	UserID int `gorm:"not null;column:user_id" json:"user"`

	// TaskID : The identity of task.
	TaskID int `gorm:"not null;column:task_id" json:"task"`

	// ConditionID : This point out which kind of task did user take.
	ConditionID int `gorm:"not null;column:condition_id" json:"condition"`

	// Answer : The Url of evidence of users' decision.
	Answer string `gorm:"not null;column:answer" json:"answer"`

	// Reason : The reason of users' decision.
	Reason string `gorm:"not null;column:reason" json:"reason"`
}

// CreateAnswer : Create new record.
func CreateAnswer(a *Answer) error {
	db := database.GetDB()
	if err := db.Create(a).Error; err != nil {
		return RaiseInternalServerError(err, "Failed to create new `Answer` resource")
	}
	return nil
}

// GetAnswerByID : Gets single record from table by ID.
func GetAnswerByID(id int) (*Answer, error) {
	a := new(Answer)
	db := database.GetDB()
	if err := db.Where("id = ?", id).First(a).Error; err != nil {
		return a, RaiseNotFoundError(err, fmt.Sprintf("Answer for ID %d is not found", id))
	}
	return a, nil
}

// ListAnswers : Gets all records from table.
func ListAnswers() ([]Answer, error) {
	ans := []Answer{}
	db := database.GetDB()
	if err := db.Find(&ans).Error; err != nil {
		return ans, RaiseInternalServerError(err, "Failed to fetch all User resource")
	}
	return ans, nil
}

// UpdateAnswer : Updates record in table.
func UpdateAnswer(a *Answer) error {
	db := database.GetDB()
	// [FIXME] `db.Save()` upserts record.
	if err := db.Save(a).Error; err != nil {
		return RaiseInternalServerError(err, fmt.Sprintf("Failed to Update User resource of ID %d", a.ID), a)
	}
	return nil
}

// DeleteAnswer : Deletes a record from table.
func DeleteAnswer(id int) error {
	db := database.GetDB()
	if err := db.Delete(&Answer{}, id).Error; err != nil {
		return RaiseInternalServerError(err, fmt.Sprintf("Failed to delete User resource of ID %d", id))
	}
	return nil
}
