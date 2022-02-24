package views

import "ushas/models"

type UserView struct {
	// ID : The ID of user.
	ID int `json:"id"`

	// Uid : External user_id (like crowdsourcing site).
	Uid string `json:"uid"`
}

func NewUserView(u *models.User) *UserView {
	v := &UserView{
		ID:  u.ID,
		Uid: u.Uid,
	}
	return v
}
