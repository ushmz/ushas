package views

import "ushas/models"

type UserView struct {
	// ID : The ID of user.
	ID int `json:"id"`

	// UID : External user_id (like crowdsourcing site).
	UID string `json:"uid"`
}

// UserResponseView : Struct for response body of `CreateUser` handler
type UserResponseView struct {
	// Exist : Given uid is exist on DB ot not.
	Exist bool `json:"exist"`

	// UserId : Unique ID used in DB.
	UserId int `json:"user"`

	// Secret : Generated secret string.
	Secret string `json:"secret"`

	// TaskIds : Shows the IDs that user perform
	TaskIds []int `json:"tasks"`

	// ConditionId : Assigned condition ID
	ConditionId int `json:"condition"`

	// GroupId : The ID assigned to the pair of "task IDs" and "condition ID"
	GroupId int `json:"group"`
}

func NewUserView(u *models.User) *UserView {
	v := &UserView{ID: u.ID, UID: u.UID}
	return v
}

func NewListUserView(us []models.User) []UserView {
	vs := []UserView{}
	for _, u := range us {
		v := UserView{ID: u.ID, UID: u.UID}
		vs = append(vs, v)
	}
	return vs
}

func NewCreateUserView(u *models.User) *UserResponseView {
	return nil
}
