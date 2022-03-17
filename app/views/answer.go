package views

import (
	"ushas/models"
)

// AnswerView : Response data of "/answer" endpoint.
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

// NewAnswerView : Returns single answer response data.
func NewAnswerView(answer *models.Answer) *AnswerView {
	v := &AnswerView{
		// ID:          a.Id,
		UserID:      answer.UserID,
		TaskID:      answer.TaskID,
		ConditionID: answer.ConditionID,
		Answer:      answer.Answer,
		Reason:      answer.Reason,
	}
	return v
}

// NewListAnswerView : Return listed answer response data.
func NewListAnswerView(answers []models.Answer) []*AnswerView {
	vs := []*AnswerView{}
	for _, a := range answers {
		v := &AnswerView{
			UserID:      a.UserID,
			TaskID:      a.TaskID,
			ConditionID: a.ConditionID,
			Answer:      a.Answer,
			Reason:      a.Reason,
		}
		vs = append(vs, v)
	}
	return vs
}
