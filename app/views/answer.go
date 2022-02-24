package views

import (
	"ushas/models"
)

type AnswerView struct {
	// ID : The ID of user.
	// ID int `json:"id"`

	// UserID : Means external ID.
	UserID int `json:"user"`

	// TaskID : The identity of task.
	TaskID int `json:"task"`

	// ConditionID : This point out which kind of task did user take.
	ConditionID int `json:"condition"`

	// Answer : The Url of evidence of users' decision.
	Answer string `json:"answer"`

	// Reason : The reason of users' decision.
	Reason string `json:"reason"`
}

func NewAnswerView(a *models.Answer) *AnswerView {
	v := &AnswerView{
		// ID:          a.Id,
		UserID:      a.UserID,
		TaskID:      a.TaskID,
		ConditionID: a.ConditionID,
		Answer:      a.Answer,
		Reason:      a.Reason,
	}
	return v
}
