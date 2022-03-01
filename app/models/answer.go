package models

import (
	"fmt"

	"ushas/database"
)

type Answer struct {
	// ID : The ID of user.
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

func CreateAnswer(a *Answer) error {
	db := database.GetDB()
	if err := db.Create(a).Error; err != nil {
		return RaiseInternalServerError(err, "Failed to create new `Answer` resource")
	}
	return nil
}

func GetAnswerByID(id int) (*Answer, error) {
	a := new(Answer)
	db := database.GetDB()
	err := db.Where("id = ?", id).First(a).Error
	if err != nil {
		return a, RaiseNotFoundError( err, fmt.Sprintf("Answer for ID %d is not found", id))
	}
	return a, nil
}

func ListAnswers() ([]Answer, error) {
	ans := []Answer{}
	db := database.GetDB()
	err := db.Find(&ans).Error
	if err != nil {
		return ans, RaiseInternalServerError(err, "Failed to fetch all User resource")
	}
	return ans, nil
}

func UpdateAnswer(a *Answer) error {
	db := database.GetDB()
	err := db.Save(a).Error
	if err != nil {
		return RaiseInternalServerError(
			err,
			fmt.Sprintf("Failed to Update User resource of ID %d", a.ID),
			a,
		)
	}
	return nil
}

func DeleteAnswer(id int) error {
	db := database.GetDB()
	err := db.Delete(&Answer{}, id).Error
	if err != nil {
		return RaiseInternalServerError(
			err,
			fmt.Sprintf("Failed to delete User resource of ID %d", id),
		)
	}
	return nil
}
