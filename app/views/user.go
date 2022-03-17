package views

import "ushas/models"

// UserView : Response data of "user" endpoint.
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

	// UserID : Unique ID used in DB.
	UserID int `json:"user"`

	// Secret : Generated secret string.
	Secret string `json:"secret"`

	// TaskIds : Shows the IDs that user perform
	TaskIds []int `json:"tasks"`

	// ConditionID : Assigned condition ID
	ConditionID int `json:"condition"`

	// GroupID : The ID assigned to the pair of "task IDs" and "condition ID"
	GroupID int `json:"group"`
}

// NewUserView : Returns single user response data.
func NewUserView(u *models.User) *UserView {
	v := &UserView{ID: u.ID, UID: u.UID}
	return v
}

// NewListUserView : Returns listed user response.
func NewListUserView(us []models.User) []UserView {
	vs := []UserView{}
	for _, u := range us {
		v := UserView{ID: u.ID, UID: u.UID}
		vs = append(vs, v)
	}
	return vs
}

// NewCreateUserView : Returns data when new user is created.
func NewCreateUserView(u *models.User) *UserResponseView {
	return nil
}
