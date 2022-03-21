package models

import (
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
		return translateGormError(err, "Failed to create answer", a)
	}
	return nil
}

// GetAnswerByID : Gets single record from table by ID.
func GetAnswerByID(id int) (*Answer, error) {
	a := new(Answer)
	db := database.GetDB()
	if err := db.Where("id = ?", id).First(a).Error; err != nil {
		return a, translateGormError(err, "Failed to get answer", id)
	}
	return a, nil
}

// ListAnswers : Gets all records from table.
func ListAnswers() ([]Answer, error) {
	ans := []Answer{}
	db := database.GetDB()
	if err := db.Find(&ans).Error; err != nil {
		return ans, translateGormError(err, "Failed to get all answers", nil)
	}
	return ans, nil
}

// UpdateAnswer : Updates record in table.
func UpdateAnswer(a *Answer) error {
	db := database.GetDB()
	// [FIXME] `db.Save()` upserts record.
	if err := db.Save(a).Error; err != nil {
		return translateGormError(err, "Failed to update answer", a)
	}
	return nil
}

// DeleteAnswer : Deletes a record from table.
func DeleteAnswer(id int) error {
	db := database.GetDB()
	if err := db.Delete(&Answer{}, id).Error; err != nil {
		return translateGormError(err, "Failed to delete answer", id)
	}
	return nil
}
